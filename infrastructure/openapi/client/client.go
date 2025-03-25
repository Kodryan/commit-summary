package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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
	ctx, cancel := context.WithTimeout(ctx, c.connectionConfig.Timeout)
	defer cancel()
	r, err := http.NewRequestWithContext(ctx, "POST", url.String(), nil)
	if err != nil {
		return "", err
	}

	r.Header.Set("Authorization", "Bearer "+c.connectionConfig.APIKey)
	r.Header.Set("Content-Type", "application/json")
	request := OpenAIRequest{
		Model:    "gpt-4o-mini-2024-07-18",
		Store:    true,
		Messages: []Message{{Role: "user", Content: "write very short commit message for the following diff: " + diff}},
	}

	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(request); err != nil {
		return "", err
	}
	r.Body = io.NopCloser(body)
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
