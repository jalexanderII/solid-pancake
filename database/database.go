package database

import (
	"github.com/hashicorp/go-hclog"

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
	db.AutoMigrate()

	Database = DbInstance{Db: db}
}