package repositories

import (
	"context"

	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/pkg/consts"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Read
func FindGamesRawFilter(filter bson.D) ([]models.Game, error) {
	cursor, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_GAME_COLLECTION).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var results []models.Game
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func CountGamesRawFilter(filter bson.D) (int64, error) {
	count, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_GAME_COLLECTION).CountDocuments(context.Background(), filter)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func GroupGamesByReleaseYear() ([]models.GameReleaseYearCount, error) {
	cursor, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_GAME_COLLECTION).Aggregate(context.Background(), mongo.Pipeline{bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: bson.D{{Key: "$year", Value: "$release_date"}}}, {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}}})
	if err != nil {
		return nil, err
	}

	var results []models.GameReleaseYearCount
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

// Create
func InsertGames(games []models.Game) ([]any, error) {
	result, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_GAME_COLLECTION).InsertMany(context.Background(), games)
	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}
