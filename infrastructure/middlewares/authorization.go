package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/saufiroja/go-graphql-boilerplate/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// get token from header
	tokenString := c.Get("Authorization")
	fmt.Println(tokenString)
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"message": "Unauthorized",
		})
	}

	// validate token
	token, err := utils.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"message": "Unauthorized",
		})
	}

	// get claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"message": "Unauthorized",
		})
	}

	// get email from claims
	email, ok := claims["email"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    401,
			"message": "Unauthorized",
		})
	}

	// set email to context
	c.Locals("email", email)

	return c.Next()
}
