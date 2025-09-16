package handlers

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/services"
)

func ConnectGameRoutes() {
	router := configs.GetGinRouter()

	gameGroup := router.Group("/game")
	gameGroup.PUT("/insert/:count", func(ctx *gin.Context) {
		count, err := strconv.Atoi(ctx.Param("count"))
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{})
			return
		}

		// Add counter
		log.Println(count)

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

func testTimer() {
	configs.GetGinRouter().GET("/timer", func(c *gin.Context) {
		for range time.Tick(2 * time.Second) {
			go func() {
				log.Println(time.Now())
			}()
		}
		c.JSON(200, gin.H{
			"success": true,
			"message": "updated game stub",
		})
	})
}
