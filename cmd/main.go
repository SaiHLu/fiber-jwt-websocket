package main

import (
	"github.com/SaiHLu/fiber-jwt/database"
	"github.com/SaiHLu/fiber-jwt/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Use(cors.New())
	app.Use(recover.New())

	app.Static("/", "./public")

	routes.SetupRoutes(app)

	log.Info(app.Listen(":3000"))
}
