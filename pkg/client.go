package pkg

import (
	"io/ioutil"
	"net/http"
)

type TopdataPackageServiceClient struct {
	BaseURL    string
	Username   string
	Password   string
	HTTPClient *http.Client
}

func NewClient(baseURL, username, password string) *TopdataPackageServiceClient {
	return &TopdataPackageServiceClient{
		BaseURL:    baseURL,
		Username:   username,
		Password:   password,
		HTTPClient: &http.Client{},
	}
}

func (c *TopdataPackageServiceClient) FetchRepositories() ([]byte, error) {
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
