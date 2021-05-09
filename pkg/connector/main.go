package connector

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Connector struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

func NewConnector(url *url.URL) *Connector {
	return &Connector{
		BaseURL:    url,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Connector) Get(endpoint string) (*http.Response, error){
	rel := &url.URL{Path: fmt.Sprintf(endpoint)}
	u := c.BaseURL.ResolveReference(rel)

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Connector) Post(endpoint string, bytes *bytes.Buffer) (*http.Response, error){
	rel := &url.URL{Path: fmt.Sprintf(endpoint)}
	u := c.BaseURL.ResolveReference(rel)

	resp, err := http.Post(u.String(), "application/json", bytes)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
