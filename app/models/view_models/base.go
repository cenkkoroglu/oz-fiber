package view_models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	Id        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
