package db

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type IPStatus string

const (
	Active    IPStatus = "active"
	Blocked   IPStatus = "blocked"
	Suspended IPStatus = "suspended"
)

type IP struct {
	ID            uuid.UUID  `gorm:"type:char(36);primary_key"`
	IPAddress     string     `gorm:"type:varchar(15);unique_index;not null"`
	Note          string     `gorm:"type:varchar(64);"`
	CreatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	LastUpdatedAt *time.Time `gorm:"default:NULL"`
	Status        IPStatus   `gorm:"type:enum('active', 'blocked', 'suspended);not null"`
}

func (ip *IP) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
