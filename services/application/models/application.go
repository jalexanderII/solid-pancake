package models

import (
	"github.com/google/uuid"
	RealEstateM "github.com/jalexanderII/solid-pancake/services/realestate/models"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ApplicantFormRequest struct {
	gorm.Model
	Name            string                `gorm:"index" json:"name"`
	SocialSecurity  string                `json:"social_security,omitempty"`
	DateOfBirth     string                `json:"date_of_birth,omitempty"`
	DriversLicense  string                `json:"drivers_license,omitempty"`
	PreviousAddress RealEstateM.Place     `gorm:"embedded" json:"previous_address,omitempty"`
	Landlord        string                `json:"landlord,omitempty"`
	LandlordNumber  string                `json:"landlord_number,omitempty"`
	Employer        string                `json:"employer,omitempty"`
	Salary          int32                 `json:"salary,omitempty"`
	ApartmentRef    int                   `json:"apartment_id"`
	Apartment       RealEstateM.Apartment `gorm:"foreignKey:ApartmentRef"`
}

type ApplicantFormResponse struct {
	gorm.Model
	ReferenceId    uuid.UUID            `gorm:"type:uuid;default:uuid_generate_v4()"  json:"reference_id"`
	Status         string               `json:"status,omitempty"`
	Attachments    pq.StringArray       `gorm:"type:text[]" json:"attachments"`
	ApplicationRef int                  `json:"application_id"`
	Application    ApplicantFormRequest `gorm:"foreignKey:ApplicationRef"`
}
