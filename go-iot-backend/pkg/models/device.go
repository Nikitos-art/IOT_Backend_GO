package models

import (
    "github.com/jinzhu/gorm"
)


type Device struct {
	gorm.Model

	Name   string `json:"name" validate:"required"`
	UserID uint   `json:"user_id"`

	APIKey string `json:"api_key"`
}