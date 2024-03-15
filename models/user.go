package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	RevokeTime time.Time
	Pkgs       []Pkg
}
