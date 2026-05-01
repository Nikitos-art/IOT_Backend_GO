package services

import (
    "github.com/Nikitos-art/go-iot-backend/pkg/config"
    "github.com/Nikitos-art/go-iot-backend/pkg/models"
    "github.com/Nikitos-art/go-iot-backend/pkg/utils"
)

func CreateDevice(device *models.Device) error {

	apiKey, err := utils.GenerateAPIKey()
	if err != nil {
		return err
	}

	device.APIKey = apiKey

	result := config.GetDB().Create(device)
	return result.Error
}

func GetDevicesByUser(userID uint) ([]models.Device, error) {
    var devices []models.Device

    result := config.GetDB().
        Where("user_id = ?", userID).
        Find(&devices)

    return devices, result.Error
}


func GetDeviceByID(id uint) (*models.Device, error) {
    var device models.Device

    result := config.GetDB().
        Where("id = ?", id).
        First(&device)

    return &device, result.Error
}

func GetDeviceByAPIKey(key string) (*models.Device, error) {
	var device models.Device

	result := config.GetDB().
		Where("api_key = ?", key).
		First(&device)

	return &device, result.Error
}

func CreateDeviceData(data *models.DeviceData) error {
	result := config.GetDB().Create(data)
	return result.Error
}

func DeleteDevice(id uint) error {
    result := config.GetDB().
        Where("id = ?", id).
        Delete(&models.Device{})

    return result.Error
}