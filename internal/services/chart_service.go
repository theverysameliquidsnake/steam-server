package services

import (
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetStubsByStatusDataset() ([]models.StubByStatus, error) {
	ignored, err := repositories.CountStubsRawFilter(bson.D{
		{Key: "ignore", Value: true},
	})
	if err != nil {
		return []models.StubByStatus{}, err
	}

	errorStub, err := repositories.CountStubsRawFilter(bson.D{
		{Key: "error", Value: true},
	})
	if err != nil {
		return []models.StubByStatus{}, err
	}

	igdbUpdate, err := repositories.CountStubsRawFilter(bson.D{
		{Key: "igdb_update", Value: true},
	})
	if err != nil {
		return []models.StubByStatus{}, err
	}

	steamcmdUpdate, err := repositories.CountStubsRawFilter(bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "steamcmd_update", Value: true}},
			bson.D{{Key: "igdb_update", Value: false}},
		}},
	})
	if err != nil {
		return []models.StubByStatus{}, err
	}

	steamUpdate, err := repositories.CountStubsRawFilter(bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "steam_update", Value: true}},
			bson.D{{Key: "steamcmd_update", Value: false}},
			bson.D{{Key: "igdb_update", Value: false}},
		}},
	})
	if err != nil {
		return []models.StubByStatus{}, err
	}

	orpStub, err := repositories.CountStubsRawFilter(bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "new", Value: false}},
			bson.D{{Key: "error", Value: true}},
			bson.D{{Key: "steam_update", Value: false}},
			bson.D{{Key: "steamcmd_update", Value: false}},
			bson.D{{Key: "igdb_update", Value: false}},
		}},
	})
	if err != nil {
		return []models.StubByStatus{}, err
	}

	newStub, err := repositories.CountStubsRawFilter(bson.D{
		{Key: "new", Value: true},
	})
	if err != nil {
		return []models.StubByStatus{}, err
	}

	return []models.StubByStatus{
		{Status: "new", Count: uint32(newStub)},
		{Status: "orphan", Count: uint32(orpStub)},
		{Status: "steam update", Count: uint32(steamUpdate)},
		{Status: "steamcmd update", Count: uint32(steamcmdUpdate)},
		{Status: "igdb update", Count: uint32(igdbUpdate)},
		{Status: "error", Count: uint32(errorStub)},
		{Status: "ignore", Count: uint32(ignored)},
	}, nil
}

func GetStubsByTypeDataset() ([]models.StubTypeCount, error) {
	stubTypeDataset, err := repositories.GroupStubsByType()
	if err != nil {
		return []models.StubTypeCount{}, err
	}

	return stubTypeDataset, nil
}

func GetGamesByYears() (models.GameReleaseYearDataset, error) {
	totalUnreleasedGames, err := repositories.CountGamesRawFilter(bson.D{{Key: "coming_soon", Value: true}})
	if err != nil {
		return models.GameReleaseYearDataset{}, err
	}

	totalGamesByReleaseYear, err := repositories.GroupGamesByReleaseYear()
	if err != nil {
		return models.GameReleaseYearDataset{}, err
	}

	return models.GameReleaseYearDataset{
		TotalUnreleasedYetGames:   uint32(totalUnreleasedGames),
		TotalGamesReleasedByYears: totalGamesByReleaseYear,
	}, nil
}

func GetChartsDatasets() (models.ChartsDatasets, error) {
	var dataset models.ChartsDatasets

	stubsByStatus, err := GetStubsByStatusDataset()
	if err != nil {
		return models.ChartsDatasets{}, err
	}
	dataset.TotalStubsByStatus = stubsByStatus

	stubsByType, err := GetStubsByTypeDataset()
	if err != nil {
		return models.ChartsDatasets{}, err
	}
	dataset.TotalStubsByType = stubsByType

	gamesByYearDataset, err := GetGamesByYears()
	if err != nil {
		return models.ChartsDatasets{}, err
	}
	dataset.GamesByYearDataset = gamesByYearDataset

	return dataset, nil
}
