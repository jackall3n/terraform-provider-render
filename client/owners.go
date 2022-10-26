package client

import (
	t "github.com/jackall3n/terraform-provider-render/client/types"
)

func (c *Client) GetOwners() ([]t.OwnerItem, error) {
	var owners []t.OwnerItem

	err := c.Get("owners", &owners)

	if err != nil {
		return nil, err
	}

	return owners, nil
}
