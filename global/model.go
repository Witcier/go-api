package global

import (
	"time"

	"gorm.io/gorm"
)

type API_MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
