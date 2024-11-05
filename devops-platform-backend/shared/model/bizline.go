package model

import "gorm.io/gorm"

type BizLine struct {
	gorm.Model
	Name          string `json:"name" column:"name"`
	ResponsibleID uint   `json:"responsible_id" column:"responsible_id"`
	Description   string `json:"description" column:"description"`
}

func (b *BizLine) TableName() string {
	return "business_line"
}
