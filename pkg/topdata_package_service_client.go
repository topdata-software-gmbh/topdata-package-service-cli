package pkg

import (
	"io"
	"net/http"
)

type TopdataPackageServiceClient struct {
	BaseURL    string
	Username   string
	Password   string
	HTTPClient *http.Client
}

func NewTopdataPackageServiceClient(baseURL, username, password string) *TopdataPackageServiceClient {
	return &TopdataPackageServiceClient{
		BaseURL:    baseURL,
		Username:   username,
		Password:   password,
		HTTPClient: &http.Client{},
	}
}

func (c *TopdataPackageServiceClient) ListRepositories() ([]byte, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/repositories", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *TopdataPackageServiceClient) Ping() ([]byte, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/ping", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
