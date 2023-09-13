package controllers

import (
	"PackX/initializers"
	"PackX/models"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

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
			"Message": "Hibás kérés",
		})
	}

	// Inserting the new package
	result := initializers.DB.Create(&csomag)

	// Error handling
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Hiba történt a csomag létrehozása közben",
		})
	}

	// Return as OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Csomag sikeresen hozzáadva",
	})
}

// Remove package with input json {id}
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
	id := c.Params("id")

	// Search for the package with the desired {id}
	var packageData *models.Package
	err := initializers.DB.Where("id = ?", id).First(&packageData).Error

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
	c.Status(fiber.StatusOK).JSON(packageData)
	return nil
}
