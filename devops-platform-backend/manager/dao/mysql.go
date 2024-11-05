package dao

import (
	"gorm.io/gorm"
)

type MySQL struct {
	db *gorm.DB
}

func NewMySQL(db *gorm.DB) *MySQL {
	return &MySQL{
		db: db,
	}
}
