package render

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	r "github.com/jackall3n/terraform-provider-render/client"
	t "github.com/jackall3n/terraform-provider-render/client/types"
	"strings"
)

func resourceServiceEnvironment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServiceEnvironmentCreate,
		ReadContext:   resourceServiceEnvironmentRead,
		UpdateContext: resourceServiceEnvironmentUpdate,
		DeleteContext: resourceServiceEnvironmentDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service": {
				Type:     schema.TypeString,
				Required: true,
			},
			"variables": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceServiceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*r.Client)

	serviceId := d.Get("service").(string)

	var variables []t.EnvVar

	if v, ok := d.GetOk("variables"); ok {
		variables = *transformEnvironmentVariables(v)
	}

	if variables == nil {
		return diag.FromErr(errors.New("no variables were supplied"))
	}

	_, err := c.UpdateServiceEnvironmentVariables(serviceId, variables)

	if err != nil {
		return diag.FromErr(err)
	}

	id := "env-" + serviceId

	d.SetId(id)

	return resourceServiceEnvironmentRead(ctx, d, m)
}

func resourceServiceEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*r.Client)

	var diags diag.Diagnostics

	id := d.Id()

	id = strings.Replace(id, "env-", "", 1)

	variables, err := c.GetServiceEnvironmentVariables(id)

	if err != nil {
		return diag.FromErr(err)
	}

	vars := map[string]interface{}{}

	fmt.Println(variables)

	for _, v := range variables {
		vars[v.EnvVar.Key] = v.EnvVar.Value
	}

	if err := d.Set("variables", vars); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceServiceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*r.Client)

	serviceId := d.Get("service").(string)

	var variables []t.EnvVar

	if v, ok := d.GetOk("variables"); ok {
		variables = *transformEnvironmentVariables(v)
	}

	if variables == nil {
		return diag.FromErr(errors.New("no variables were supplied"))
	}

	_, err := c.UpdateServiceEnvironmentVariables(serviceId, variables)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceServiceEnvironmentRead(ctx, d, m)
}

func resourceServiceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*r.Client)

	var diags diag.Diagnostics

	serviceId := d.Get("service").(string)

	_, err := c.UpdateServiceEnvironmentVariables(serviceId, []t.EnvVar{})

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func transformEnvironmentVariables(v interface{}) *[]t.EnvVar {
	if v == nil {
		return nil
	}

	var variables []t.EnvVar

	for key, value := range v.(map[string]interface{}) {
		env := t.EnvVar{
			Key:   key,
			Value: value.(string),
		}

		variables = append(variables, env)
	}

	return &variables
}
