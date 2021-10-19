package main

import (
	"fmt"
	"github.com/tegarap/cekak/api/config"
	"github.com/tegarap/cekak/api/database"
	"github.com/tegarap/cekak/api/handlers"
	"github.com/tegarap/cekak/api/models"
	"github.com/tegarap/cekak/api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main()  {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open(cfg.ConnectionString), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	db.AutoMigrate(models.Site{})

	siteModel := database.NewUrlDbModel(db)
	siteHandler := handlers.NewUrlHandler(siteModel)

	app := fiber.New()
	routes.Routes(app, siteHandler)

	if err = app.Listen(fmt.Sprintf(":%d", cfg.HttpPort)); err != nil {
		log.Fatalln("Error starting server", err.Error())
	}
}