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
	Confirmation         Confirmation          `json:"confirmation"`
	Transfers            []TransferDetails     `json:"transfers,omitempty"`

	// Max 16 keys
	// Max key name length 32
	// Max key value length 512
	Metadata map[string]interface{} `json:"metadata,omitempty"`
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
