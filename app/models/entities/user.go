package entities

import (
	"github.com/cenkkoroglu/oz-fiber/app/models/view_models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	UserName string `gorm:"size:255;not null;unique"`
	Email    string `gorm:"size:100;not null;unique"`
	Password string `gorm:"size:255;not null;"`
}

func (u *User) HashPassword(password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	u.Password = string(bytes)
}

func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (u *User) ToUserVm() *view_models.User {
	return &view_models.User{
		Base: view_models.Base{
			Id:        u.Id,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: u.DeletedAt,
		},
		UserName: u.UserName,
		Email:    u.Email,
	}
}
