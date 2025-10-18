package repositories

import (
	"context"
	"os"

	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetGameCollection() *mongo.Collection {
	return configs.GetMongoClient().Database(os.Getenv("MONGO_DATABASE")).Collection(os.Getenv("MONGO_GAMES_COLLECTION"))
}

// Read
func FindGamesRawFilter(filter bson.D) ([]models.Game, error) {
	cursor, err := GetGameCollection().Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []models.Game
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func CountGamesRawFilter(filter bson.D) (int64, error) {
	count, err := GetGameCollection().CountDocuments(context.Background(), filter)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func GroupGamesByReleaseYear() ([]models.GameReleaseYearCount, error) {
	cursor, err := GetGameCollection().Aggregate(context.Background(), mongo.Pipeline{bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{{Key: "$year", Value: "$release_date"}}},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		}},
	}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []models.GameReleaseYearCount
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

// Create
func InsertGames(games []models.Game) ([]any, error) {
	result, err := GetGameCollection().InsertMany(context.Background(), games)
	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}

// Update
func UpdateGameSecondTime(filter bson.D, update bson.D) error {
	_, err := GetGameCollection().UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
