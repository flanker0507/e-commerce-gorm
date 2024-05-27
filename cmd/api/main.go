package main

import (
	"e-commerce-gorm/apps/auth"
	"e-commerce-gorm/apps/product"
	"e-commerce-gorm/eksternal/database"
	"e-commerce-gorm/internal/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectMysql(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	auth.Init(router, db)
	product.Init(router, db)

	router.Listen(config.Cfg.App.Port)

}
