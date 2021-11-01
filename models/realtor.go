package models

import "gorm.io/gorm"

type Realtor struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name        string `gorm:"index" json:"name"`
	Address     *Place `gorm:"embedded" json:"address"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
