package entities

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (b *BaseEntity) New() *BaseEntity {
	return &BaseEntity{
		Id:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}
