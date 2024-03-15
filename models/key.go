package models

import (
	"gorm.io/gorm"
	"time"
)

type Key struct {
	gorm.Model
	ValidFor time.Time
	User     User
	ApiKey   string
}
