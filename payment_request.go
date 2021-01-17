package yookassa

import (
	"time"
)

type PaymentRequest struct {
	Amount            Amount                 `json:"amount"`
	Capture           bool                   `json:"capture,omitempty"`
	PaymentToken      string                 `json:"payment_token,omitempty"`
	PaymentMethodID   string                 `json:"payment_method_id,omitempty"`
	PaymentMethodData *PaymentMethod         `json:"payment_method_data,omitempty"`
	SavePaymentMethod bool                   `json:"save_payment_method,omitempty"`
	Receipt           *ReceiptRequestData    `json:"receipt,omitempty"`
	ClientIP          string                 `json:"client_ip,omitempty"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	AirlineTicket     *AirlineTicketData     `json:"airline,omitempty"`
	Transfers         []TransferRequestData  `json:"transfers,omitempty"`
}

type ReceiptRequestData struct {
	Customer CustomerInfo `json:"customer"`
}

type CustomerInfo struct {
	FullName string `json:"full_name,omitempty"`
	INN      string `json:"inn,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type TransferRequestData struct {
	AccountID         string `json:"account_id"`
	Amount            Amount `json:"amount"`
	PlatformFeeAmount Amount `json:"platform_fee_amount"`
}

type AirlineTicketData struct {
	TicketNumber     string                       `json:"ticket_number,omitempty"`
	BookingReference string                       `json:"booking_reference,omitempty"`
	Passengers       []AirlineTocketPassengerData `json:"passengers,omitempty"`
	Legs             []AirlineTicketLegData       `json:"legs"`
}

type AirlineTocketPassengerData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AirlineTicketLegData struct {
	DepartureAirport   string    `json:"departure_airport"`
	DepartureDate      time.Time `json:"departure_date"`
	DestinationAirport string    `json:"destination_airport"`
	CarrierCode        string    `json:"carrier_code"`
}
