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

	handlers.InitMongoRoutes()
	handlers.InitStubRoutes()
	handlers.ConnectGameRoutes()

	err = configs.RunRouter()
	if err != nil {
		log.Fatal(err)
	}
}
