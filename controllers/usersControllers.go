package controllers

import (
	"PackX/exceptions"
	"PackX/initializers"
	"PackX/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterNewUser(c *fiber.Ctx) error {

	// Creating the new user
	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.CreateInvalidInputException("Invalid User"))
	}

	// Check whether a user with same email exists in the db
	// If the user already exists, return an error
	if initializers.DB.Where("email = ?", newUser.Email).First(&models.User{}).Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.CreateUserAlreadyExistsException("User already exists"))
	}

	// Password hashing
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	newUser.Password = string(hashedPassword)

	// Save the user to the database
	initializers.DB.Create(&newUser)

	// Send a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "User added successfully",
	})
}

func Login(c *fiber.Ctx) error {

	// Make a temporary user model with the given json input data
	loginUser := new(models.User)
	if err := c.BodyParser(loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.CreateInvalidInputException("Something went wrong"))
	}

	// Getting the credentials from the user
	login_email := loginUser.Email
	login_passw := loginUser.Password

	// Search for the email in DB
	var userMatch models.User
	initializers.DB.First(&userMatch, "email = ?", login_email)

	// Check the passwords
	if bcrypt.CompareHashAndPassword([]byte(userMatch.Password), []byte(login_passw)) == nil {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"Message": "Login successfully",
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": "Wrong password",
		})
	}
}

// Get the access of the user based on URL {id}
func GetAccessLevel(c *fiber.Ctx) error {

	id := c.Params("id")

	// Get the user with the given id
	var user models.User
	initializers.DB.First(&user, "id = ?", id)

	// Return the value
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": user.AccessLevel,
	})
}
