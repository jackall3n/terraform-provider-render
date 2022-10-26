package render

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	r "github.com/jackall3n/terraform-provider-render/client"
	t "github.com/jackall3n/terraform-provider-render/client/types"
	"reflect"
)

func resourceService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServiceCreate,
		ReadContext:   resourceServiceRead,
		UpdateContext: resourceServiceUpdate,
		DeleteContext: resourceServiceDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repo": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_details": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Describes the Service being deployed.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"env": {
							Type:     schema.TypeString,
							Required: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"env_specific_details": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"build_command": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"start_command": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceServiceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*r.Client)

	service := t.Service{
		Name:    d.Get("name").(string),
		Type:    d.Get("type").(string),
		Repo:    d.Get("repo").(string),
		OwnerId: c.Owner.Id,
	}

	if serviceDetails := transformServiceDetails(d); serviceDetails != nil {
		service.ServiceDetails = *serviceDetails
	}

	deploy, err := c.CreateService(service)

	if err != nil {
		return diag.FromErr(err)
	}

	s := deploy.Service

	d.SetId(s.ID)

	return resourceServiceRead(ctx, d, m)
}

func transformServiceDetails(d *schema.ResourceData) *t.ServiceDetails {
	raw, ok := d.GetOk("service_details")

	if !ok {
		return nil
	}

	value := raw.([]interface{})[0].(map[string]interface{})

	var serviceDetails t.ServiceDetails

	if property := reflect.ValueOf(value["env"]); property.IsValid() {
		serviceDetails.Env = value["env"].(string)
	}

	if property := reflect.ValueOf(value["region"]); property.IsValid() {
		serviceDetails.Region = value["region"].(string)
	}

	if envSpecificDetails := transformEnvSpecificDetails(value["env_specific_details"]); envSpecificDetails != nil {
		serviceDetails.EnvSpecificDetails = *envSpecificDetails
	}

	return &serviceDetails
}

func transformEnvSpecificDetails(v interface{}) *t.EnvSpecificDetails {
	raw := v.([]interface{})

	if len(raw) == 0 || raw[0] == nil {
		return nil
	}

	value := raw[0].(map[string]interface{})

	var envSpecificDetails t.EnvSpecificDetails

	if property := reflect.ValueOf(value["build_command"]); property.IsValid() {
		envSpecificDetails.BuildCommand = value["build_command"].(string)
	}

	if property := reflect.ValueOf(value["start_command"]); property.IsValid() {
		envSpecificDetails.StartCommand = value["start_command"].(string)
	}

	return &envSpecificDetails
}

func resourceServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*r.Client)

	var diags diag.Diagnostics

	id := d.Id()

	s, err := c.GetService(id)

	if err != nil {
		return diag.FromErr(err)
	}

	properties := map[string]interface{}{
		"id":   s.ID,
		"name": s.Name,
		"type": s.Type,
		"repo": s.Repo,
	}

	for key, value := range properties {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*r.Client)

	id := d.Id()

	service := t.Service{
		Name: d.Get("name").(string),
		Type: d.Get("type").(string),
	}

	_, err := c.UpdateService(id, service)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceServiceRead(ctx, d, m)
}

func resourceServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*r.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	_, err := c.DeleteService(id)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
