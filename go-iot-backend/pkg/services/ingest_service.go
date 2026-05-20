package services

import (
	"encoding/json"

	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/config"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/models"
)

func IngestData(apiKey string, payload map[string]interface{}) error {

	db := config.GetDB()

	var device models.Device

	if err := db.Where("api_key = ?", apiKey).First(&device).Error; err != nil {
		return err
	}

	var temperature float64
	var humidity float64

	if v, ok := payload["temperature"].(float64); ok {
		temperature = v
	}

	if v, ok := payload["humidity"].(float64); ok {
		humidity = v
	}

	rawBytes, _ := json.Marshal(payload)

	telemetry := models.Telemetry{
		DeviceID: device.ID,
		// Data:     data,
		Temperature: temperature,
		Humidity:    humidity,
		Raw:         string(rawBytes),
	}

	return db.Create(&telemetry).Error
}
