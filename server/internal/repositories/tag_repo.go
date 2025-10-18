package repositories

import (
	"context"
	"os"

	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetTagsCollection() *mongo.Collection {
	return configs.GetMongoClient().Database(os.Getenv("MONGO_DATABASE")).Collection(os.Getenv("MONGO_TAGS_COLLECTION"))
}

// Read
func GetAllTags() (map[uint32]string, error) {
	cursor, err := GetTagsCollection().Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []models.Tag
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	tagsMap := make(map[uint32]string)
	for _, elem := range results {
		tagsMap[elem.Id] = elem.Name
	}

	return tagsMap, nil
}

// Create
func InsertTags(tags []models.Tag) ([]any, error) {
	result, err := GetTagsCollection().InsertMany(context.Background(), tags)
	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}

// Delete
func DeleteTags() error {
	err := GetTagsCollection().Drop(context.Background())
	if err != nil {
		return err
	}

	return nil
}
