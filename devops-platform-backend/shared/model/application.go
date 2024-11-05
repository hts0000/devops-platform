package model

import "gorm.io/gorm"

type App struct {
	gorm.Model
	Name string `json:"name" column:"name"`
}

func (b *App) TableName() string {
	return "application"
}
