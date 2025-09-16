package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/services"
)

func InitStubRoutes() {
	router := configs.GetGinRouter()

	stubGroup := router.Group("/stub")
	stubGroup.PUT("/refresh", func(ctx *gin.Context) {
		err := services.RefreshStubs()
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{})
			return
		}
		ctx.JSON(200, gin.H{})
	})
}
