package model

import "gorm.io/gorm"

type BizLine struct {
	gorm.Model
	Name          string `gorm:"column:business_line_name"`
	ResponsibleID uint   `gorm:"column:responsible_id"`
	Description   string `gorm:"column:description"`
}

func (b *BizLine) TableName() string {
	return "business_line"
}
