package initializers

import (
	"PackX/enums"
	"PackX/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Defining the application databse structure with gorm
var DB *gorm.DB

// Connectiing to the database based on the environment variables
func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Error handling
	if err != nil {
		fmt.Println("Nem sikerült kapcsolódni az adatbázishoz")
	}
}

// FOR TESTING PURPOSES
func DropTables() {
	DB.Exec("DROP TABLE IF EXISTS public.users;")
	DB.Exec("DROP TABLE IF EXISTS public.packages;")
	DB.Exec("DROP TABLE IF EXISTS public.statuses;")
	DB.Exec("DROP TABLE IF EXISTS public.packagestatuses;")
	DB.Exec("DROP TABLE IF EXISTS public.couriers;")
	DB.Exec("DROP TABLE IF EXISTS public.lockers;")
	DB.Exec("DROP TABLE IF EXISTS public.lockergroups;")
}

// Migrating the DB tables into Go models
func SyncDB() {
	DB.AutoMigrate(&models.Package{})
	DB.AutoMigrate(&models.Courier{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Status{})
	DB.AutoMigrate(&models.PackageStatus{})
	DB.AutoMigrate(&models.Locker{})
	DB.AutoMigrate(&models.LockerGroup{})
}

// Generate test datas
func GenerateTestEntries() {

	// Users

	felh1 := models.User{
		Name:    "Kovács Bea",
		Address: "Liliom utca 4.",
		Phone:   "+36201956673",
		Email:   "k.bea@mail.com",
	}
	DB.Create(&felh1)

	felh2 := models.User{
		Name:    "Szalma Géza",
		Address: "Egressy körút 58.",
		Phone:   "+36605385438",
		Email:   "szalmag@mail.com",
	}
	DB.Create(&felh2)

	felh3 := models.User{
		Name:    "Veres Péter",
		Address: "Malom út 12.",
		Phone:   "+36504098931",
		Email:   "vrsptr@mail.com",
	}
	DB.Create(&felh3)

	// Packages

	csomag1 := models.Package{
		UserID:    2,
		Size:      enums.Sizes.Medium,
		Price:     37990,
		Note:      "Utánvét",
		CourierID: 1,
	}
	DB.Create(&csomag1)

	csomag2 := models.Package{
		UserID:    1,
		Size:      enums.Sizes.Small,
		Price:     225000,
		Note:      "Javítás",
		CourierID: 1,
	}
	DB.Create(&csomag2)

	csomag3 := models.Package{
		UserID:    2,
		Size:      enums.Sizes.Medium,
		Price:     17490,
		Note:      "-",
		CourierID: 2,
	}
	DB.Create(&csomag3)

	csomag4 := models.Package{
		UserID:    3,
		Size:      enums.Sizes.Small,
		Price:     3989,
		Note:      "Cserekészülék",
		CourierID: 2,
	}
	DB.Create(&csomag4)

	csomag5 := models.Package{
		UserID:    1,
		Size:      enums.Sizes.Large,
		Price:     55990,
		Note:      "-",
		CourierID: 1,
	}
	DB.Create(&csomag5)

	// Possible package statuses

	statusz1 := models.Status{
		Id:   1,
		Name: enums.Statuses.Dispatch,
	}
	DB.Create(&statusz1)

	statusz2 := models.Status{
		Id:   2,
		Name: enums.Statuses.Transit,
	}
	DB.Create(&statusz2)

	statusz3 := models.Status{
		Id:   3,
		Name: enums.Statuses.Warehouse,
	}
	DB.Create(&statusz3)

	statusz4 := models.Status{
		Id:   4,
		Name: enums.Statuses.Delivery,
	}
	DB.Create(&statusz4)

	statusz5 := models.Status{
		Id:   5,
		Name: enums.Statuses.Delivered,
	}
	DB.Create(&statusz5)

	statusz6 := models.Status{
		Id:   6,
		Name: enums.Statuses.Cancelled,
	}
	DB.Create(&statusz6)

	// Package statuses

	csomagstatusz1 := models.PackageStatus{
		Package_id: 1,
		Status_id:  1,
	}
	DB.Create(&csomagstatusz1)

	csomagstatusz2 := models.PackageStatus{
		Package_id: 2,
		Status_id:  4,
	}
	DB.Create(&csomagstatusz2)

	csomagstatusz3 := models.PackageStatus{
		Package_id: 3,
		Status_id:  3,
	}
	DB.Create(&csomagstatusz3)

	csomagstatusz4 := models.PackageStatus{
		Package_id: 4,
		Status_id:  4,
	}
	DB.Create(&csomagstatusz4)

	csomagstatusz5 := models.PackageStatus{
		Package_id: 5,
		Status_id:  2,
	}
	DB.Create(&csomagstatusz5)

	// Couriers

	futar1 := models.Courier{
		Name:  "Kiss Bendegúz",
		Phone: "+36403437791",
	}
	DB.Create(&futar1)

	futar2 := models.Courier{
		Name:  "Némedi Emma",
		Phone: "+36301984673",
	}
	DB.Create(&futar2)

	// Lockers

	locker1 := models.Locker{
		Address:       "Szent István út 23.",
		LockerGroupID: 1,
	}
	DB.Create(&locker1)

	locker2 := models.Locker{
		Address:       "Kiss Ernő utca 5.",
		LockerGroupID: 1,
	}
	DB.Create(&locker2)

	locker3 := models.Locker{
		Address:       "Lomnic utca 30.",
		LockerGroupID: 1,
	}
	DB.Create(&locker3)

	locker4 := models.Locker{
		Address:       "Paragvári utca 74.",
		LockerGroupID: 2,
	}
	DB.Create(&locker4)

	locker5 := models.Locker{
		Address:       "Gömör utca 3.",
		LockerGroupID: 2,
	}
	DB.Create(&locker5)

	locker6 := models.Locker{
		Address:       "Éhen Gyula tér 3.",
		LockerGroupID: 2,
	}
	DB.Create(&locker6)

	locker7 := models.Locker{
		Address:       "Sziget utca 7.",
		LockerGroupID: 2,
	}
	DB.Create(&locker7)

	// Lockergroups

	lgroup1 := models.LockerGroup{
		ID:   1,
		City: "Győr",
	}
	DB.Create(&lgroup1)

	lgroup2 := models.LockerGroup{
		ID:   2,
		City: "Szombathely",
	}
	DB.Create(&lgroup2)
}
