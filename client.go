package autodns

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"os"

	"models"
)

const DEFAULT_URL = "https://gateway.autodns.com"

var REQUIRED_ENV = []string{"AUTODNS_USER", "AUTODNS_PASSWORD", "AUTODNS_CONTEXT"}

type Callback func(resp *http.Response, err error)

type Client struct {
	Url      string
	Context  string
	SystemNs []models.SystemNameServer
	User     string
	password string
}

func NewClient(user, password, context string, url string, sysNS []models.SystemNameServer) (*Client, error) {
	if user == "" || password == "" || context == "" {
		return nil, errors.New("invalid client settings")
	}
	if url == "" {
		url = DEFAULT_URL
	}
	return &Client{
		Url:      url,
		Context:  context,
		SystemNs: sysNS,
		User:     user,
		password: password,
	}, nil
}

func NewClientFromEnv() (*Client, error) {
	for _, k := range REQUIRED_ENV {
		if os.Getenv(k) == "" {
			return nil, fmt.Errorf("%s needs to be set", k)
		}
	}
	url := os.Getenv("AUTODNS_URL")
	user := os.Getenv("AUTODNS_USER")
	password := os.Getenv("AUTODNS_PASSWORD")
	context := os.Getenv("AUTODNS_CONTEXT")
	return NewClient(user, password, context, url, nameServersFromEnv())
}

func nameServersFromEnv() []models.SystemNameServer {
	keys := []string{"AUTODNS_NS1", "AUTODNS_NS2", "AUTODNS_NS3", "AUTODNS_NS4"}
	out := make([]models.SystemNameServer, 0)
	for _, key := range keys {
		ns, ok := os.LookupEnv(key)
		if ok {
			out = append(out, models.SystemNameServer(ns))
		}
	}
	return out
}

func (c *Client) GetAuth() *models.Auth {
	return &models.Auth{
		User:     c.User,
		Password: c.password,
		Context:  c.Context,
	}
}

func (c *Client) PerformRequest(request *models.Request, callback Callback) {
	payload, err := xml.Marshal(request)
	if err != nil {
		callback(nil, err)
	}
	responseBody := bytes.NewBuffer(payload)
	resp, err := http.Post(c.Url, "application/json", responseBody)
	callback(resp, err)
}
