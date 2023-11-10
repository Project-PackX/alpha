package controllers

import (
	"database/sql"
	"fmt"

	"github.com/Project-PackX/backend/enums"
	"github.com/Project-PackX/backend/initializers"
	"github.com/Project-PackX/backend/models"
	"github.com/Project-PackX/backend/utils"
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

var SUBJECT_PACKAGE_STATUS_MODIFIED = "The status of your package has changed"
var BODY_PACKAGE_STATUS_MODIFIED = `
	<span><h4 style="color:black;">Dear Customer, your package status has changed</h4></span>
	<p style="color:black;">Status: <em>%s</em><br></p>
	<p>Track ID: <em>%s</em></p>
	<p>Sincerely,<br>PackX</br></p>
`

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

// Get all package info wchich belong to the specific courier id based on URL
func GetPackagesUnderCourier(c *fiber.Ctx) error {

	// Getting the {cid} from URL
	cid := c.Params("id")

	// Getting the packs from the desired courier
	var packs []models.Package
	initializers.DB.Find(&packs, "courier_id = ?", cid)

	var stats []string
	for i := 0; i < len(packs); i++ {

		// Getting the package status code from the packagestatus table
		var statusindex models.PackageStatus
		initializers.DB.Find(&statusindex, "package_id = ?", packs[i].ID)

		// Getting the right status row in the status table
		var statusname models.Status
		initializers.DB.Find(&statusname, "id = ?", statusindex.Status_id)

		stats = append(stats, statusname.Name)
	}

	// Sending back the list of packages with every information
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"packages": packs,
		"statuses": stats,
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
	candidatePackage := new(models.Package)
	if err := c.BodyParser(candidatePackage); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Wrong format",
		})
	}

	var senderLocker *models.Locker
	var destinationLocker *models.Locker

	initializers.DB.Find(&senderLocker, "ID = ?", candidatePackage.SenderLockerId)
	initializers.DB.Find(&destinationLocker, "ID = ?", candidatePackage.DestinationLockerId)

	var packagesToSenderLocker []models.PackageLocker
	var packagesToDestinationLocker []models.PackageLocker
	initializers.DB.Find(&packagesToSenderLocker, "locker_id = ?", candidatePackage.SenderLockerId)
	initializers.DB.Find(&packagesToDestinationLocker, "locker_id = ?", candidatePackage.DestinationLockerId)

	nPackagesSenderLocker := len(packagesToSenderLocker)
	nPackagesDestinationLocker := len(packagesToDestinationLocker)

	if senderLocker.Capacity <= uint(nPackagesSenderLocker) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Sender locker's capacity is full",
		})
	}

	if destinationLocker.Capacity <= uint(nPackagesDestinationLocker) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": "Destination locker's capacity is full",
		})
	}

	// Get the size of a package
	size := candidatePackage.Size
	if size == "small" {
		candidatePackage.Size = enums.Sizes.Small
	} else if size == "medium" {
		candidatePackage.Size = enums.Sizes.Medium
	} else {
		candidatePackage.Size = enums.Sizes.Large
	}

	// Generate a random 6 digit number for the package code
	pCode := utils.RandomString(6)
	candidatePackage.Code = pCode

	// Generate TrackID
	trackId := utils.RandomString(10)
	candidatePackage.TrackID = trackId

	// Calculate CO2 savings
	candidatePackage.Co2 = utils.CalculateEmissionDifference(utils.CalculateDistance(senderLocker.Latitude, senderLocker.Longitude, destinationLocker.Latitude, destinationLocker.Longitude))

	// Inserting the new package
	result := initializers.DB.Create(&candidatePackage)

	// Adding dispatch status
	packageStatus := new(models.PackageStatus)
	packageStatus.Package_id = candidatePackage.ID
	packageStatus.Status_id = 1
	saveResult := initializers.DB.Create(&packageStatus)

	// Error handling
	if result.Error != nil || saveResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Something went wrong when creating the package",
		})
	}

	var packageToSenderLocker models.PackageLocker
	packageToSenderLocker.Package_id = candidatePackage.ID
	packageToSenderLocker.Locker_id = candidatePackage.SenderLockerId

	var packageToDestinationLocker models.PackageLocker
	packageToDestinationLocker.Package_id = candidatePackage.ID
	packageToDestinationLocker.Locker_id = candidatePackage.DestinationLockerId

	initializers.DB.Save(packageToSenderLocker)
	initializers.DB.Save(packageToDestinationLocker)

	var sender *models.User

	initializers.DB.Find(&sender, "ID = ?", candidatePackage.UserID)

	var body = fmt.Sprintf(BODY_ADD_PACKAGE, senderLocker.City+", "+senderLocker.Address, destinationLocker.City+", "+destinationLocker.Address, candidatePackage.TrackID)
	utils.SendEmail([]string{candidatePackage.ReceiverEmail, sender.Email}, SUBJECT_ADD_PACKAGE, body)

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(candidatePackage)
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
			"Message": "Error during the deletion of the package",
		})
	}

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Package deleted successfully",
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

	if err != nil {
		if err == sql.ErrNoRows {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "There is no package with such ID",
			})
		} else {
			c.SendStatus(fiber.StatusInternalServerError)
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
	newPackageStatus := new(models.PackageStatus)

	// Check error
	if err := c.BodyParser(newPackageStatus); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Update the StatusID based on the ID
	initializers.DB.Model(&models.PackageStatus{}).Where("package_id = ?", newPackageStatus.Package_id).Update("status_id", newPackageStatus.Status_id)

	var csomag *models.Package
	var status *models.Status
	initializers.DB.Find(&csomag, "ID = ?", newPackageStatus.Package_id)
	initializers.DB.Find(&status, "ID = ?", newPackageStatus.Status_id)

	var body = fmt.Sprintf(BODY_PACKAGE_STATUS_MODIFIED, status.Name, csomag.TrackID)
	utils.SendEmail([]string{csomag.ReceiverEmail}, SUBJECT_PACKAGE_STATUS_MODIFIED, body)

	if status.Name == enums.Statuses.Warehouse {
		initializers.DB.Where("package_id = ? AND locker_id = ?", newPackageStatus.Package_id, csomag.SenderLockerId).Delete(&models.PackageLocker{})
	}

	if status.Name == enums.Statuses.Delivered {
		initializers.DB.Where("package_id = ? AND locker_id = ?", newPackageStatus.Package_id, csomag.DestinationLockerId).Delete(&models.PackageLocker{})
	}

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Package status updated successfully",
	})
}

// Change package Status '+1'
func ChangeStatusUp(c *fiber.Ctx) error {

	id := c.Params("id")

	// Search for the package with the desired {id}
	var ps *models.PackageStatus
	err := initializers.DB.Where("package_id = ?", id).First(&ps).Error

	// Check error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get current status id
	statID := ps.Status_id

	// Only update if not delivered yet
	newStatID := statID
	if ps.Status_id < 5 {
		newStatID = newStatID + 1
	}

	// Update the StatusID
	initializers.DB.Model(&models.PackageStatus{}).Where("package_id = ?", ps.Package_id).Update("status_id", newStatID)

	var csomag *models.Package
	var status *models.Status
	initializers.DB.Find(&csomag, "ID = ?", ps.Package_id)
	initializers.DB.Find(&status, "ID = ?", ps.Status_id)

	var body = fmt.Sprintf(BODY_PACKAGE_STATUS_MODIFIED, status.Name, csomag.TrackID)
	utils.SendEmail([]string{csomag.ReceiverEmail}, SUBJECT_PACKAGE_STATUS_MODIFIED, body)

	if status.Name == enums.Statuses.Warehouse {
		initializers.DB.Where("package_id = ? AND locker_id = ?", ps.Package_id, csomag.SenderLockerId).Delete(&models.PackageLocker{})
	}

	if status.Name == enums.Statuses.Delivered {
		initializers.DB.Where("package_id = ? AND locker_id = ?", ps.Package_id, csomag.DestinationLockerId).Delete(&models.PackageLocker{})
	}

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(status)
}

func MakeCanceled(c *fiber.Ctx) error {
	id := c.Params("id")

	initializers.DB.Model(&models.Package{}).Where("package_id = ?", id).Update("status_id", 6)

	var csomag *models.Package
	initializers.DB.Where("package_id = ?", id).Find(&csomag)

	initializers.DB.Where("package_id = ? AND locker_id = ?", id, csomag.DestinationLockerId).Delete(&models.PackageLocker{})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Package successfully canceled",
	})
}

func GetAllAndPerPackageEmission(c *fiber.Ctx) error {

	var packages []models.Package
	initializers.DB.Find(&packages)

	all := 0.0
	for _, _package := range packages {
		all += _package.Co2
	}

	emissionPerPackage := all / float64(len(packages))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"All":                all,
		"EmissionPerPackage": emissionPerPackage,
	})
}
