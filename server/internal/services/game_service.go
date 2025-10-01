package services

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/internal/repositories"
	"github.com/theverysameliquidsnake/steam-db/pkg/utils"
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
		date, err := time.Parse("2 Jan, 2006", publicAppDetailsSteam.ReleaseDate.Date)
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

	for _, elem := range publicAppDetailsSteam.Genres {
		game.Genres = append(game.Genres, elem.Description)
	}

	// Insert first part of update of game
	_, err = repositories.InsertGames([]models.Game{game})
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

	return game, nil
}
