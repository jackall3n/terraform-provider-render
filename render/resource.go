package render

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func providerResource() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"render_service":             resourceService(),
		"render_service_environment": resourceServiceEnvironment(),
	}
}
