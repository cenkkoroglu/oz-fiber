package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(*gorm.DB) (err error) {
	base.Id = uuid.New()
	return
}
