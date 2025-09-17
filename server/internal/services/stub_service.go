package services

import (
	"errors"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RefreshStubs() error {
	response, err := http.Get("https://api.steampowered.com/ISteamApps/GetAppList/v2/")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var publicApps []models.PublicAppSteam
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal([]byte(jsoniter.Get(body, "applist", "apps").ToString()), &publicApps)
	if err != nil {
		return err
	}

	existingStubsMap := make(map[uint32]models.Stub)
	existingStubs, err := repositories.FindStubsRawFilter(bson.D{})
	if err != nil {
		return err
	}
	for _, elem := range existingStubs {
		existingStubsMap[elem.AppId] = elem
	}

	var stubs []models.Stub
	for _, elem := range publicApps {
		_, isStubExists := existingStubsMap[elem.AppId]
		if !isStubExists {
			stubs = append(stubs, models.Stub{AppId: elem.AppId, Name: elem.Name, NeedsUpdate: true, Skip: false})
		}
	}
	_, err = repositories.InsertStubs(stubs)
	if err != nil {
		return err
	}

	return nil
}

func GetStubRequiredToUpdate() (models.Stub, error) {
	result, err := repositories.FindStubsRawFilter(bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "needs_update", Value: true}},
			bson.D{{Key: "skip", Value: false}},
		}},
	})
	if err != nil {
		return models.Stub{}, err
	}

	if len(result) > 0 {
		return result[0], nil
	}

	return models.Stub{}, errors.New("stub not found")
}
