package repositories

import (
	"context"

	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	consts "github.com/theverysameliquidsnake/steam-db/pkg"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func FindStubsRawFilter(filter bson.D) ([]models.Stub, error) {
	cursor, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_STUB_COLLECTION).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var results []models.Stub
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func InsertStubs(stubs []models.Stub) ([]any, error) {
	result, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_STUB_COLLECTION).InsertMany(context.Background(), stubs)
	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}

func SetStubType(appId uint32, appType string) error {
	_, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_STUB_COLLECTION).UpdateOne(
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

func SetStubNeedsUpdateStatus(appId uint32, needsUpdate bool) error {
	_, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_STUB_COLLECTION).UpdateOne(
		context.Background(),
		bson.M{"appid": appId},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "needs_update", Value: needsUpdate}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func SetStubSkipStatus(appId uint32, skip bool) error {
	_, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_STUB_COLLECTION).UpdateOne(
		context.Background(),
		bson.M{"appid": appId},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "skip", Value: skip}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func SetStubNeedsUpdateAndSkipStatuses(appId uint32, needsUpdate bool, skip bool) error {
	_, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_STUB_COLLECTION).UpdateOne(
		context.Background(),
		bson.M{"appid": appId},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "needs_update", Value: needsUpdate}, {Key: "skip", Value: skip}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}
