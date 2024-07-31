package models

import "gorm.io/gorm"

type Bid struct {
	gorm.Model
	Price  float64 `json:"price"`
	UserID uint    `json:"user_id"`
	User   User    `json:"user" gorm:"foreignKey:UserID"`
}
