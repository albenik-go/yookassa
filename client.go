package yookassa

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

type Client struct {
	baseURL string
	shopID  string
	apiKEY  string
	http    *http.Client
}

func New(id, key string, opts ...func(c *Client)) *Client {
	c := &Client{
		baseURL: "https://api.yookassa.ru/v3/payments",
		shopID:  id,
		apiKEY:  key,
		http:    &http.Client{},
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

// CreatePayment sends the new payment request with specific indepotence key `k`.
func (c *Client) CreatePayment(k string, p *PaymentRequest) (*PaymentResponse, error) {
	return c.CreatePaymentContext(context.Background(), k, p)
}

// CreatePaymentContext sends the new payment request with specific indepotence key `k`.
func (c *Client) CreatePaymentContext(ctx context.Context, k string, p *PaymentRequest) (*PaymentResponse, error) {
	data, err := jsoniter.Marshal(p)
	if err != nil {
		return nil, NewPaymentError(fmt.Errorf("request encode error: %w", err))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewReader(data))
	if err != nil {
		return nil, NewPaymentError(fmt.Errorf("cannot prepare http request: %w", err))
	}

	req.Header.Set("Idempotence-Key", k)

	return c.fetchPayment(req)
}

func (c *Client) FetchPaymentInfo(id string) (*PaymentResponse, error) {
	return c.FetchPaymentInfoContext(context.Background(), id)
}

func (c *Client) FetchPaymentInfoContext(ctx context.Context, id string) (*PaymentResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+id, nil)
	if err != nil {
		return nil, NewPaymentError(fmt.Errorf("cannot prepare http request: %w", err))
	}

	return c.fetchPayment(req)
}

func (c *Client) fetchPayment(req *http.Request) (*PaymentResponse, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, NewPaymentError(fmt.Errorf("request error: %w", err))
	}
	defer resp.Body.Close()

	var payload *PaymentResponse
	if err = jsoniter.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, NewPaymentError(fmt.Errorf("response decode error: %w", err))
	}

	return payload, nil
}

// NewIdempotenceKey generates a new V4 UUID to use as indepotency key
// as recommended in https://yookassa.ru/developers/using-api/basics.
func NewIdempotenceKey() (string, error) {
	key, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("indepotence key (v4 uuid) generation error: %w", err)
	}
	return key.String(), nil
}
