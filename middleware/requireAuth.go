package middleware

import (
	"PackX/initializers"
	"PackX/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func RequireJwtTokenAuth(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Message": "Unauthorized"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the token using the secret key
		claims := token.Claims.(jwt.MapClaims)

		// If expired
		exp := claims["exp"].(float64)
		if exp <= float64(time.Now().Unix()) {
			return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Message": "Unauthorized"})
		}

		// If user does not exist
		userId := claims["user_id"]
		var foundUser models.User
		initializers.DB.First(&foundUser, "id = ?", userId)
		if foundUser.ID == 0 {
			return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Message": "Unauthorized"})
		}

		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Message": "Unauthorized"})
	}

	// Token is valid; proceed with the request
	return c.Next()
}
