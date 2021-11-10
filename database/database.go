package database

import (
	"github.com/hashicorp/go-hclog"
	ApplicationM "github.com/jalexanderII/solid-pancake/services/application/models"
	LifeCycleM "github.com/jalexanderII/solid-pancake/services/lifecycle/models"
	RealEstateM "github.com/jalexanderII/solid-pancake/services/realestate/models"
	UserM "github.com/jalexanderII/solid-pancake/services/users/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DbInstance is a struct that holds database pointer
type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dbLogger := hclog.Default()

	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		dbLogger.Error("failed to connect database", "error", err)
	}
	dbLogger.Info("Connection Opened to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	// Migrate the schema
	dbLogger.Info("Running Migrations")
	db.AutoMigrate(
		// Real estate models
		&RealEstateM.Apartment{}, &RealEstateM.Building{}, &RealEstateM.Realtor{},
		// Application models
		&ApplicationM.ApplicantFormRequest{}, &ApplicationM.ApplicantFormResponse{},
		// User models
		&UserM.User{},
		// LifeCycle models
		&LifeCycleM.RentalDetails{}, &LifeCycleM.PaymentRequest{}, &LifeCycleM.PaymentResponse{}, &LifeCycleM.PaymentConfirmation{},
	)

	Database = DbInstance{Db: db}
}
