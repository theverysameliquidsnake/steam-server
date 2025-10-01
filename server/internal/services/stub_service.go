package services

import (
	"errors"
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/internal/repositories"
	"github.com/theverysameliquidsnake/steam-db/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func RefreshStubs() (int, error) {
	client, err := utils.UseProxyClient()
	if err != nil {
		return -1, err
	}

	response, err := client.Get("https://api.steampowered.com/ISteamApps/GetAppList/v2/")
	if err != nil {
		return -1, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return -1, err
	}

	var publicApps []models.PublicAppSteam
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal([]byte(jsoniter.Get(body, "applist", "apps").ToString()), &publicApps)
	if err != nil {
		return -1, err
	}

	existingStubsMap := make(map[uint32]models.Stub)
	existingStubs, err := repositories.FindStubsRawFilter(bson.D{})
	if err != nil {
		return -1, err
	}
	for _, elem := range existingStubs {
		existingStubsMap[elem.AppId] = elem
	}

	var stubs []models.Stub
	for _, elem := range publicApps {
		_, isStubExists := existingStubsMap[elem.AppId]
		if !isStubExists {
			stubs = append(stubs, models.Stub{AppId: elem.AppId, Name: elem.Name, New: true})
		}
	}

	if len(stubs) == 0 {
		return 0, nil
	}

	result, err := repositories.InsertStubs(stubs)
	if err != nil {
		return -1, err
	}

	return len(result), nil
}

func GetStubRequiredToUpdate() (models.Stub, error) {
	utils.Lock()
	defer utils.Unlock()
	result, err := repositories.FindStubsRawFilter(bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "new", Value: true}},
			bson.D{{Key: "error", Value: false}},
			bson.D{{Key: "ignore", Value: false}},
		}},
	})
	if err != nil {
		return models.Stub{}, err
	}

	if len(result) > 0 {
		// Set Stub's "new" = false
		err = repositories.SetStubNewStatus(result[0].AppId, false)
		if err != nil {
			revertErr := repositories.SetStubNewStatus(result[0].AppId, true)
			return models.Stub{}, errors.Join(err, revertErr)
		}

		return result[0], nil
	}

	return models.Stub{}, errors.New("stub not found")
}

func GetAllStubs(offset int64) ([]models.Stub, error) {
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "appid", Value: 1}})
	opts.SetLimit(50)
	opts.SetSkip(offset)

	result, err := repositories.FindStubsRawFilterOptions(bson.D{}, *opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}
