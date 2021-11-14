package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jalexanderII/solid-pancake/database"
	ApplicationM "github.com/jalexanderII/solid-pancake/services/application/models"
	LifeCycleM "github.com/jalexanderII/solid-pancake/services/lifecycle/models"
	"gorm.io/gorm"
)

type RentalDetails struct {
	ID uint `json:"id"`
	// Payments data
	TotalPayments         int32   `json:"total_payments,omitempty"`
	OnTimePayments        int32   `json:"on_time_payments,omitempty"`
	PercentPaymentsOnTime float32 `json:"percent_payments_on_time,omitempty"`
	TotalRentPaid         int32   `json:"total_rent_paid,omitempty"`
	// Application data
	TotalApplications int32   `json:"total_applications,omitempty"`
	AverageSalary     float32 `json:"average_salary,omitempty"`
}

func CreateRentalDetailsResponse(rentalDetails LifeCycleM.RentalDetails) RentalDetails {
	return RentalDetails{
		ID:                    rentalDetails.ID,
		TotalPayments:         rentalDetails.TotalPayments,
		OnTimePayments:        rentalDetails.OnTimePayments,
		PercentPaymentsOnTime: rentalDetails.PercentPaymentsOnTime,
		TotalRentPaid:         rentalDetails.TotalRentPaid,
		TotalApplications:     rentalDetails.TotalApplications,
		AverageSalary:         rentalDetails.AverageSalary,
	}
}

func GetRentalDetails(c *fiber.Ctx) error {
	var rentalDetails []LifeCycleM.RentalDetails
	database.Database.Db.Find(&rentalDetails)
	responseRentalDetails := make([]RentalDetails, len(rentalDetails))

	for idx, rentalDetail := range rentalDetails {
		responseRentalDetails[idx] = CreateRentalDetailsResponse(rentalDetail)
	}
	return c.Status(fiber.StatusOK).JSON(responseRentalDetails)
}

type UserRentalDetails struct {
	ID                    uint    `json:"id"`
	UserRef               int     `json:"user_id"`
	TotalPayments         int32   `json:"total_payments"`
	OnTimePayments        int32   `json:"on_time_payments"`
	PercentPaymentsOnTime float32 `json:"percent_payments_on_time"`
	TotalRentPaid         int32   `json:"total_rent_paid"`
}

func CreateUserRentalDetailsResponse(rentalDetails LifeCycleM.UserRentalDetails) UserRentalDetails {
	return UserRentalDetails{
		ID:                    rentalDetails.ID,
		UserRef:               rentalDetails.UserRef,
		TotalPayments:         rentalDetails.TotalPayments,
		OnTimePayments:        rentalDetails.OnTimePayments,
		PercentPaymentsOnTime: rentalDetails.PercentPaymentsOnTime,
		TotalRentPaid:         rentalDetails.TotalRentPaid,
	}
}

func GetUserRentalDetails(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure id is and uint")
	}

	var userRentalDetails []LifeCycleM.UserRentalDetails
	database.Database.Db.Where("user_ref = ?", id).Find(&userRentalDetails)
	responseUserRentalDetails := make([]UserRentalDetails, len(userRentalDetails))

	for idx, UserRentalDetail := range userRentalDetails {
		responseUserRentalDetails[idx] = CreateUserRentalDetailsResponse(UserRentalDetail)
	}
	return c.Status(fiber.StatusOK).JSON(responseUserRentalDetails)
}

type RentalDetailsData struct {
	// Types that are assignable to Datatype:
	//	*ApplicationData
	//	*PaymentData
	Data isRentalDetailsDatatype `json:"datatype"`
}

func (m *RentalDetailsData) getDatatype() isRentalDetailsDatatype {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RentalDetailsData) getApplicationData() *ApplicationM.ApplicantFormResponse {
	if x, ok := m.getDatatype().(*ApplicationData); ok {
		return x.ApplicationData
	}
	return nil
}

func (m *RentalDetailsData) getPaymentData() *LifeCycleM.PaymentConfirmation {
	if x, ok := m.getDatatype().(*PaymentData); ok {
		return x.PaymentData
	}
	return nil
}

type isRentalDetailsDatatype interface {
	isRentalDetailsDatatype()
}

type ApplicationData struct {
	ApplicationData *ApplicationM.ApplicantFormResponse `json:"application_data"`
}

type PaymentData struct {
	PaymentData *LifeCycleM.PaymentConfirmation `json:"payment_data"`
}

func (*ApplicationData) isRentalDetailsDatatype() {}

func (*PaymentData) isRentalDetailsDatatype() {}

// SendToDataPipeline accepts a RentalDetailsData and processes the data depending on what type the data is
// Types that are assignable to Datatype:
//	*ApplicationData
//	*PaymentData
func SendToDataPipeline(d *RentalDetailsData) error {
	switch x := d.Data.(type) {
	case *ApplicationData:
		// Upload data related to the application
		err := createApplicationData(x)
		if err != nil {
			return fmt.Errorf("unexpected error processing application rental details %v", err)
		}
	case *PaymentData:
		// Upload data related to payments
		err := createPaymentData(x)
		if err != nil {
			return fmt.Errorf("unexpected error processing payment rental details %v", err)
		}
	default:
		return fmt.Errorf("data has unexpected type %T", x)
	}
	return nil
}

func createPaymentData(p *PaymentData) error {
	var (
		userRentalDetails LifeCycleM.UserRentalDetails
		rentalDetails     LifeCycleM.RentalDetails
	)
	payment := p.PaymentData.Payment
	onTime := 0
	if p.PaymentData.OnTime == true {
		onTime = 1
	}
	onTimePayments := int32(onTime)
	percentPaymentsOnTime := float32(onTime / 1)
	totalRentPaid := int32(payment.Payment.Amount)

	err1 := database.Database.Db.Where("total_payments >= ?", 1).Last(&rentalDetails).Error
	err2 := database.Database.Db.Where("user_ref = ?", payment.UserRef).Last(&userRentalDetails).Error
	if err1 == gorm.ErrRecordNotFound && err2 == gorm.ErrRecordNotFound {
		// will only happen once
		return database.Database.Db.Transaction(func(tx *gorm.DB) error {
			inputDetails := LifeCycleM.RentalDetails{
				TotalPayments:         1,
				OnTimePayments:        onTimePayments,
				PercentPaymentsOnTime: percentPaymentsOnTime,
				TotalRentPaid:         totalRentPaid,
			}

			userDetails := LifeCycleM.UserRentalDetails{
				UserRef:               payment.UserRef,
				TotalPayments:         1,
				OnTimePayments:        onTimePayments,
				PercentPaymentsOnTime: percentPaymentsOnTime,
				TotalRentPaid:         totalRentPaid,
			}
			if err := tx.Create(&inputDetails).Error; err != nil {
				return fmt.Errorf("couldn't create rental details record, %v", err)
			}
			if err := tx.Create(&userDetails).Error; err != nil {
				return fmt.Errorf("couldn't create user rental details record, %v", err)
			}
			// return nil will commit the whole transaction
			return nil
		})
	} else if err2 == gorm.ErrRecordNotFound {
		userDetails := LifeCycleM.UserRentalDetails{
			UserRef:               payment.UserRef,
			TotalPayments:         1,
			OnTimePayments:        onTimePayments,
			PercentPaymentsOnTime: percentPaymentsOnTime,
			TotalRentPaid:         totalRentPaid,
		}
		if err := database.Database.Db.Create(&userDetails).Error; err != nil {
			return fmt.Errorf("couldn't create user rental details record, %v", err)
		}
	}

	prevTotalPayments := rentalDetails.TotalPayments
	prevOnTimePayments := rentalDetails.OnTimePayments
	prevTotalRentPaid := rentalDetails.TotalRentPaid
	updatedRentalDetails := LifeCycleM.RentalDetails{
		TotalPayments:         prevTotalPayments + 1,
		OnTimePayments:        prevOnTimePayments + int32(onTime),
		PercentPaymentsOnTime: float32(prevOnTimePayments + int32(onTime)/(prevTotalPayments+1)),
		TotalRentPaid:         prevTotalRentPaid + int32(payment.Payment.Amount),
	}
	if err := database.Database.Db.Create(&updatedRentalDetails).Error; err != nil {
		return fmt.Errorf("couldn't create rental details record, %v", err)
	}

	if err2 == nil {
		prevUserTotalPayments := userRentalDetails.TotalPayments
		prevUserOnTimePayments := userRentalDetails.OnTimePayments
		prevUserTotalRentPaid := userRentalDetails.TotalRentPaid
		updatedUserRentalDetails := LifeCycleM.UserRentalDetails{
			UserRef:               payment.UserRef,
			TotalPayments:         prevUserTotalPayments + 1,
			OnTimePayments:        prevUserOnTimePayments + int32(onTime),
			PercentPaymentsOnTime: float32(prevUserOnTimePayments + int32(onTime)/(prevUserTotalPayments+1)),
			TotalRentPaid:         prevUserTotalRentPaid + int32(payment.Payment.Amount),
		}
		if err := database.Database.Db.Create(&updatedUserRentalDetails).Error; err != nil {
			return fmt.Errorf("couldn't create user rental details record, %v", err)
		}
	}
	return nil
}

func createApplicationData(a *ApplicationData) error {
	var rentalDetails LifeCycleM.RentalDetails
	application := a.ApplicationData.Application

	if err := database.Database.Db.Where("total_applications >= ?", 1).Last(&rentalDetails).Error; err == gorm.ErrRecordNotFound {
		// will only happen once
		inputDetails := LifeCycleM.RentalDetails{
			TotalApplications: 1,
			AverageSalary:     float32(application.Salary),
		}
		if err := database.Database.Db.Create(&inputDetails).Error; err != nil {
			return fmt.Errorf("couldn't create rental details record, %v", err)
		}
		return nil
	}
	prevTotal := rentalDetails.TotalApplications
	prevAvg := rentalDetails.AverageSalary
	updatedRentalDetails := LifeCycleM.RentalDetails{
		TotalApplications: prevTotal + 1,
		AverageSalary:     addToAverage(prevAvg, prevTotal, application.Salary),
	}
	if err := database.Database.Db.Create(&updatedRentalDetails).Error; err != nil {
		return fmt.Errorf("couldn't create rental details record, %v", err)
	}
	return nil
}

func addToAverage(prevAvg float32, prevTotal int32, value int32) float32 {
	return (float32(prevTotal)*prevAvg + float32(value)) / float32(prevTotal+1)
}
