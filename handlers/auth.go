package handlers

import (
	"time"

	"github.com/SaiHLu/fiber-jwt/common/config"
	"github.com/SaiHLu/fiber-jwt/common/utils"
	"github.com/SaiHLu/fiber-jwt/database"
	"github.com/SaiHLu/fiber-jwt/models"
	"github.com/SaiHLu/fiber-jwt/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SignUp(c *fiber.Ctx) error {
	db := database.DB

	user := new(models.User)
	c.BodyParser(&user)

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "errors": err.Error()})
	}

	user.Password = hash
	if err = db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "errors": err.Error()})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

func Login(c *fiber.Ctx) error {
	db := database.DB

	body := new(validations.Login)
	c.BodyParser(&body)

	var user models.User
	db.Where(&models.User{Email: body.Email}).First(&user)

	if !utils.CompareHashPassword(user.Password, body.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid credentials.",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
