package models

import (
	"gorm.io/gorm"
)

type Apartment struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name           string          `gorm:"index" json:"name" `
	Address        *Place          `gorm:"embedded" json:"address"`
	Rent           int             `gorm:"not null" json:"rent"`
	Size           float32         `json:"size"`
	Features       *Features       `gorm:"embedded" json:"features"`
	ListingMetrics *ListingMetrics `gorm:"embedded" json:"listing_metrics"`
	Description    string          `json:"description"`
	Amenities      []string        `gorm:"type:text[]" json:"amenities"`
	// A belongs to association sets up a one-to-one connection with another model,
	// such that each instance of the declaring model “belongs to” one instance of the other model.
	BuildingRef  int           `json:"building_id"`
	Building     *Building     `gorm:"foreignKey:BuildingRef"`
	PriceHistory *PriceHistory `gorm:"embedded" json:"price_history"`
	RealtorRef   int           `json:"realtor_id"`
	Realtor      *Realtor      `gorm:"foreignKey:RealtorRef"`
	BuildingID   int
	BuildingType string
}

type Features struct {
	Beds  int     `gorm:"not null" json:"beds"`
	Baths float32 `gorm:"not null" json:"baths"`
	Rooms int     `json:"rooms,omitempty"`
}

type PriceHistory struct {
	Date []string  `gorm:"type:text[]" json:"date"`
	Rent []float32 `gorm:"type:float[]" json:"rent"`
}
