package model

import "gorm.io/gorm"

type SayHelloRecord struct {
	gorm.Model
	Name string `json:"name" column:"name"`
}

func (r *SayHelloRecord) TableName() string {
	return "sayhello_record"
}
