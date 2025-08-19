package apiclientcoincapio

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client  *http.Client
	baseURL string
	apiKey  string
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
		baseURL: "https://api.coincap.io/v2/",
		apiKey:  "3d60337d2f7e66fca106d315654f5b176fbbcb53d2de493ef360d9b7f93f6b4a",
	}, nil
}

func (c *Client) GetAssets() ([]AssetData, error) {
	req, err := http.NewRequest("GET", c.baseURL+"assets", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r AssetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.Data, nil
}
