package render

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	r "github.com/jackall3n/terraform-provider-render/client"
	"strconv"
	"time"
)

func dataSourceServices() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServicesRead,
		Schema: map[string]*schema.Schema{
			"applications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*r.Client)

	var diags diag.Diagnostics

	response, err := c.GetServices()

	if err != nil {
		return diag.FromErr(err)
	}

	services := make([]map[string]interface{}, 0)

	for _, s := range response {
		service := make(map[string]interface{})

		service["id"] = s.Service.ID
		service["name"] = s.Service.Name
		service["type"] = s.Service.Type

		services = append(services, service)
	}

	if err := d.Set("services", services); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
