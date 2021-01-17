package yookassa

import (
	"net/http"
)

func WithBaseURL(url string) func(*Client) {
	return func(c *Client) {
		c.baseURL = url
	}
}

func WithHTTPClient(client *http.Client) func(*Client) {
	return func(c *Client) {
		c.http = client
	}
}
