package client

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient       *http.Client
	connectionConfig ConnectionConfig
}

type ConnectionConfig struct {
	Timeout time.Duration
	APIKey  string
}

func NewClient(httpClient *http.Client, connectionConfig ConnectionConfig) *Client {
	return &Client{httpClient: httpClient, connectionConfig: connectionConfig}
}

func (c *Client) GetSummary(ctx context.Context, diff string) (string, error) {
	url := &url.URL{Scheme: "https", Host: "api.openai.com", Path: "/v1/chat/completions"}
	r, err := http.NewRequestWithContext(ctx, "POST	", url.String(), nil)
	if err != nil {
		return "", err
	}

	r.Header.Set("Authorization", "Bearer "+c.connectionConfig.APIKey)
	r.Header.Set("Content-Type", "application/json")
	return "", nil
}
