package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/services"
)

func ConnectGameRoutes() {
	router := configs.GetGinRouter()

	gameGroup := router.Group("/game")
	gameGroup.PUT("/insert", func(ctx *gin.Context) {
		result, err := services.GetStubRequiredToUpdate()
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{})
			return
		}

		err = services.GetSteamAppDetails(result.AppId)
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{})
			return
		}

		ctx.JSON(200, gin.H{})
	})
}
