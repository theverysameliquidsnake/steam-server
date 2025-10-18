package repositories

import (
	"context"
	"os"

	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetLogsCollection() *mongo.Collection {
	return configs.GetMongoClient().Database(os.Getenv("MONGO_DATABASE")).Collection(os.Getenv("MONGO_LOGS_COLLECTION"))
}

// Create
func InsertLogs(logs []models.Log) error {
	_, err := GetLogsCollection().InsertMany(context.Background(), logs)
	if err != nil {
		return err
	}

	return nil
}
