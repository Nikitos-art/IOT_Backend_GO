package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/models"
    "github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/services"
    "github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/middleware"
)

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device

	// decode request
	err := json.NewDecoder(r.Body).Decode(&device)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// get user from middleware
	userID := middleware.GetUserID(r)
	device.UserID = userID

	// create device
	err = services.CreateDevice(&device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 🔥 IMPORTANT: set status FIRST
	w.WriteHeader(http.StatusCreated)

	// response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      device.ID,
		"name":    device.Name,
		"api_key": device.APIKey,
	})
}

func GetDevicesByUser(w http.ResponseWriter, r *http.Request) {

	userID := middleware.GetUserID(r)

	devices, err := services.GetDevicesByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(devices)
}

func GetDeviceByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    device, err := services.GetDeviceByID(uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(device)
}


func DeleteDevice(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    err = services.DeleteDevice(uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}