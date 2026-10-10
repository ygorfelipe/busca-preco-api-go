package httpclient

import "net/http"

type Client struct {
	client *http.Client
}

func New() *Client {
	return &Client{client: &http.Client{}}
}

func (c *Client) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}
