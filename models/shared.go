package models

import (
	"gorm.io/datatypes"
)

type Place struct {
	Address string  `json:"address,omitempty"`
	Street  string  `gorm:"not null" json:"street,omitempty"`
	City    string  `gorm:"not null" json:"city,omitempty"`
	State   string  `gorm:"not null" json:"state,omitempty"`
	Zip     string  `gorm:"not null" json:"zip,omitempty"`
	Unit    string  `json:"unit,omitempty"`
	Lat     float64 `gorm:"not null" json:"lat,omitempty"`
	Lng     float64 `gorm:"not null" json:"lng,omitempty"`
}

type ListingMetrics struct {
	AvailableOn     datatypes.Date   `json:"available_on,omitempty"`
	LastPriceChange *LastPriceChange `gorm:"embedded" json:"last_price_change,omitempty"`
	DaysOnMarket    int              `json:"days_on_market,omitempty"`
}

type LastPriceChange struct {
	Rent       int            `json:"rent"`
	PriceDelta int            `json:"price_delta,"`
	ChangedOn  datatypes.Date `json:"changed_on,"`
}
