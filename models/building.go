package models

import "gorm.io/gorm"

type Building struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name           string          `gorm:"index" json:"name"`
	Address        *Place          `gorm:"embedded" json:"address"`
	Apartments     []Apartment     `gorm:"polymorphic:Building;" json:"apartments"`
	Size           int             `json:"size,omitempty"`
	ListingMetrics *ListingMetrics `gorm:"embedded" json:"listing_metrics,omitempty"`
	Description    string          `json:"description,omitempty"`
	Amenities      []string        `gorm:"type:text[]" json:"amenities"`
	RealtorRef     int             `json:"realtor_id"`
	Realtor        *Realtor        `gorm:"foreignKey:RealtorRef"`
}
