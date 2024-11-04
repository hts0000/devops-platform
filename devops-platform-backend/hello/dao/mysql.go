package dao

import (
	"shared/model"

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

func (m *MySQL) CreateSayHelloRecord(r *model.SayHelloRecord) error {
	return m.db.Create(r).Error
}

func (m *MySQL) DeleteSayHelloRecord(r *model.SayHelloRecord) error {
	return m.db.Delete(r).Error
}
