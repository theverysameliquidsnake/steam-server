package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/repositories"
	"github.com/theverysameliquidsnake/steam-db/internal/services"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func InitStubRoutes() {
	router := configs.GetGinRouter()

	stubGroup := router.Group("/stub")

	// PUT /stub/refresh
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

	// GET /stub/request
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

	// GET /stub/all/<skip count>
	stubGroup.GET("/all/:offset", func(ctx *gin.Context) {
		offset, err := strconv.ParseInt(ctx.Param("offset"), 10, 64)
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{
				"success": false,
				"message": "",
				"error":   "Internal server error",
			})
			return
		}

		result, err := services.GetAllStubs(offset)
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

	// PATCH /stub/ignore
	stubGroup.PATCH("/ignore", func(ctx *gin.Context) {
		appId, err := strconv.ParseUint(ctx.PostForm("appid"), 10, 32)
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{
				"success": false,
				"message": "",
				"error":   "Internal server error",
			})
			return
		}

		ignore, err := strconv.ParseBool(ctx.PostForm("ignore"))
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{
				"success": false,
				"message": "",
				"error":   "Internal server error",
			})
			return
		}

		err = repositories.SetStubIgnoreStatus(uint32(appId), ignore)
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
			"message": fmt.Sprintf("Added %d to ignore", appId),
			"error":   "",
		})
	})

	// GET /stub/count
	stubGroup.GET("/count", func(ctx *gin.Context) {
		result, err := repositories.CountStubsRawFilter(bson.D{})
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
			"data": map[string]int64{
				"count": result,
			},
			"message": fmt.Sprintf("%d stubs count", result),
			"error":   "",
		})
	})
}
