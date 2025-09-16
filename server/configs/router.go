package configs

import "github.com/gin-gonic/gin"

var router *gin.Engine

func CreateRouter() *gin.Engine {
	router = gin.Default()
	return router
}

func RunRouter() error {
	if err := router.Run(); err != nil {
		return err
	}

	return nil
}

func GetGinRouter() *gin.Engine {
	return router
}
