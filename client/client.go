package client

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/rabbitprincess/goany/goany"
)

type Client struct {
	HTTPClient *http.Client
	Headers    map[string]string
}

func NewClient(cli *http.Client) *Client {
	if cli == nil {
		cli = http.DefaultClient
	}

	return &Client{HTTPClient: cli}
}

func (c *Client) Get(ctx context.Context, url string, req *goany.Request) (*goany.Response, error) {
	if req != nil {
		params := req.MarshalUrlParams()
		url = url + "?" + params.Encode()
	}

	return c.Do(ctx, http.MethodGet, url, nil)
}

func (c *Client) Post(ctx context.Context, url string, req *goany.Request) (*goany.Response, error) {
	body, err := req.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, http.MethodPost, url, body)
}

func (c *Client) Do(ctx context.Context, method, url string, body []byte) (*goany.Response, error) {
	var bodyBuf io.Reader
	if body != nil {
		bodyBuf = bytes.NewReader(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyBuf)
	if err != nil {
		return nil, err
	}
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respAny, err := goany.NewResponseFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	return respAny.SetHTTPStatus(resp.StatusCode), nil
}
