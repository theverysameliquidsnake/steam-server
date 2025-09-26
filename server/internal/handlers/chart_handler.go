package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/services"
)

func InitChartRoutes() {
	router := configs.GetGinRouter()

	chartGroup := router.Group("/chart")

	// GET /chart/dataset
	chartGroup.GET("/dataset", func(ctx *gin.Context) {
		results, err := services.GetChartsDatasets()
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
			"data":    results,
			"message": "Returned charts datasets",
			"error":   "",
		})
	})
}
