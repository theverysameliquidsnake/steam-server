package services

import (
	"context"

	"github.com/theverysameliquidsnake/steam-db/configs"
)

func ResetMongo() error {
	if err := configs.GetMongoDatabase().Drop(context.Background()); err != nil {
		return err
	}

	return nil
}
