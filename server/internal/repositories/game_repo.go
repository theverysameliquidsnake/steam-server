package repositories

import (
	"context"

	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	consts "github.com/theverysameliquidsnake/steam-db/pkg"
)

func InsertGames(games []models.Game) ([]any, error) {
	result, err := configs.GetMongoClient().Database(consts.MONGO_DATABASE).Collection(consts.MONGO_GAME_COLLECTION).InsertMany(context.Background(), games)
	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}
