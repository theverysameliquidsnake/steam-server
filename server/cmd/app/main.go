package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/theverysameliquidsnake/steam-db/configs"
	"github.com/theverysameliquidsnake/steam-db/internal/handlers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	_, err := configs.ConnectToMongo()
	if err != nil {
		log.Fatal(err)
	}
	defer configs.DisconnectFromMongo()

	configs.CreateRouter()
	configs.SetCORS()

	handlers.InitMongoRoutes()
	handlers.InitStubRoutes()
	handlers.InitTagRoutes()
	handlers.InitGameRoutes()
	handlers.InitChartRoutes()

	//utils.StartPlaywright()
	//defer utils.StopPlaywright()

	configs.InitIGDBToken()

	err = configs.RunRouter()
	if err != nil {
		log.Fatal(err)
	}
}
