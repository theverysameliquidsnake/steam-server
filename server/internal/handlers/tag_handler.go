package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/services"
)

func InitTagRoutes() {
	router := configs.GetGinRouter()

	tagGroup := router.Group("/tag")

	// PUT /tag/refresh
	tagGroup.PUT("/refresh", func(ctx *gin.Context) {
		count, err := services.RefreshTags()
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
			"message": fmt.Sprintf("Inserted %d tag(s)", count),
			"error":   "",
		})
	})
}
