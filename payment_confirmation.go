package yookassa

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

const (
	ConfirmationRedirect = "redirect"
	ConfirmationQR       = "qr"
	ConfirmationEmbeded  = "embedded"
	ConfirmationExternal = "external"
)

// Request

type RedirectConfirmationRequestData struct {
	ReturnURL string `json:"return_url"`
	Locale    string `json:"locale,omitempty"`
	Enforce   bool   `json:"enforce,omitempty"`
}

type redirectConfirmationRequestData struct {
	Type      string `json:"type"`
	ReturnURL string `json:"return_url"`
	Locale    string `json:"locale,omitempty"`
	Enforce   bool   `json:"enforce,omitempty"`
}

func (c *RedirectConfirmationRequestData) EncodeJSON() ([]byte, error) {
	return jsoniter.Marshal(&redirectConfirmationRequestData{
		Type:      ConfirmationRedirect,
		ReturnURL: c.ReturnURL,
		Locale:    c.Locale,
		Enforce:   c.Enforce,
	})
}

type SimpleConfirmationRequestData struct {
	Type   string `json:"type"`
	Locale string `json:"locale,omitempty"`
}

// Response

type Confirmation struct {
	Type    string      `json:"type"`
	Details interface{} `json:"details,omitempty"`
}

func (c *Confirmation) UnmarshalJSON(data []byte) error {
	c.Type = jsoniter.Get(data, "type").ToString()

	switch c.Type {
	case ConfirmationRedirect:
		c.Details = &RedirectConfirmationDetails{}
	case ConfirmationQR:
		c.Details = &QRCodeConfirmationDetails{}
	case ConfirmationEmbeded:
		c.Details = &EmbededConfirmationDetails{}
	case ConfirmationExternal:
		c.Details = nil
	default:
		return fmt.Errorf("invalid confirmation type %q", c.Type) // nolint:goerr113
	}

	if c.Details != nil {
		return jsoniter.Unmarshal(data, c.Details)
	}
	return nil
}

type RedirectConfirmationDetails struct {
	Type            string `json:"type"`
	Enforce         bool   `json:"enforce"`
	ConfirmationURL string `json:"confirmation_url"`
	ReturnURL       string `json:"return_url"`
}

type EmbededConfirmationDetails struct {
	Type              string `json:"type"`
	ConfirmationToken string `json:"confirmation_token"`
}

type QRCodeConfirmationDetails struct {
	Type             string `json:"type"`
	ConfirmationData string `json:"confirmation_data"`
}
