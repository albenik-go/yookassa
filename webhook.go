package yookassa

const (
	EventWaitingForCapture = "payment.waiting_for_capture"
	EventPaymentSucceeded  = "payment.succeeded"
	EventPaymentCancelled  = "payment.canceled"
	EventRefundSucceeded   = "refund.succeeded"
)

type Event struct {
	Type   string           `json:"type"`
	Name   string           `json:"event"`
	Object *PaymentResponse `json:"object"`
}
