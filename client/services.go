package client

import (
	"fmt"
	t "github.com/jackall3n/terraform-provider-render/client/types"
)

func (c *Client) GetServices() ([]t.ServiceItem, error) {
	var services []t.ServiceItem

	err := c.Get("services", &services)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	return services, nil
}

func (c *Client) GetServiceEnvironmentVariables(id string) ([]t.EnvVarItem, error) {
	var results []t.EnvVarItem

	err := c.Get(fmt.Sprintf("services/%s/env-vars", id), &results)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c *Client) UpdateServiceEnvironmentVariables(id string, variables []t.EnvVar) (*[]t.EnvVarItem, error) {
	var results []t.EnvVarItem

	err := c.Put(fmt.Sprintf("services/%s/env-vars", id), variables, &results)

	if err != nil {
		return nil, err
	}

	return &results, nil
}

func (c *Client) GetService(id string) (*t.Service, error) {
	var service t.Service

	err := c.Get(fmt.Sprintf("services/%s", id), &service)

	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (c *Client) CreateService(service t.Service) (*t.ServiceDeploy, error) {
	var result t.ServiceDeploy

	err := c.Post("services", service, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateService(id string, service t.Service) (*t.Service, error) {
	var result t.Service

	err := c.Put(fmt.Sprintf("services/%s", id), service, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteService(id string) (interface{}, error) {
	var result interface{}

	err := c.Delete(fmt.Sprintf("services/%s", id), &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
