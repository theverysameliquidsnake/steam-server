package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/services"
)

func InitGameRoutes() {
	router := configs.GetGinRouter()

	gameGroup := router.Group("/game")

	// PUT /game/insert/<app id>
	gameGroup.PUT("/insert/:appid", func(ctx *gin.Context) {
		appId, err := strconv.ParseUint(ctx.Param("appid"), 10, 32)
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{
				"success": false,
				"message": "",
				"error":   "Internal server error",
			})
			return
		}

		game, err := services.GetSteamAppDetails(uint32(appId))
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{
				"success": false,
				"message": "",
				"error":   "Internal server error",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"success": true,
			"message": fmt.Sprintf("Added: %s (%d)", game.Name, game.AppId),
			"error":   "",
		})
	})
}
