package handlers

import (
	"time"

	"github.com/google/uuid"
	"github.com/jalexanderII/solid-pancake/database"
	LifeCycleM "github.com/jalexanderII/solid-pancake/services/lifecycle/models"
)

// PaymentConfirmation To be used as a serializer
type PaymentConfirmation struct {
	ReferenceId uuid.UUID       `json:"reference_id"`
	Timestamp   int64           `json:"time_stamp"`
	Status      string          `json:"status,omitempty"`
	Payment     PaymentResponse `json:"payment" validate:"dive"`
	OnTime      bool            `json:"on_time"`
}

// CreatePaymentConfirmation Takes in a model and returns a serializer
func CreatePaymentConfirmation(paymentConfirmation LifeCycleM.PaymentConfirmation) PaymentConfirmation {
	var paymentResponse LifeCycleM.PaymentResponse
	database.Database.Db.First(&paymentResponse, paymentConfirmation.PaymentRef)
	return PaymentConfirmation{
		ReferenceId: paymentConfirmation.ReferenceId,
		Timestamp:   paymentConfirmation.Timestamp,
		Status:      paymentConfirmation.Status,
		Payment:     CreatePaymentResponse(paymentResponse),
		OnTime:      paymentConfirmation.OnTime,
	}
}

func PaymentConfirmationResponse(id int, status string, onTime bool) (PaymentConfirmation, error) {
	var paymentResponse LifeCycleM.PaymentResponse
	if err := findPaymentResponse(id, &paymentResponse); err != nil {
		return PaymentConfirmation{}, err
	}

	paymentConfirmation := LifeCycleM.PaymentConfirmation{
		ReferenceId: uuid.New(),
		Timestamp:   time.Now().Unix(),
		Status:      status,
		Payment:     paymentResponse,
		OnTime:      onTime,
	}
	if err := database.Database.Db.Create(&paymentConfirmation).Error; err != nil {
		return PaymentConfirmation{}, err
	}
	paymentConfirmationResponse := CreatePaymentConfirmation(paymentConfirmation)
	return paymentConfirmationResponse, nil
}

func SendToPaymentService(paymentResponseID int) (PaymentConfirmation, error) {
	// TODO: Pass payment to external payment provider
	status := "COMPLETE"
	return PaymentConfirmationResponse(paymentResponseID, status, true)
}
