package httper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// ClientCfg holds the configuration for the HTTP client, including a URL prefix and timeout duration.
type ClientCfg struct {
	Prefix  string        `yaml:"prefix"  env:"HTTP_CLIENT_PREFIX"  env-default:""`
	Timeout time.Duration `yaml:"timeout" env:"HTTP_CLIENT_TIMEOUT" env-default:"5s"`
}

// Client represents an HTTP client with a configurable prefix and an underlying http.Client.
type Client struct {
	prefix string
	client *http.Client
}

// NewClient creates and returns a new Client instance configured with the provided ClientCfg.
func NewClient(cfg *ClientCfg) *Client {
	return &Client{
		prefix: cfg.Prefix,
		client: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

// Get performs a GET request to the specified URL, optionally prepending the prefix.
func (c *Client) Get(url string) (*Resp, error) {
	if c.prefix != "" {
		url = c.prefix + url
	}

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return newResp(resp, body), nil
}

// GetJson performs a GET request and decodes the JSON response into the provided interface.
func (c *Client) GetJson(url string, to interface{}) (*Resp, error) {
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp.ByteBody, to)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Post performs a POST request with the given data to the specified URL, optionally prepending the prefix.
func (c *Client) Post(url string, data interface{}) (*Resp, error) {
	if c.prefix != "" {
		url = c.prefix + url
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(body)

	resp, err := c.client.Post(url, string(JsonType), reader)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return newResp(resp, body), nil
}

// PostJson performs a POST request with JSON data and binds the response to the provided interface.
func (c *Client) PostJson(url string, data interface{}, to interface{}) (*Resp, error) {
	resp, err := c.Post(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp.ByteBody, to)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Do executes a custom HTTP request and returns the response.
// It allows for more complex requests by using a Req object that can include additional settings.
func (c *Client) Do(req *Req) (*Resp, error) {
	if c.prefix != "" {
		newUrl, err := url.Parse(c.prefix + fmt.Sprintf("%v", req.URL))
		if err != nil {
			return nil, err
		}

		req.URL = newUrl
	}

	resp, err := c.client.Do(req.Request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if req.NeedUnmarshal {
		err = req.unmarshal(body)
		if err != nil {
			return nil, err
		}
	}

	return newResp(resp, body), nil
}
