package services

import (
	"context"

	"github.com/theverysameliquidsnake/steam-db/configs"
	consts "github.com/theverysameliquidsnake/steam-db/pkg"
)

func ResetMongo() error {
	if err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Drop(context.Background()); err != nil {
		return err
	}

	return nil
}
