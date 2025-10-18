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

func GetSteamCMDResponse(appId uint32) (*models.SteamCMD, error) {
	// Get App Details from SteamCMD API
	client, err := utils.UseProxyClient()
	if err != nil {
		return nil, err
	}

	response, err := client.Get(fmt.Sprintf("https://api.steamcmd.net/v1/info/%d", appId))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if jsoniter.Get(body, "status").ToString() != "success" {
		return nil, errors.New("jsoniter: could not confirm success from SteamCMD")
	}

	var steamCMD models.SteamCMD
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal([]byte(jsoniter.Get(body, "data", strconv.FormatUint(uint64(appId), 10), "common").ToString()), &steamCMD)
	if err != nil {
		return nil, err
	}

	return &steamCMD, nil
}
