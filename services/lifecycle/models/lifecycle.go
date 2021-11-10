package models

import (
	"github.com/google/uuid"
	RealEstateM "github.com/jalexanderII/solid-pancake/services/realestate/models"
	UserM "github.com/jalexanderII/solid-pancake/services/users/models"
	"gorm.io/gorm"
)

type RentalDetails struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	UserRef               int        `json:"user_id"`
	User                  UserM.User `gorm:"foreignKey:UserRef"`
	OnTimePayments        int        `json:"on_time_payments"`
	PercentPaymentsOnTime float32    `json:"percent_payments_on_time"`
	TotalRentPaid         int        `json:"total_rent_paid"`
}

type PaymentRequest struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Purpose string     `gorm:"type:enum('rent', 'other');default:'rent'" json:"purpose"`
	Amount  float32    `json:"amount"`
	Period  string     `json:"period"`
	UserRef int        `json:"user_id"`
	User    UserM.User `gorm:"foreignKey:UserRef"`
}

type PaymentResponse struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Payment        PaymentInfo           `gorm:"embedded" json:"payment"`
	UserRef        int                   `json:"user_id"`
	User           UserM.User            `gorm:"foreignKey:UserRef"`
	ApartmentRef   int                   `json:"apartment_id"`
	Apartment      RealEstateM.Apartment `gorm:"foreignKey:ApartmentRef"`
	PaymentRef     int                   `json:"payment_request_id"`
	PaymentRequest PaymentRequest        `gorm:"foreignKey:PaymentRef"`
}

type PaymentConfirmation struct {
	ReferenceId uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()" json:"reference_id"`
	Timestamp   int64           `gorm:"autoCreateTime" json:"time_stamp"` // Use unix seconds as creating time
	Status      string          `json:"status"`
	PaymentRef  int             `json:"payment_id"`
	Payment     PaymentResponse `gorm:"foreignKey:PaymentRef"`
	OnTime      bool            `json:"on_time"`
}

type PaymentInfo struct {
	Purpose    string  `gorm:"type:enum('rent', 'other');default:'rent'" json:"purpose"`
	Amount     float32 `json:"amount"`
	CardNumber string  `json:"card_number"`
	NameOnCard string  `json:"name"`
	CVC        int     `json:"cvc"`
	ZipCode    int     `json:"zip_code"`
}
