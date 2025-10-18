package services

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/pkg/utils"
)

func GetSteamResponse(appId uint32) (*models.AppDetailsSteam, error) {
	// Get App Details from Steam API
	client, err := utils.UseProxyClient()
	if err != nil {
		return nil, err
	}

	response, err := client.Get(fmt.Sprintf("https://store.steampowered.com/api/appdetails/?appids=%d&l=english", appId))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if !jsoniter.Get(body, strconv.FormatUint(uint64(appId), 10), "success").ToBool() {
		return nil, errors.New("jsoniter: could not confirm success from Steam API")
	}

	var publicAppDetailsSteam models.AppDetailsSteam
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal([]byte(jsoniter.Get(body, strconv.FormatUint(uint64(appId), 10), "data").ToString()), &publicAppDetailsSteam)
	if err != nil {
		return nil, err
	}

	return &publicAppDetailsSteam, nil
}
