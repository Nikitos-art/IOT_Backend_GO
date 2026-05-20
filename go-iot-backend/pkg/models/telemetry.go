package models

import "github.com/jinzhu/gorm"

type Telemetry struct {
	gorm.Model

	DeviceID uint `json:"device_id"`
	// Data     string  `json:"data"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`

	Raw string `json:"raw"`
}
