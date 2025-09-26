package services

import (
	"context"

	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/pkg/consts"
)

func ResetMongo() error {
	if err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Drop(context.Background()); err != nil {
		return err
	}

	return nil
}
