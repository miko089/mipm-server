package models

import (
	"gorm.io/gorm"
)

type State uint

const (
	Actual State = iota
	Deprecated
	Archive
)

type Pkg struct {
	gorm.Model
	Author   User
	Versions []PkgVersion
	CurState State
}
