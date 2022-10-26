package render

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigure,
		Schema:               providerSchema(),
		DataSourcesMap:       providerDataSource(),
		ResourcesMap:         providerResource(),
	}
}
