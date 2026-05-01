package models

import "time"

type DeviceData struct {
	ID        uint      `gorm:"primaryKey"`
	DeviceID  uint      `gorm:"index"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time
}