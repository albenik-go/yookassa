package yookassa

import (
	"time"
)

const (
	StatusPending  = "pending"
	StatusWaiting  = "waiting_for_capture"
	SatusSucceeded = "succeeded"
	StatusCanceled = "canceled"
)

type PaymentRequest struct {
	Amount            Amount                 `json:"amount"`
	Description       string                 `json:"description,omitempty"`
	Receipt           *ReceiptRequestData    `json:"receipt,omitempty"`
	Recipient         *RecipientRequest      `json:"recipient,omitempty"`
	PaymentToken      string                 `json:"payment_token,omitempty"`
	PaymentMethodID   string                 `json:"payment_method_id,omitempty"`
	PaymentMethodData *PaymentMethod         `json:"payment_method_data,omitempty"`
	Confirmation      interface{}            `json:"confirmation"`
	SavePaymentMethod bool                   `json:"save_payment_method,omitempty"`
	Capture           bool                   `json:"capture,omitempty"`
	ClientIP          string                 `json:"client_ip,omitempty"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	AirlineTicket     *AirlineTicketData     `json:"airline,omitempty"`
	Transfers         []TransferRequestData  `json:"transfers,omitempty"`
}

type PaymentResponse struct {
	ID                   string                `json:"id"`
	Status               string                `json:"status"`
	Test                 bool                  `json:"test"`
	Paid                 bool                  `json:"paid"`
	Refundable           bool                  `json:"refundable"`
	Amount               Amount                `json:"amount"`
	IncomeAmount         *Amount               `json:"income_amount,omitempty"`
	RefundedAmount       *Amount               `json:"refunded_amount,omitempty"`
	Created              time.Time             `json:"created_at"`
	Captured             time.Time             `json:"captured_at"`
	Expires              time.Time             `json:"expires_at"`
	Description          string                `json:"description"`
	Recipient            Recipient             `json:"recipient"`
	PaymentMethod        PaymentMethod         `json:"payment_method"`
	ReceiptRegistration  string                `json:"receipt_registration,omitempty"`
	AuthorizationDetails *AuthorizationDetails `json:"authorization_details,omitempty"`
	CancellationDetails  *CancellationDetails  `json:"cancellation_details,omitempty"`
	Confirmation         ConfirmationInfo      `json:"confirmation"`
	Transfers            []TransferDetails     `json:"transfers,omitempty"`

	// Max 16 keys
	// Max key name length 32
	// Max key value length 512
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type RecipientRequest struct {
	GatewayID string `json:"gateway_id"`
}

type Recipient struct {
	AccountID string `json:"account_id"`
	GatewayID string `json:"gateway_id"`
}

type Amount struct {
	Value    string
	Currency string
}

type AuthorizationDetails struct {
	RRN      string `json:"rrn"`
	AuthCode string `json:"auth_code"`
}

type CancellationDetails struct {
	Party  string `json:"party"`
	Reason string `json:"reason"`
}

type TransferDetails struct {
	AccountID         string `json:"account_id"`
	Status            string `json:"status"`
	Amount            Amount `json:"amount"`
	PlatformFeeAmount Amount `json:"platform_fee_amount"`
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
