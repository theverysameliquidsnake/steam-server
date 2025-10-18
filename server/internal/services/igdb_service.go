package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/pkg/utils"
)

func GetIGDBResponse(appId uint32) (*models.GameIGDB, error) {
	// Get Details from IGDB API
	client, err := utils.UseProxyClient()
	if err != nil {
		return nil, err
	}

	// Seek in external sources
	payload := fmt.Sprintf("fields *; where uid = \"%d\" & external_game_source = 1;", appId)
	request, err := http.NewRequest("POST", "https://api.igdb.com/v4/external_games", strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "text/plain")
	for key, value := range configs.GetIGDBHeaders() {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Check by external game uid
	var externalGames []models.ExternalGameIGDB
	err = json.Unmarshal([]byte(jsoniter.Get(body).ToString()), &externalGames)
	if err != nil {
		return nil, err
	}

	if len(externalGames) == 0 {
		return nil, errors.New("igdb: external game not found")
	}

	// Get Details from IGDB API
	payload = fmt.Sprintf("fields *, genres.*, themes.*, franchises.*, collections.*, keywords.*; where id = %d;", externalGames[0].Game)
	request, err = http.NewRequest("POST", "https://api.igdb.com/v4/games", strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "text/plain")
	for key, value := range configs.GetIGDBHeaders() {
		request.Header.Set(key, value)
	}

	response, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Check by igdb id
	var gamesIGDB []models.GameIGDB
	err = json.Unmarshal([]byte(jsoniter.Get(body).ToString()), &gamesIGDB)
	if err != nil {
		return nil, err
	}

	if len(gamesIGDB) == 0 {
		return nil, errors.New("igdb: game not found")
	}

	return &gamesIGDB[0], nil
}
