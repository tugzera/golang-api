package models

import (
	"time"

	"gorm.io/gorm"
)

type ArmorModel struct {
	gorm.Model

	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
