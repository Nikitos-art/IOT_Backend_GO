package services

import (
	"testing"

	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/config"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupTestDB() {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// create schema needed for test
	db.AutoMigrate(&models.Device{})

	// inject into global config so service uses this DB
	config.SetDB(db)
}

func TestCreateDevice(t *testing.T) {

	// arrange test DB
	setupTestDB()

	device := models.Device{
		Name:   "test-device",
		UserID: 1,
	}

	// act
	err := CreateDevice(&device)

	// assert error
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// assert API key generated
	if device.APIKey == "" {
		t.Errorf("expected API key to be generated")
	}
}