package controllers


import (
	"encoding/json"
	"net/http"

	"github.com/Nikitos-art/go-bookstore/pkg/models"
	"github.com/Nikitos-art/go-bookstore/pkg/services"
)


func IngestData(w http.ResponseWriter, r *http.Request) {

	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		http.Error(w, "missing api key", http.StatusUnauthorized)
		return
	}

	device, err := services.GetDeviceByAPIKey(apiKey)
	if err != nil {
		http.Error(w, "invalid api key", http.StatusUnauthorized)
		return
	}

	var data models.DeviceData
	json.NewDecoder(r.Body).Decode(&data)

	data.DeviceID = device.ID

	err = services.CreateDeviceData(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ingested"))
}