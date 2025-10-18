package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/services"
)

func InitMongoRoutes() {
	router := configs.GetGinRouter()

	mongoGroup := router.Group("/mongo")

	// DELETE /mongo/drop
	mongoGroup.DELETE("/drop", func(ctx *gin.Context) {
		err := services.ResetMongo()
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
			"message": "MongoDB cleared",
			"error":   "",
		})
	})
}
