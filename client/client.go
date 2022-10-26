package client

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	t "github.com/jackall3n/terraform-provider-render/client/types"
)

type Client struct {
	BaseURL      string
	Organisation string
	Token        string

	HTTPClient *resty.Client

	Owner *t.Owner
}

const HostUrl string = "https://api.render.com/v1"

func NewClient(host, email, apiKey *string) (*Client, error) {
	c := Client{
		HTTPClient: resty.New().EnableTrace().SetHeader("Accept", "application/json"),
		BaseURL:    HostUrl,
	}

	if host != nil {
		c.BaseURL = *host
	}

	c.SetApiKey(*apiKey)

	owners, err := c.GetOwners()

	if err != nil {
		return nil, err
	}

	for _, o := range owners {
		if o.Owner.Email == *email {
			c.Owner = &o.Owner
		}
	}

	if c.Owner == nil {
		return nil, errors.New("unable to find an owner")
	}

	return &c, nil
}

func (c *Client) SetApiKey(apiKey string) *Client {
	c.HTTPClient.SetAuthScheme("Bearer").SetAuthToken(apiKey)

	return c
}

func (c *Client) Get(path string, res interface{}) error {
	resp, err := c.HTTPClient.R().
		EnableTrace().
		SetResult(&res).
		Get(fmt.Sprintf("%s/%s", c.BaseURL, path))

	if err != nil {
		return err
	}

	if false {
		return errors.New(resp.String())
	}

	return nil
}

func (c *Client) Post(path string, body interface{}, res interface{}) error {
	resp, err := c.HTTPClient.R().
		EnableTrace().
		SetBody(body).
		SetResult(&res).
		Post(fmt.Sprintf("%s/%s", c.BaseURL, path))

	if err != nil {
		return err
	}

	if false {
		return errors.New(resp.String())
	}

	return nil
}

func (c *Client) Put(path string, body interface{}, res interface{}) error {
	_, err := c.HTTPClient.R().
		EnableTrace().
		SetBody(body).
		SetResult(&res).
		Put(fmt.Sprintf("%s/%s", c.BaseURL, path))

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(path string, res interface{}) error {
	_, err := c.HTTPClient.R().
		EnableTrace().
		SetResult(&res).
		Delete(fmt.Sprintf("%s/%s", c.BaseURL, path))

	if err != nil {
		return err
	}

	return nil
}
