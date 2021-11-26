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
	TotalPayments         int32   `json:"total_payments"`
	OnTimePayments        int32   `json:"on_time_payments"`
	PercentPaymentsOnTime float32 `json:"percent_payments_on_time"`
	TotalRentPaid         int32   `json:"total_rent_paid"`
	TotalApplications     int32   `json:"total_applications"`
	AverageSalary         float32 `json:"average_salary"`
}

type UserRentalDetails struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	UserRef               int        `json:"user_id"`
	User                  UserM.User `gorm:"foreignKey:UserRef"`
	TotalPayments         int32      `json:"total_payments"`
	OnTimePayments        int32      `json:"on_time_payments"`
	PercentPaymentsOnTime float32    `json:"percent_payments_on_time"`
	TotalRentPaid         int32      `json:"total_rent_paid"`
}

type PaymentRequest struct {
	// gorm.Model Embedded Struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Purpose string     `json:"purpose"`
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

type PaymentInfo struct {
	Purpose    string  `json:"purpose"`
	Amount     float32 `json:"amount"`
	CardNumber string  `json:"card_number"`
	NameOnCard string  `json:"name"`
	CVC        int32   `json:"cvc"`
	ZipCode    int32   `json:"zip_code"`
}

type PaymentConfirmation struct {
	ReferenceId uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()" json:"reference_id"`
	Timestamp   int64           `gorm:"autoCreateTime" json:"time_stamp"`
	Status      string          `json:"status"`
	PaymentRef  int             `json:"payment_id"`
	Payment     PaymentResponse `gorm:"foreignKey:PaymentRef"`
	OnTime      bool            `json:"on_time"`
}
