package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/services"
)

func InitStubRoutes() {
	router := configs.GetGinRouter()

	stubGroup := router.Group("/stub")
	stubGroup.PUT("/refresh", func(ctx *gin.Context) {
		records, err := services.RefreshStubs()
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
			"message": fmt.Sprintf("Inserted %d record(s)", records),
			"error":   "",
		})
	})

	stubGroup.GET("/request", func(ctx *gin.Context) {
		result, err := services.GetStubRequiredToUpdate()
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
			"data": map[string]uint32{
				"appid": result.AppId,
			},
			"message": fmt.Sprintf("Requested: %s (%d)", result.Name, result.AppId),
			"error":   "",
		})
	})

	stubGroup.GET("/all", func(ctx *gin.Context) {
		result, err := services.GetAllStubs()
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
			"data":    result,
			"message": fmt.Sprintf("Returned %d stubs", len(result)),
			"error":   "",
		})
	})
}
