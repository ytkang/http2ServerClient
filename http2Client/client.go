package main

import (
	"bytes"
	"crypto/tls"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	c := &Client{
		client: &http.Client{
			Transport: &http2.Transport{
				// So http2.Transport doesn't complain the URL scheme isn't 'https'
				AllowHTTP: false,
				// Pretend we are dialing a TLS endpoint.
				// Note, we ignore the passed tls.Config
				DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
		},
	}
	return c
}

func (c *Client) Post(host, path string, data []byte) ([]byte, error) {
	req := &http.Request{
		Method: "POST",
		URL: &url.URL{
			Scheme: "https",
			Host:   host,
			Path:   path,
		},
		Header: http.Header{},
		Body:   ioutil.NopCloser(bytes.NewReader(data)),
	}

	// Sends the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 500 {
		return nil, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
