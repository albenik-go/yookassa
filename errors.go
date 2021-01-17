package yookassa

import (
	"errors"
	"fmt"
)

type PaymentError struct {
	InternalError error
}

func (e *PaymentError) Error() string {
	return fmt.Sprintf("cannot create payment: %s", e.InternalError.Error())
}

func (e *PaymentError) Is(err error) bool {
	return err == e || errors.Is(err, e.InternalError) // nolint:errorlint,goerr113
}

func NewPaymentError(err error) *PaymentError {
	return &PaymentError{InternalError: err}
}
