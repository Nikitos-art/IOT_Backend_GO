package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/config"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/middleware"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/models"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRouter() *mux.Router {
	// in-memory DB
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Device{})

	config.SetDB(db)

	// router setup
	r := mux.NewRouter()

	// protected route
	r.Handle("/devices", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := middleware.GetUserID(r)

		dev := models.Device{
			Name:   "test-device",
			UserID: userID,
		}

		services.CreateDevice(&dev)

		w.WriteHeader(http.StatusCreated)
	}))).Methods("POST")

	return r
}


func TestCreateDeviceHTTP(t *testing.T) {

	router := setupRouter()

	body := bytes.NewBuffer([]byte(`{"name":"sensor-1"}`))

	req := httptest.NewRequest("POST", "/devices", body)

	// fake JWT (must match your ParseToken expectations OR use bypass logic later)
	req.Header.Set("Authorization", "Bearer test-token")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
}