package pkg

import (
	"net/http"
)

type topdataPackageServiceClient struct {
	BaseURL      string
	Username     string
	Password     string
	HTTPClient   *http.Client
}

func NewClient(baseURL, username, password string) *topdataPackageServiceClient {
	return &topdataPackageServiceClient{
		BaseURL:    baseURL,
		Username:   username,
		Password:   password,
		HTTPClient: &http.Client{},
	}
}
