package services

import (
	"github.com/cenkkoroglu/oz-fiber/app/models/entities"
	"github.com/cenkkoroglu/oz-fiber/app/models/request_models"
	"github.com/cenkkoroglu/oz-fiber/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"time"
)

type userService struct {
	DBConn *gorm.DB
}

func (u *userService) GetByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	if err := u.DBConn.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) Register(registerRequest *request_models.Register) (*entities.User, error) {
	user := entities.User{
		UserName: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}

	user.HashPassword(user.Password)

	err := u.DBConn.Create(&user).Error
	return &user, err
}

func (u *userService) GenerateToken(user *entities.User) (*string, error) {
	config := config.GetConfig()

	claims := jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Minute * time.Duration(config.JwtExpireMinute)).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}

type UserService interface {
	GenerateToken(user *entities.User) (*string, error)
	Register(registerRequest *request_models.Register) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
}

func NewUserService(_dbConn *gorm.DB) UserService {
	return &userService{
		DBConn: _dbConn,
	}
}
