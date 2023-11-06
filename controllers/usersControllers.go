package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/Project-PackX/backend/enums"
	"github.com/Project-PackX/backend/exceptions"
	"github.com/Project-PackX/backend/initializers"
	"github.com/Project-PackX/backend/middleware"
	"github.com/Project-PackX/backend/models"
	"github.com/Project-PackX/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var SUBJECT_RESET_PASSWORD = "Password reset code"
var BODY_RESET_PASSWORD = `
	<span><h4 style="color:black;">Dear Customer</h4></span>
	<p style="color:black;">a password reset was requested by someone. </p>
	<p>Your password reset code is: %s</p>
	<p>The code is valid for 1 hour.</p>
	<p>If the request was not made by you, please contact our support immediately.</p>
	<p>Sincerely,<br>PackX</br></p>
`
var NUM_DIGITS_RESET_CODE = 14

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

	// Adding permission level based on email
	// If ends with 'packx.hu' -> admin
	// If ends with 'packx-courier.hu' -> courier
	// Normal user otherwise
	if strings.Split(newUser.Email, "@")[1] == "packx.hu" {
		newUser.AccessLevel = enums.AccessLevel.Admin
	} else if strings.Split(newUser.Email, "@")[1] == "packx-courier.hu" {
		newUser.AccessLevel = enums.AccessLevel.Courier
	} else {
		newUser.AccessLevel = enums.AccessLevel.Normal
	}

	// Save the user to the database
	initializers.DB.Create(&newUser)

	// Send a success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token":   generateJwtToken(*newUser),
		"message": "User added successfully",
		"user":    newUser,
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
	initializers.DB.Where("email = ?", loginEmail).First(&userMatch)

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
		"token": generateJwtToken(userMatch),
		"user":  userMatch,
	})
}

func GetUserById(c *fiber.Ctx) error {

	id := c.Params("id")

	var user models.User
	initializers.DB.Where("id = ?", id).First(&user)

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User was not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func EditUser(c *fiber.Ctx) error {

	userInput := new(models.User)
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.CreateInvalidInputException("Bad input format"))
	}

	var userToEdit models.User
	userToEditID := c.Params("id")
	initializers.DB.First(&userToEdit, "ID = ?", userToEditID)

	if userToEditID == "0" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User was not found",
		})
	}

	if userInput.Name != "" && userInput.Name != userToEdit.Name {
		userToEdit.Name = userInput.Name
	}

	if userInput.Email != "" && userInput.Email != userToEdit.Email {
		userToEdit.Email = userInput.Email
	}

	if userInput.Address != "" && userInput.Address != userToEdit.Address {
		userToEdit.Address = userInput.Address
	}

	if userInput.Phone != "" && userInput.Phone != userToEdit.Phone {
		userToEdit.Phone = userInput.Phone
	}

	initializers.DB.Save(&userToEdit)

	return c.Status(fiber.StatusAccepted).JSON(userToEdit)
}

func DeleteUserById(c *fiber.Ctx) error {

	id := c.Params("id")
	tokenString := c.Get("Authorization")

	var userToDelete models.User
	initializers.DB.Where("id = ?", id).Find(&userToDelete)

	if userToDelete.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User was not found",
		})
	}

	if !userIsAllowedToDelete(userToDelete, tokenString) {
		c.SendStatus(fiber.StatusForbidden)
	}

	initializers.DB.Unscoped().Where("id = ?", id).Delete(&models.User{})

	return c.SendStatus(fiber.StatusOK)
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

func SetAccessLevel(c *fiber.Ctx) error {
	// Making a user model with the given parameters
	user := new(models.User)

	// Check error
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Hibás kérés",
		})
	}

	// Update the StatusID based on the ID
	initializers.DB.Model(&models.User{}).Where("email = ?", user.Email).Update("access_level", user.AccessLevel)

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User access level updated successfully",
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
	return c.Status(fiber.StatusOK).JSON(packs)
}

func SendPasswordResetCode(c *fiber.Ctx) error {

	email := c.Get("email")

	// Get the user with the given email
	var user models.User
	initializers.DB.Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.CreateUserAlreadyExistsException("User was not found with email: " + user.Email))
	}

	resetCode := utils.RandomString(NUM_DIGITS_RESET_CODE)
	var resetPasswordCode models.ResetPasswordCode
	resetPasswordCode.Code = resetCode
	resetPasswordCode.User_id = user.ID

	initializers.DB.Save(&resetPasswordCode)

	var body = fmt.Sprintf(BODY_RESET_PASSWORD, resetCode)
	utils.SendEmail([]string{user.Email}, SUBJECT_RESET_PASSWORD, body)

	return c.SendStatus(fiber.StatusOK)
}

func CheckResetCode(c *fiber.Ctx) error {

	code := c.Get("code")

	if code == "" {
		c.SendStatus(fiber.StatusBadRequest)
	}

	var resetPasswordCode models.ResetPasswordCode
	initializers.DB.Where("code = ?", code).Find(&resetPasswordCode)

	if resetPasswordCode.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Code is invalid",
		})
	}

	if time.Now().UnixNano()-resetPasswordCode.CreatedAt.UnixNano() > 3600000000000 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Code is invalid",
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

type passwordChangeInput struct {
	Password      string
	PasswordAgain string
}

func ResetPassword(c *fiber.Ctx) error {

	var passwordChangeInput = new(passwordChangeInput)

	if err := c.BodyParser(passwordChangeInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.CreateInvalidInputException("Bad input format"))
	}

	if !doTwoPasswordsMatch(passwordChangeInput.Password, passwordChangeInput.PasswordAgain) {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.CreateInvalidInputException("Passwords do not match"))
	}

	email := c.Get("email")

	var user models.User
	initializers.DB.Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(exceptions.CreateUserAlreadyExistsException("User was not found with email: " + user.Email))
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordChangeInput.Password), 14)
	user.Password = string(hashedPassword)
	initializers.DB.Save(&user)

	return c.SendStatus(fiber.StatusOK)
}

func doTwoPasswordsMatch(pw1 string, pw2 string) bool {
	return pw1 == pw2
}

func generateJwtToken(user models.User) string {
	claims := jwt.MapClaims{
		"access_level": user.AccessLevel,
		"user_id":      user.ID,
		"exp":          time.Now().Add(time.Hour * 24).Unix(), // Token expiration time set to 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(middleware.SecretKey)

	if err != nil {
		return err.Error()
	}

	return signedToken
}

func userIsAllowedToDelete(user models.User, tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(middleware.SecretKey), nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)

	tokenUserId := claims["user_id"]
	var deleterUser models.User
	initializers.DB.Where("id = ?", tokenUserId).Find(&deleterUser)

	return (tokenUserId != user.ID && deleterUser.AccessLevel == 3) || (deleterUser.ID == user.ID)
}
