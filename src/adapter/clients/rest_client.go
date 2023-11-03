package clients

import (
	"bytes"
	"net/http"
)

type RestClient struct {
	BaseURL string
	Headers map[string]string
}

func NewRestClient(baseURL string, headers map[string]string) *RestClient {
	return &RestClient{BaseURL: baseURL, Headers: headers}
}

func (c *RestClient) Post(path string, payload []byte) (*http.Response, error) {
	url := c.BaseURL + path

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
