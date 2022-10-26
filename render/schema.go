package render

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_key": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("RENDER_API_KEY", nil),
		},
		"email": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("RENDER_EMAIL", nil),
		},
	}
}
