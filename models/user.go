package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username   string
	RevokeTime time.Time
	Pkgs       []Pkg
}
