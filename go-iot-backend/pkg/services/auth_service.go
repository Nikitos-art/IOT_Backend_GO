package services

import (
    "errors"

    "golang.org/x/crypto/bcrypt"
    "github.com/jinzhu/gorm"

    "github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/models"
    "github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/utils"
)

type AuthService struct {
    db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
    return &AuthService{db: db}
}

func (s *AuthService) Register(email, password string) (*models.User, error) {
    // check if user exists
    var existing models.User
    if err := s.db.Where("email = ?", email).First(&existing).Error; err == nil {
        return nil, errors.New("user already exists")
    }

    // hash password
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := models.User{
        Email:    email,
        Password: string(hashed),
    }

    if err := s.db.Create(&user).Error; err != nil {
        return nil, err
    }

    return &user, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	var user models.User

	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}

	// compare password (assuming you already hash it in register)
	if err := utils.CheckPassword(user.Password, password); err != nil {
		return "", err
	}

	// create JWT token
	token, err := utils.CreateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// func (s *AuthService) Login(email, password string) (*models.User, error) {
//     var user models.User

//     if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
//         return nil, errors.New("invalid credentials")
//     }

//     if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
//         return nil, errors.New("invalid credentials")
//     }

//     return &user, nil
// }