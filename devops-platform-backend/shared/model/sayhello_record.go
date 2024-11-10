package model

import "gorm.io/gorm"

type SayHelloRecord struct {
	gorm.Model
	Name string `gorm:"column:name"`
}

func (r *SayHelloRecord) TableName() string {
	return "sayhello_record"
}
