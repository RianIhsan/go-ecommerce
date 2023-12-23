package main

import (
	"github.com/RianIhsan/go-ecommerce/external/database"
	"github.com/RianIhsan/go-ecommerce/internal/config"
	"log"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	if db != nil {
		log.Println("Success connect to database")
	}
}
