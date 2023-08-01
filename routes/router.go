package routes

import (
	"fmt"

	"github.com/SaiHLu/fiber-jwt/handlers"
	"github.com/SaiHLu/fiber-jwt/middlewares"
	"github.com/SaiHLu/fiber-jwt/validations"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	api.Post("/auth/signup", validations.SignUpValidation, handlers.SignUp)
	api.Post("/auth/login", validations.LoginValidation, handlers.Login)

	protected := app.Group("/protected", middlewares.Protected())

	protected.Get("/me", func(c *fiber.Ctx) error {
		token := c.Locals("user").(*jwt.Token)
		fmt.Println("token: ", token.Claims.(jwt.MapClaims))
		return c.JSON("OK")
	})

	app.Get("/ws", middlewares.WSMiddleware, websocket.New(handlers.WSHandler))
}
