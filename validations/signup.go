package validations

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type SignUp struct {
	Username string `validate:"required,min=3,max=50" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=50" json:"password"`
}

func SignUpValidation(c *fiber.Ctx) error {
	var errors []*Error

	body := new(SignUp)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := Validate.Struct(body); err != nil {
		fmt.Printf("Errors: %+v\n", err.(validator.ValidationErrors))

		for _, err := range err.(validator.ValidationErrors) {
			var el Error
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}

		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}
