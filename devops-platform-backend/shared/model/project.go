package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string `json:"name" column:"name"`
}

func (b *Project) TableName() string {
	return "project"
}
