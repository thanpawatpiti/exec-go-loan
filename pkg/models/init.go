package models

import "gorm.io/gorm"

type Init struct {
	DB *gorm.DB
}

func NewInit(db *gorm.DB) *Init {
	return &Init{DB: db}
}
