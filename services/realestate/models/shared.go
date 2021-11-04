package models

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
	AvailableOn  string `json:"available_on"`
	DaysOnMarket int    `json:"days_on_market"`
}
