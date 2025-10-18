package repositories

import (
	"context"
	"errors"
	"os"

	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetStubCollection() *mongo.Collection {
	return configs.GetMongoClient().Database(os.Getenv("MONGO_DATABASE")).Collection(os.Getenv("MONGO_STUBS_COLLECTION"))
}

// Read
func FindStubsRawFilter(filter bson.D) ([]models.Stub, error) {
	cursor, err := GetStubCollection().Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []models.Stub
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func FindStubsRawFilterOptions(filter bson.D, opts options.FindOptionsBuilder) ([]models.Stub, error) {
	cursor, err := GetStubCollection().Find(context.Background(), filter, &opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []models.Stub
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func FindStubRawFilterAndUpdate(filter bson.D, update bson.D) (models.Stub, error) {
	var stub models.Stub
	err := GetStubCollection().FindOneAndUpdate(context.Background(), filter, update).Decode(&stub)
	if err != nil {
		return models.Stub{}, err
	}

	return stub, nil
}

func CountStubsRawFilter(filter bson.D) (int64, error) {
	count, err := GetStubCollection().CountDocuments(context.Background(), filter)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func GroupStubsByType() ([]models.StubTypeCount, error) {
	cursor, err := GetStubCollection().Aggregate(context.Background(), mongo.Pipeline{bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$type"}, {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []models.StubTypeCount
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

// Create
func InsertStubs(stubs []models.Stub) ([]any, error) {
	result, err := GetStubCollection().InsertMany(context.Background(), stubs)
	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}

// Update
func SetStubType(appId uint32, appType string) error {
	_, err := GetStubCollection().UpdateOne(
		context.Background(),
		bson.M{"appid": appId},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "type", Value: appType}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func SetStubNewStatus(appId uint32, new bool) error {
	_, err := GetStubCollection().UpdateOne(
		context.Background(),
		bson.M{"appid": appId},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "new", Value: new}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func SetStubNumberUpdateStatus(appId uint32, updateNumber int, updateStatus bool) error {
	var field string
	switch updateNumber {
	case 1:
		field = "steam_update"
	case 2:
		field = "steamcmd_update"
	case 3:
		field = "igdb_update"
	default:
		return errors.New("update field not specified")
	}

	_, err := GetStubCollection().UpdateOne(
		context.Background(),
		bson.M{"appid": appId},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: field, Value: updateStatus}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func SetStubErrorStatus(appId uint32, error bool) error {
	_, err := GetStubCollection().UpdateOne(
		context.Background(),
		bson.M{"appid": appId},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "error", Value: error}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func SetStubIgnoreStatus(appId uint32, ignore bool) error {
	_, err := GetStubCollection().UpdateOne(
		context.Background(),
		bson.M{"appid": appId},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "ignore", Value: ignore}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}
