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
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/internal/repositories"
)

func GetSteamAppDetails(appId uint32) (models.Game, error) {
	// Set Stub's needs_update = false
	err := repositories.SetStubNeedsUpdateStatus(appId, false)
	if err != nil {
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Get App Details from Steam API
	response, err := http.Get(fmt.Sprintf("https://store.steampowered.com/api/appdetails/?appids=%d&l=english", appId))
	if err != nil {
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
		return models.Game{}, errors.Join(err, revertErr)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	if !jsoniter.Get(body, strconv.FormatUint(uint64(appId), 10), "success").ToBool() {
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
		return models.Game{}, errors.Join(errors.New("jsoniter: could not confirm success"), revertErr)
	}

	var publicAppDetailsSteam models.AppDetailsSteam
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal([]byte(jsoniter.Get(body, strconv.FormatUint(uint64(appId), 10), "data").ToString()), &publicAppDetailsSteam)
	if err != nil {
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Check if App Requested from Steam API is a game
	if publicAppDetailsSteam.Type != "game" {
		typeErr := repositories.SetStubType(appId, publicAppDetailsSteam.Type)
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, false, true)
		return models.Game{}, errors.Join(errors.New("assertion: not a game type app"), revertErr, typeErr)
	}

	// Get App Details from SteamSpy API
	response, err = http.Get(fmt.Sprintf("https://steamspy.com/api.php?request=appdetails&appid=%d", appId))
	if err != nil {
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
		return models.Game{}, errors.Join(err, revertErr)
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	var publicAppDetailsSteamSpy models.AppDetailsSteamSpy
	err = json.Unmarshal(body, &publicAppDetailsSteamSpy)
	if err != nil {
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	// Set Stub's type
	err = repositories.SetStubType(appId, publicAppDetailsSteam.Type)
	if err != nil {
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
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
			revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
			return models.Game{}, errors.Join(err, revertErr)
		}
		game.ReleaseDate = date
	}

	for _, elem := range publicAppDetailsSteam.Screenshots {
		game.Screenshots = append(game.Screenshots, models.GameScreenshot{PathThumbnail: elem.PathThumbnail, PathFull: elem.PathFull})
	}

	for _, elem := range publicAppDetailsSteam.Movies {
		game.Movies = append(game.Movies, models.GameMovie{Name: elem.Name, Thumbnail: elem.Thumbnail, Webm: models.GameMovieEntry{P480: elem.Webm.P480, Max: elem.Webm.Max}, Mp4: models.GameMovieEntry{P480: elem.Mp4.P480, Max: elem.Mp4.Max}})
	}

	for _, elem := range publicAppDetailsSteam.Genres {
		game.Genres = append(game.Genres, elem.Description)
	}

	for key := range publicAppDetailsSteamSpy.Tags {
		game.Tags = append(game.Tags, key)
	}

	game.ReviewsPositive = publicAppDetailsSteamSpy.Positive
	game.ReviewsNegative = publicAppDetailsSteamSpy.Negative

	if owners := strings.Split(publicAppDetailsSteamSpy.Owners, ".."); len(owners) == 2 {
		ownersMin, err := strconv.ParseUint(strings.ReplaceAll(strings.TrimSpace(owners[0]), ",", ""), 10, 32)
		if err != nil {
			revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
			return models.Game{}, errors.Join(err, revertErr)
		}
		game.OwnersMin = uint32(ownersMin)
		ownersMax, err := strconv.ParseUint(strings.ReplaceAll(strings.TrimSpace(owners[1]), ",", ""), 10, 32)
		if err != nil {
			revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
			return models.Game{}, errors.Join(err, revertErr)
		}
		game.OwnersMax = uint32(ownersMax)
	}

	_, err = repositories.InsertGames([]models.Game{game})
	if err != nil {
		revertErr := repositories.SetStubNeedsUpdateAndSkipStatuses(appId, true, true)
		return models.Game{}, errors.Join(err, revertErr)
	}

	return game, nil
}
