package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	timeout    time.Duration
	clientHttp *http.Client
}

func New(options ...func(*Client)) *Client {
	c := &Client{}
	for _, opt := range options {
		opt(c)
	}
	c.clientHttp = http.DefaultClient
	return c
}

func WithBaseURL(baseURL string) func(*Client) {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func WithTimeout(timeout time.Duration) func(*Client) {
	return func(c *Client) {
		c.timeout = timeout
	}
}

func (c *Client) Post(ctx context.Context, endpoint string, headers map[string]string, body any) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	b, _ := json.Marshal(body)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/"+endpoint, bytes.NewReader(b))

	c.setHeaders(req, headers)

	return c.clientHttp.Do(req)
}

func (c *Client) Get(ctx context.Context, endpoint string, headers map[string]string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+endpoint, nil)

	c.setHeaders(req, headers)

	return c.clientHttp.Do(req)
}

func (c *Client) Delete(ctx context.Context, endpoint string, headers map[string]string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodDelete, c.baseURL+"/"+endpoint, nil)

	c.setHeaders(req, headers)

	return c.clientHttp.Do(req)
}

func (c *Client) setHeaders(req *http.Request, headers map[string]string) *http.Request {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req
}
