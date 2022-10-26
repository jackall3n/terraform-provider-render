package render

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func providerDataSource() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"render_services": dataSourceServices(),
	}
}
