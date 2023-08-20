package initializers

import (
	"PackX/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// A programban szereplő adatbázis definiálása
var DB *gorm.DB

// Adatbázishot csatlakozás
func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Println("Nem sikerült kapcsolódni az adatbázishoz")
	}
}

// Tesztadatok miatt kell csak, hogy egységes legyen mindenkinek
func DropTables() {
	DB.Exec("DROP TABLE IF EXISTS public.users;")
	DB.Exec("DROP TABLE IF EXISTS public.packages;")
	DB.Exec("DROP TABLE IF EXISTS public.statuses;")
	DB.Exec("DROP TABLE IF EXISTS public.packagestatuses;")
	DB.Exec("DROP TABLE IF EXISTS public.couriers;")
}

// Automigrálás adatbázisból Go struct-okba
func SyncDB() {
	DB.AutoMigrate(&models.Package{})
	DB.AutoMigrate(&models.Courier{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Status{})
	DB.AutoMigrate(&models.PackageStatus{})
}

func GenerateTestEntries() {

	// Felhasználók

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

	// Csomagok

	csomag1 := models.Package{
		UserID:             2,
		DestinationAddress: "Szikra utca 26. II/28",
		Content:            "Cipő",
		Price:              37990,
		Note:               "Utánvét",
		CourierID:          1,
	}
	DB.Create(&csomag1)

	csomag2 := models.Package{
		UserID:             1,
		DestinationAddress: "Szendrei Lipót sétány 40.",
		Content:            "Laptop",
		Price:              225000,
		Note:               "Javítás",
		CourierID:          1,
	}
	DB.Create(&csomag2)

	csomag3 := models.Package{
		UserID:             2,
		DestinationAddress: "Nagy István út 81.",
		Content:            "Kabát",
		Price:              17490,
		Note:               "-",
		CourierID:          2,
	}
	DB.Create(&csomag3)

	csomag4 := models.Package{
		UserID:             3,
		DestinationAddress: "Kő utca 5.",
		Content:            "25W töltőfej",
		Price:              3989,
		Note:               "Cserekészülék",
		CourierID:          2,
	}
	DB.Create(&csomag4)

	csomag5 := models.Package{
		UserID:             1,
		DestinationAddress: "Rózsavölgy körút 67/B",
		Content:            "Bútor",
		Price:              55990,
		Note:               "-",
		CourierID:          1,
	}
	DB.Create(&csomag5)

	// Lehetséges csomag státuszok

	statusz1 := models.Status{
		Id:   1,
		Name: "Feladva",
	}
	DB.Create(&statusz1)

	statusz2 := models.Status{
		Id:   2,
		Name: "Átvéve",
	}
	DB.Create(&statusz2)

	statusz3 := models.Status{
		Id:   3,
		Name: "Raktárban",
	}
	DB.Create(&statusz3)

	statusz4 := models.Status{
		Id:   4,
		Name: "Szállítás alatt",
	}
	DB.Create(&statusz4)

	statusz5 := models.Status{
		Id:   5,
		Name: "Kézbesítve",
	}
	DB.Create(&statusz5)

	statusz6 := models.Status{
		Id:   6,
		Name: "Törölve",
	}
	DB.Create(&statusz6)

	// A bevitt csomagok státuszai

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

	// Futárok

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

}
