# Render Provider

The Render provider is used to interact with https://render.com

Use the navigation to the left to read about the available resources.

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "render" {
  apiKey = "your-api-key"
  email  = "your-render-email"
}
```