package pkg

import (
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
