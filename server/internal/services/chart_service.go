package services

import (
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetChartsDatasets() (models.ChartsDatasets, error) {
	totalStubs, err := repositories.CountStubsRawFilter(bson.D{})
	if err != nil {
		return models.ChartsDatasets{}, err
	}

	totalUntouchedStubs, err := repositories.CountStubsRawFilter(bson.D{{Key: "$and", Value: bson.A{bson.D{{Key: "needs_update", Value: true}}, bson.D{{Key: "skip", Value: false}}, bson.D{{Key: "ignore", Value: false}}}}})
	if err != nil {
		return models.ChartsDatasets{}, err
	}

	stubTypeDataset, err := repositories.GroupStubsByType()
	if err != nil {
		return models.ChartsDatasets{}, err
	}

	totalUnreleasedGames, err := repositories.CountGamesRawFilter(bson.D{{Key: "coming_soon", Value: true}})
	if err != nil {
		return models.ChartsDatasets{}, err
	}

	totalGamesByReleaseYear, err := repositories.GroupGamesByReleaseYear()
	if err != nil {
		return models.ChartsDatasets{}, err
	}

	var dataset models.ChartsDatasets
	dataset.TotalCountOfStubs = uint32(totalStubs)
	dataset.TotalCountOfUntouchedStubs = uint32(totalUntouchedStubs)
	dataset.TotalCountOfUnreleasedYetGames = uint32(totalUnreleasedGames)
	dataset.TotalStubsByType = stubTypeDataset
	dataset.TotalGamesReleasedByYears = totalGamesByReleaseYear

	return dataset, nil
}
