package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/config"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/controllers"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/middleware"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/models"
)


func setupFullFlowRouter() *mux.Router {

	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Device{})
	config.SetDB(db)

	r := mux.NewRouter()

	// DEVICE CREATE (protected by JWT bypass)
	r.Handle("/devices", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateDevice))).Methods("POST")

	// INGEST (API KEY AUTH - we simulate later)
	r.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	return r
}


func TestFullFlow(t *testing.T) {

	router := setupFullFlowRouter()

	// ----------------------------
	// STEP 1: CREATE DEVICE (JWT)
	// ----------------------------
	deviceBody := bytes.NewBuffer([]byte(`{"name":"sensor-1"}`))

	req1 := httptest.NewRequest("POST", "/devices", deviceBody)
	req1.Header.Set("Authorization", "Bearer test-token")

	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	if w1.Code != http.StatusCreated {
		t.Fatalf("device creation failed: %d", w1.Code)
	}

	// parse response
	var deviceResp map[string]interface{}
	json.NewDecoder(w1.Body).Decode(&deviceResp)

	apiKey := deviceResp["api_key"]
	if apiKey == nil {
		t.Fatal("expected api_key to be generated")
	}

	// ----------------------------
	// STEP 2: INGEST (API KEY)
	// ----------------------------
	payload := bytes.NewBuffer([]byte(`{"temperature": 22.5}`))

	req2 := httptest.NewRequest("POST", "/ingest", payload)
	req2.Header.Set("X-API-Key", apiKey.(string))

	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Fatalf("ingest failed: %d", w2.Code)
	}
}