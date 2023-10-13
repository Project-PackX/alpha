package controllers

import (
	"PackX/exceptions"
	"PackX/initializers"
	"PackX/middleware"
	"PackX/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// List all users
func ListUsers(c *fiber.Ctx) error {

	var users []models.User // Slice that will contain all users

	// Execute 'SELECT * FROM public.users' query
	result := initializers.DB.Find(&users)

	// Error handling
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Something bad happened during the query",
		})
	}

	// Separately treated case where no record exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "There are no users in the database",
		})
	}

	// Returning the users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"users":   users,
	})
}

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
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token":   generateJwtToken(*newUser),
		"message": "User added successfully",
	})
}

func Login(c *fiber.Ctx) error {

	// Make a temporary user model with the given json input data
	loginUser := new(models.User)
	if err := c.BodyParser(loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.CreateInvalidInputException("Something went wrong"))
	}

	// Getting the credentials from the user
	loginEmail := loginUser.Email
	loginPassword := loginUser.Password

	// Search for the email in DB
	var userMatch models.User
	initializers.DB.First(&userMatch, "email = ?", loginEmail)

	if userMatch.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Username or password is incorrect",
		})
	}

	// Check the passwords
	if bcrypt.CompareHashAndPassword([]byte(userMatch.Password), []byte(loginPassword)) != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Username or password is incorrect",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"token":   generateJwtToken(userMatch),
		"message": "Login successful",
		"name":    userMatch.Name,
		"email":   userMatch.Email,
	})
}

// Get the access of the user based on URL {id}
func GetAccessLevel(c *fiber.Ctx) error {

	id := c.Params("id")

	// Get the user with the given id
	var user models.User
	initializers.DB.First(&user, "id = ?", id)

	// Return the value
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": user.AccessLevel,
	})
}

// Get all package info wchich belong to the specific user id based on URL
func GetPackagesUnderUser(c *fiber.Ctx) error {

	// Getting the {id} from URL
	id := c.Params("id")

	// Getting the packs from the desired user
	var packs []models.Package
	initializers.DB.Find(&packs, "user_id = ?", id)

	// Sending back the list of packages with every information
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": packs,
	})
}

func generateJwtToken(user models.User) string {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expiration time set to 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(middleware.SecretKey)

	if err != nil {
		return err.Error()
	}

	return signedToken
}
