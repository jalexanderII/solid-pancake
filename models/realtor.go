package models

import "gorm.io/gorm"

type Realtor struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name        string `gorm:"index" json:"name"`
	Company     string `json:"company"`
	PhoneNumber string `json:"phone_number"`
}
