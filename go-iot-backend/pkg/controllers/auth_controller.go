package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/Nikitos-art/go-iot-backend/pkg/services"
)

type AuthController struct {
    authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
    return &AuthController{
        authService: authService,
    }
}

type RegisterRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest

    json.NewDecoder(r.Body).Decode(&req)

    user, err := c.authService.Register(req.Email, req.Password)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "id":    user.ID,
        "email": user.Email,
    })
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest

    json.NewDecoder(r.Body).Decode(&req)

    token, err := c.authService.Login(req.Email, req.Password)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "token": token,
    })
}