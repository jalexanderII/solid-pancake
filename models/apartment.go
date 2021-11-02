package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Apartment struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name           string         `gorm:"index" json:"name" `
	Address        Place          `gorm:"embedded" json:"address"`
	Rent           int            `gorm:"not null" json:"rent"`
	Size           float32        `json:"size"`
	Features       Features       `gorm:"embedded" json:"features"`
	ListingMetrics ListingMetrics `gorm:"embedded" json:"listing_metrics"`
	Description    string         `json:"description"`
	Amenities      pq.StringArray `gorm:"type:text[]" json:"amenities"`
	// A belongs to association sets up a one-to-one connection with another model,
	// such that each instance of the declaring model “belongs to” one instance of the other model.
	BuildingRef  int          `json:"building_id"`
	Building     Building     `gorm:"foreignKey:BuildingRef"`
	RealtorRef   int          `json:"realtor_id"`
	Realtor      Realtor      `gorm:"foreignKey:RealtorRef"`
}

type Features struct {
	Beds  int     `gorm:"not null" json:"beds"`
	Baths float32 `gorm:"not null" json:"baths"`
	Rooms int     `json:"rooms"`
}
