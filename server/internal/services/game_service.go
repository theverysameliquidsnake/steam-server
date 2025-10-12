package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/internal/repositories"
	"github.com/theverysameliquidsnake/steam-db/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetSteamAppDetails(appId uint32) (models.Game, error) {
	// Get App Details from Steam API
	client, err := utils.UseProxyClient()
	if err != nil {
		return models.Game{}, err
	}

	response, err := client.Get(fmt.Sprintf("https://store.steampowered.com/api/appdetails/?appids=%d&l=english", appId))
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	if !jsoniter.Get(body, strconv.FormatUint(uint64(appId), 10), "success").ToBool() {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(errors.New("jsoniter: could not confirm success"), revertErr)
	}

	var publicAppDetailsSteam models.AppDetailsSteam
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal([]byte(jsoniter.Get(body, strconv.FormatUint(uint64(appId), 10), "data").ToString()), &publicAppDetailsSteam)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Check if App Requested from Steam API is a game
	if publicAppDetailsSteam.Type != "game" {
		typeErr := repositories.SetStubType(appId, publicAppDetailsSteam.Type)
		revertErr := repositories.SetStubIgnoreStatus(appId, true)
		return models.Game{}, errors.Join(errors.New("assertion: not a game type app"), revertErr, typeErr)
	}

	// Set Stub's type
	err = repositories.SetStubType(appId, publicAppDetailsSteam.Type)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Construct Mongo document
	game := models.Game{
		AppId:       publicAppDetailsSteam.SteamAppId,
		Name:        publicAppDetailsSteam.Name,
		Description: publicAppDetailsSteam.ShortDescription,
		HeaderImage: publicAppDetailsSteam.HeaderImage,
		Developers:  publicAppDetailsSteam.Developers,
		Publishers:  publicAppDetailsSteam.Publishers,
	}

	game.ComingSoon = publicAppDetailsSteam.ReleaseDate.ComingSoon
	if len(publicAppDetailsSteam.ReleaseDate.Date) > 0 {
		pattern := "Jan 2, 2006"
		if publicAppDetailsSteam.ReleaseDate.Date[0] >= '0' && publicAppDetailsSteam.ReleaseDate.Date[0] <= '9' {
			pattern = "2 Jan, 2006"
		}
		date, err := time.Parse(pattern, publicAppDetailsSteam.ReleaseDate.Date)
		if err != nil {
			revertErr := repositories.SetStubErrorStatus(appId, true)
			return models.Game{}, errors.Join(err, revertErr)
		}
		game.ReleaseDate = date
	}

	for _, elem := range publicAppDetailsSteam.Screenshots {
		game.Screenshots = append(game.Screenshots, models.GameScreenshot{PathThumbnail: elem.PathThumbnail, PathFull: elem.PathFull})
	}

	for _, elem := range publicAppDetailsSteam.Movies {
		game.Movies = append(game.Movies, models.GameMovie{
			Name:      elem.Name,
			Thumbnail: elem.Thumbnail,
			Webm: models.GameMovieEntry{
				P480: elem.Webm.P480,
				Max:  elem.Webm.Max,
			},
			Mp4: models.GameMovieEntry{
				P480: elem.Mp4.P480,
				Max:  elem.Mp4.Max,
			},
		})
	}

	// Insert first part of update of game
	resultIds, err := repositories.InsertGames([]models.Game{game})
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Set stub's "first update"
	err = repositories.SetStubFirstUpdateStatus(appId, true)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Get Details from IGDB API
	payload := fmt.Sprintf("fields *; where uid = \"%d\" & external_game_source = 1;", publicAppDetailsSteam.SteamAppId)
	request, err := http.NewRequest("POST", "https://api.igdb.com/v4/external_games", strings.NewReader(payload))
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}
	request.Header.Set("Content-Type", "text/plain")
	for key, value := range configs.GetIGDBHeaders() {
		request.Header.Set(key, value)
	}
	response, err = client.Do(request)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Check by external game uid
	var externalGames []models.ExternalGame
	err = json.Unmarshal([]byte(jsoniter.Get(body).ToString()), &externalGames)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	if len(externalGames) == 0 {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Get Details from IGDB API
	payload = fmt.Sprintf("fields *, genres.*, themes.*, franchises.*, collections.*, keywords.*; where id = %d;", externalGames[0].Game)
	request, err = http.NewRequest("POST", "https://api.igdb.com/v4/games", strings.NewReader(payload))
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}
	request.Header.Set("Content-Type", "text/plain")
	for key, value := range configs.GetIGDBHeaders() {
		request.Header.Set(key, value)
	}
	response, err = client.Do(request)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Check by igdb id
	var gamesIGDB []models.GameIGDB
	err = json.Unmarshal([]byte(jsoniter.Get(body).ToString()), &gamesIGDB)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	if len(gamesIGDB) == 0 {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Parsed values
	var parsedGenres []string
	for _, value := range gamesIGDB[0].Genres {
		parsedGenres = append(parsedGenres, value.Name)
	}

	var parsedThemes []string
	for _, value := range gamesIGDB[0].Themes {
		parsedThemes = append(parsedThemes, value.Name)
	}

	var parsedFranchises []string
	for _, value := range gamesIGDB[0].Franchises {
		parsedFranchises = append(parsedFranchises, value.Name)
	}

	var parsedSeries []string
	for _, value := range gamesIGDB[0].Series {
		parsedSeries = append(parsedSeries, value.Name)
	}

	var parsedKeywords []string
	for _, value := range gamesIGDB[0].Keywords {
		parsedKeywords = append(parsedKeywords, value.Name)
	}

	// Insert second part of update of game
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "genres", Value: parsedGenres},
		{Key: "themes", Value: parsedThemes},
		{Key: "franchises", Value: parsedFranchises},
		{Key: "series", Value: parsedSeries},
		{Key: "keywords", Value: parsedKeywords},
	}}}

	err = repositories.UpdateGameSecondTime(bson.D{{Key: "_id", Value: resultIds[0]}}, update)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Parse Steam page for remaining details
	/*parsedGame, err := utils.ParseSteamPage(appId)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Insert second part of update of game
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "tags", Value: parsedGame.Tags},
		{Key: "reviews_positive", Value: parsedGame.ReviewsPositive},
		{Key: "reviews_negative", Value: parsedGame.ReviewsNegative},
		{Key: "review_score", Value: (float32(parsedGame.ReviewsPositive) / (float32(parsedGame.ReviewsPositive) + float32(parsedGame.ReviewsNegative))) * 100},
	}}}
	err = repositories.UpdateGameSecondTime(bson.D{{Key: "_id", Value: resultIds[0]}}, update)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}*/

	// Set stub's "second update"
	err = repositories.SetStubSecondUpdateStatus(appId, true)
	if err != nil {
		revertErr := repositories.SetStubErrorStatus(appId, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	return game, nil
}
