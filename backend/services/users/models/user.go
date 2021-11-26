package models

import "gorm.io/gorm"

type User struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name     string `gorm:"unique_index;not null" json:"name"`
	Username string `gorm:"unique_index;" json:"username"`
	Email    string `gorm:"unique_index;" json:"email"`
	Password string `json:"password"`
}
