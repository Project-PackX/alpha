package controllers

import (
	"PackX/enums"
	"PackX/initializers"
	"PackX/models"
	"PackX/utils"
	"database/sql"
	"fmt"
	"math/rand"

	"time"

	"github.com/gofiber/fiber/v2"
)

var SUBJECT_ADD_PACKAGE = "Your package has beent sent"
var BODY_ADD_PACKAGE = `
	<span><h4 style="color:black;">Dear Customer, your package is being delivered</h4></span>
	<p style="color:black;">From: %s<br>
	To: %s</p>
	<p>You can track your package with this track ID:<em>%s</em></p>
	<p>Sincerely,<br>PackX</br></p>
`

// Generate a random string (letters + numbers) with the given length
func randomString(length int) string {
	characters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = characters[rand.Intn(len(characters))]
	}

	return string(bytes)
}

// Just load a test html page
func PostsIndex(c *fiber.Ctx) error {
	return c.Render("packs/index", fiber.Map{})
}

// List all the users with their packages
func ListUsersWithPackages(c *fiber.Ctx) error {

	// Slice that will contain all information
	var users []models.User
	initializers.DB.Model(&models.User{}).Preload("Packages").Find(&users)

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message":  "Success",
		"Csomagok": users,
	})
}

// Return the package status based on the {id} in the URL
func GetPackageStatus(c *fiber.Ctx) error {

	// Database is needed in this function
	db := initializers.DB

	// Getting the {id} from the URL
	id := c.Params("id")

	// Getting the package status code from the packagestatus table
	var statusindex models.PackageStatus
	db.Find(&statusindex, "package_id = ?", id)

	// Getting the right status row in the status table
	var statusname models.Status
	db.Find(&statusname, "id = ?", statusindex.Status_id)

	// Returning the answer
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"PackageID":  statusindex.Package_id,
		"StatusName": statusname.Name,
	})

}

// List all packages
func ListPackages(c *fiber.Ctx) error {

	var csomagok []models.Package // Slice that will contain all packages

	// Execute 'SELECT * FROM public.csomagok' query
	result := initializers.DB.Find(&csomagok)

	// Error handling
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Hiba történt a lekérdezés közben",
		})
	}

	// Separately treated case where no record exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Message": "Nincs egy aktív csomag sem",
		})
	}

	// Returning the packages
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message":  "Success",
		"Csomagok": csomagok,
	})
}

// Insert new package to the database
func AddNewPackage(c *fiber.Ctx) error {

	// Creating the new package with the input json body
	csomag := new(models.Package)
	if err := c.BodyParser(csomag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Wrong format",
		})
	}

	// Get the size of a package
	meret := csomag.Size
	if meret == "small" {
		csomag.Size = enums.Sizes.Small
	} else if meret == "medium" {
		csomag.Size = enums.Sizes.Medium
	} else {
		csomag.Size = enums.Sizes.Large
	}

	// Generate a random 6 digit number for the package code
	pcode := randomString(6)
	csomag.Code = pcode

	// Generate delivery date
	ddate := time.Now()
	if csomag.Rapid {
		ddate = ddate.Add(time.Hour * 3 * 24) // Add 3 days
	} else {
		ddate = ddate.Add(time.Hour * 5 * 24) // Add 5 days
	}
	csomag.DeliveryDate = ddate

	// Generate TrackID
	trackid := randomString(10)
	csomag.TrackID = trackid

	// Inserting the new package
	result := initializers.DB.Create(&csomag)

	// Adding dispatch status
	csomagstatusz := new(models.PackageStatus)
	csomagstatusz.Package_id = csomag.ID
	csomagstatusz.Status_id = 1
	result2 := initializers.DB.Create(&csomagstatusz)

	// Error handling
	if result.Error != nil || result2.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Hiba történt a csomag létrehozása közben",
		})
	}

	var sender *models.User
	var senderLocker *models.Locker
	var destinationLocker *models.Locker

	initializers.DB.Find(&sender, "ID = ?", csomag.UserID)
	initializers.DB.Find(&senderLocker, "ID = ?", csomag.SenderLockerId)
	initializers.DB.Find(&destinationLocker, "ID = ?", csomag.DestinationLockerId)

	var body = fmt.Sprintf(BODY_ADD_PACKAGE, senderLocker.City+", "+senderLocker.Address, destinationLocker.City+", "+destinationLocker.Address, csomag.TrackID)
	utils.SendEmail([]string{csomag.ReceiverEmail, sender.Email}, SUBJECT_ADD_PACKAGE, body)

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(csomag)
}

// Remove package with URL input {id}
func DeletePackageByID(c *fiber.Ctx) error {

	//Getting the {id} from URL
	id := c.Params("id")

	// Removing the package based on the {id}
	result := initializers.DB.Delete(&models.Package{}, id)

	// Error handling
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Hiba történt a csomag törlése közben",
		})
	}

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Csomag sikeresen törölve",
	})
}

// Getting the details of the {id}. package
func ListPackageByID(c *fiber.Ctx) error {

	// Getting the {id} from URL
	id := c.Params("trackid")

	// Search for the package with the desired {id}
	var packageData *models.Package
	err := initializers.DB.Where("track_id = ?", id).First(&packageData).Error
	packID := packageData.ID

	// Search for Status
	// Getting the package status code from the packagestatus table
	var statusindex models.PackageStatus
	initializers.DB.Find(&statusindex, "package_id = ?", packID)

	// Getting the right status row in the status table
	var statusname models.Status
	initializers.DB.Find(&statusname, "id = ?", statusindex.Status_id)

	// Error handling
	if err != nil {
		// If there are no package with that {id}
		if err == sql.ErrNoRows {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "Nincs ilyen azonosítójú csomag",
			})
		} else { // Other error happened during the query
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Message": "Hiba történt a lekérdezés közben",
			})
		}
		return err
	}

	// Return as OK
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Data":   packageData,
		"Status": statusname.Name,
	})
	return nil
}

// Change a package status via input JSON (ID, NewStatusID)
func ChangeStatus(c *fiber.Ctx) error {

	// Making a package model with the given parameters
	newPackage := new(models.PackageStatus)

	// Check error
	if err := c.BodyParser(newPackage); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Hibás kérés",
		})
	}

	// Update the StatusID based on the ID
	initializers.DB.Model(&models.PackageStatus{}).Where("package_id = ?", newPackage.Package_id).Update("status_id", newPackage.Status_id)

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Package status updated successfully",
	})
}

func MakeCanceled(c *fiber.Ctx) error {
	id := c.Params("id")

	initializers.DB.Model(&models.PackageStatus{}).Where("package_id = ?", id).Update("status_id", 6)

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Package successfully canceled",
	})
}
