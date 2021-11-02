package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Building struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name        string         `gorm:"index" json:"name"`
	Address     Place          `gorm:"embedded" json:"address"`
	Description string         `json:"description,omitempty"`
	Amenities   pq.StringArray `gorm:"type:text[]" json:"amenities"`
	RealtorRef  int            `json:"realtor_id"`
	Realtor     Realtor        `gorm:"foreignKey:RealtorRef"`
}
