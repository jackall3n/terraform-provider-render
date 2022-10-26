package render

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jackall3n/terraform-provider-render/client"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("api_key").(string)
	email := d.Get("email").(string)

	var host *string

	hVal, ok := d.GetOk("host")

	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c, err := client.NewClient(host, &email, &apiKey)

	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
