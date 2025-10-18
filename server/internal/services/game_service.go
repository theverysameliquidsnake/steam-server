package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/internal/repositories"
)

func ConstructGameObject(appId uint32) (*models.Game, error) {
	logMsg := fmt.Sprintf("Processing %d ...", appId)
	configs.PrintLog(logMsg)
	err := repositories.InsertLogs([]models.Log{{Timestamp: time.Now(), Message: logMsg, AppId: appId}})
	if err != nil {
		return nil, err
	}

	// Get Details from Steam API
	publicAppDetailsSteam, err := GetSteamResponse(appId)
	if err != nil {
		return nil, SetStubErrorAndRevert(appId, err)
	}

	// Set Stub's type
	err = repositories.SetStubType(appId, publicAppDetailsSteam.Type)
	if err != nil {
		return nil, SetStubErrorAndRevert(appId, err)
	}

	logMsg = fmt.Sprintf("Success from Steam API for %d ...", appId)
	configs.PrintLog(logMsg)
	err = repositories.InsertLogs([]models.Log{{Timestamp: time.Now(), Message: logMsg, AppId: appId}})
	if err != nil {
		return nil, err
	}

	// Check if App Requested from Steam API is a game
	if publicAppDetailsSteam.Type != "game" {
		revertErr := repositories.SetStubIgnoreStatus(appId, true)
		return nil, errors.Join(revertErr, errors.New("assertion: not a game type app"))
	}

	// Set stub's "steam update"
	err = repositories.SetStubNumberUpdateStatus(appId, 1, true)
	if err != nil {
		return nil, SetStubErrorAndRevert(appId, err)
	}

	// Get Details from SteamCMD API
	steamCMD, err := GetSteamCMDResponse(publicAppDetailsSteam.SteamAppId)
	if err != nil {
		return nil, SetStubErrorAndRevert(appId, err)
	}

	logMsg = fmt.Sprintf("Success from SteamCMD API for %d ...", appId)
	configs.PrintLog(logMsg)
	err = repositories.InsertLogs([]models.Log{{Timestamp: time.Now(), Message: logMsg, AppId: appId}})
	if err != nil {
		return nil, err
	}

	// Set stub's "steamcmd update"
	err = repositories.SetStubNumberUpdateStatus(appId, 2, true)
	if err != nil {
		return nil, SetStubErrorAndRevert(appId, err)
	}

	// Get Details from IGDB API
	gameIGDB, err := GetIGDBResponse(publicAppDetailsSteam.SteamAppId)
	if err != nil {
		return nil, SetStubErrorAndRevert(appId, err)
	}

	logMsg = fmt.Sprintf("Success from IGDB API for %d ...", appId)
	configs.PrintLog(logMsg)
	err = repositories.InsertLogs([]models.Log{{Timestamp: time.Now(), Message: logMsg, AppId: appId}})
	if err != nil {
		return nil, err
	}

	// Set stub's "igdb update"
	err = repositories.SetStubNumberUpdateStatus(appId, 3, true)
	if err != nil {
		return nil, SetStubErrorAndRevert(appId, err)
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
			return nil, SetStubErrorAndRevert(appId, err)
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

	game.AI = len(steamCMD.AIContentType) > 0

	tagsMap, err := repositories.GetAllTags()
	if err != nil {
		return nil, SetStubErrorAndRevert(appId, err)
	}

	for _, value := range steamCMD.StoreTags {
		tagId, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return nil, SetStubErrorAndRevert(appId, err)
		}

		if len(tagsMap[uint32(tagId)]) > 0 {
			game.Tags = append(game.Tags, tagsMap[uint32(tagId)])
		} else {
			game.HasUnmappedTags = true
			game.Tags = append(game.Tags, value)
		}
	}

	for _, value := range gameIGDB.Genres {
		game.GenresIGDB = append(game.GenresIGDB, value.Name)
	}

	for _, value := range gameIGDB.Themes {
		game.ThemesIGDB = append(game.ThemesIGDB, value.Name)
	}

	for _, value := range gameIGDB.Franchises {
		game.FranchisesIGDB = append(game.FranchisesIGDB, value.Name)
	}

	for _, value := range gameIGDB.Series {
		game.SeriesIGDB = append(game.SeriesIGDB, value.Name)
	}

	for _, value := range gameIGDB.Keywords {
		game.KeywordsIGDB = append(game.KeywordsIGDB, value.Name)
	}

	// Insert document
	_, err = repositories.InsertGames([]models.Game{game})
	if err != nil {
		return nil, SetStubErrorAndRevert(appId, err)
	}

	logMsg = fmt.Sprintf("Inserted document for %d ...", appId)
	configs.PrintLog(logMsg)
	err = repositories.InsertLogs([]models.Log{{Timestamp: time.Now(), Message: logMsg, AppId: appId}})
	if err != nil {
		return nil, err
	}

	return &game, nil
}
