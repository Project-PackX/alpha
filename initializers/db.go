package initializers

import (
	"fmt"
	"os"
	"time"

	"github.com/Project-PackX/backend/enums"
	"github.com/Project-PackX/backend/models"
	"github.com/Project-PackX/backend/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Defining the application databse structure with gorm
var DB *gorm.DB

// Connectiing to the database based on the environment variables
func ConnectToDatabase() {
	logger := utils.Logger

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), 5432)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	// Error handling
	if err != nil {
		logger.Error("Couldn't connect to the database")
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
	DB.Exec("DROP TABLE IF EXISTS public.packageslockers")
	DB.Exec("DROP TABLE IF EXISTS public.reset_password_code")
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
	DB.AutoMigrate(&models.PackageLocker{})
	DB.AutoMigrate(&models.ResetPasswordCode{})
}

// Generate test datas
func GenerateTestEntries() {
	type UserTestData struct {
		name        string
		address     string
		phone       string
		email       string
		accessLevel uint
	}

	testUsers := []UserTestData{
		{
			name:        "Kovács Bea",
			address:     "Liliom utca 4.",
			phone:       "+36201956673",
			email:       "k.bea@mail.com",
			accessLevel: enums.AccessLevel.Normal,
		},
		{
			name:        "Szalma Géza",
			address:     "Egressy körút 58.",
			phone:       "+36605385438",
			email:       "szalmag@mail.com",
			accessLevel: enums.AccessLevel.Normal,
		},
		{
			name:        "Veres Péter",
			address:     "Malom út 12.",
			phone:       "+36504098931",
			email:       "vrsptr@mail.com",
			accessLevel: enums.AccessLevel.Admin,
		},
	}

	for _, u := range testUsers {
		userModel := models.User{
			Name:        u.name,
			Address:     u.address,
			Phone:       u.phone,
			Email:       u.email,
			AccessLevel: u.accessLevel,
		}

		DB.Create(&userModel)
	}

	// Packages
	type PackageTestData struct {
		userId       uint
		size         string
		price        float64
		note         string
		courierId    uint
		deliveryDate time.Time
	}

	testPackages := []PackageTestData{
		{
			userId:       2,
			size:         enums.Sizes.Medium,
			price:        37990,
			note:         "Utánvét",
			courierId:    1,
			deliveryDate: time.Now().Add(time.Hour * 24 * 5),
		},
		{
			userId:       1,
			size:         enums.Sizes.Small,
			price:        225000,
			note:         "Javítás",
			courierId:    1,
			deliveryDate: time.Now().Add(time.Hour * 24 * 5),
		},
		{
			userId:       2,
			size:         enums.Sizes.Medium,
			price:        17490,
			note:         "-",
			courierId:    2,
			deliveryDate: time.Now().Add(time.Hour * 24 * 5),
		},
		{
			userId:       3,
			size:         enums.Sizes.Small,
			price:        3989,
			note:         "Cserekészülék",
			courierId:    2,
			deliveryDate: time.Now().Add(time.Hour * 24 * 5)},
		{
			userId:       1,
			size:         enums.Sizes.Large,
			price:        55990,
			note:         "-",
			courierId:    1,
			deliveryDate: time.Now().Add(time.Hour * 24 * 5),
		},
		{
			userId:       1,
			size:         enums.Sizes.Small,
			price:        3490,
			note:         "-",
			courierId:    2,
			deliveryDate: time.Now().Add(time.Hour * 24 * 5),
		},
	}

	for _, p := range testPackages {
		packageModel := models.Package{
			UserID:       p.userId,
			Size:         p.size,
			Price:        p.price,
			Note:         p.note,
			CourierID:    p.courierId,
			DeliveryDate: p.deliveryDate,
		}

		DB.Create(&packageModel)
	}

	// Possible package statuses
	type StatusTestData struct {
		id   uint
		name string
	}

	testStatuses := []StatusTestData{
		{
			id:   1,
			name: enums.Statuses.Dispatch,
		},
		{
			id:   2,
			name: enums.Statuses.Transit,
		},
		{
			id:   3,
			name: enums.Statuses.Warehouse,
		},
		{
			id:   4,
			name: enums.Statuses.Delivery,
		},
		{
			id:   5,
			name: enums.Statuses.Delivered,
		},
		{
			id:   6,
			name: enums.Statuses.Canceled,
		},
	}

	for _, s := range testStatuses {
		statusModel := models.Status{
			Id:   s.id,
			Name: s.name,
		}

		DB.Create(&statusModel)
	}

	// Package statuses
	type PackageStatusTestData struct {
		packageId uint
		statusId  uint
	}

	testPackageStatues := []PackageStatusTestData{
		{
			packageId: 1,
			statusId:  1,
		},
		{
			packageId: 2,
			statusId:  4,
		},
		{
			packageId: 3,
			statusId:  3,
		},
		{
			packageId: 4,
			statusId:  5,
		},
		{
			packageId: 5,
			statusId:  2,
		},
	}

	for _, ps := range testPackageStatues {
		packageStatusModel := models.PackageStatus{
			Package_id: ps.packageId,
			Status_id:  ps.statusId,
		}

		DB.Create(&packageStatusModel)
	}

	// Couriers
	type CourierTestData struct {
		name  string
		phone string
	}

	testCouriers := []CourierTestData{
		{
			name:  "Kiss Bendegúz",
			phone: "+36403437791",
		},
		{
			name:  "Némedi Emma",
			phone: "+36301984673",
		},
	}

	for _, t := range testCouriers {
		courierModel := models.Courier{
			Name:  t.name,
			Phone: t.phone,
		}
		DB.Create(&courierModel)
	}

	// Lockers
	type LockerTestData struct {
		city      string
		address   string
		capacity  uint
		latitude  float64
		longitude float64
	}

	testLockers := []LockerTestData{
		{
			city:      "Győr",
			address:   "Szent István út 23.",
			capacity:  7,
			latitude:  47.683337112393474,
			longitude: 17.623955422088322,
		},
		{
			city:      "Győr",
			address:   "Kiss Ernő utca 5.",
			capacity:  5,
			latitude:  47.68834622211481,
			longitude: 17.623955422088322,
		},
		{
			city:      "Győr",
			address:   "Lomnic utca 30.",
			capacity:  5,
			latitude:  47.67035183513254,
			longitude: 17.63988174907635,
		},
		{
			city:      "Szombathely",
			address:   "Paragvári utca 74.",
			capacity:  5,
			latitude:  47.24350217822487,
			longitude: 17.623955422088322,
		},
		{
			city:      "Szombathely",
			address:   "Gömör utca 3.",
			capacity:  5,
			latitude:  47.22994517025941,
			longitude: 16.60908613742963,
		},
		{
			city:      "Szombathely",
			address:   "Éhen Gyula tér 3.",
			capacity:  10,
			latitude:  47.23677632684086,
			longitude: 16.631628691405695,
		},
		{
			city:      "Szombathely",
			address:   "Sziget utca 7.",
			capacity:  15,
			latitude:  47.23858035784418,
			longitude: 16.64677093558233,
		},
	}

	for _, l := range testLockers {
		lockerModel := models.Locker{
			City:      l.city,
			Address:   l.address,
			Capacity:  l.capacity,
			Latitude:  l.latitude,
			Longitude: l.longitude,
		}

		DB.Create(&lockerModel)
	}

	// Locker Groups
	type LockerGroupTestData struct {
		id   uint
		city string
	}

	testLockerGroups := []LockerGroupTestData{
		{
			id:   1,
			city: "Győr",
		},
		{
			id:   2,
			city: "Szombathely",
		},
	}

	for _, lg := range testLockerGroups {
		lockerGroupModel := models.LockerGroup{
			ID:   lg.id,
			City: lg.city,
		}

		DB.Create(&lockerGroupModel)
	}

	// Package Lockers
	type PackageLockersTestData struct {
		packageId uint
		lockerId  uint
	}

	testPackageLockers := []PackageLockersTestData{
		{
			packageId: 1,
			lockerId:  2,
		},
		{
			packageId: 2,
			lockerId:  1,
		},
		{
			packageId: 3,
			lockerId:  6,
		},
		{
			packageId: 4,
			lockerId:  1,
		},
		{
			packageId: 5,
			lockerId:  6,
		},
	}

	for _, pl := range testPackageLockers {
		packageLockerModel := models.PackageLocker{
			Package_id: pl.packageId,
			Locker_id:  pl.lockerId,
		}

		DB.Create(&packageLockerModel)
	}
}
