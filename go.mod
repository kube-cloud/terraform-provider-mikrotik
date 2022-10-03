module github.com/kube-cloud/terraform-provider-mikrotik

go 1.16

require (
	github.com/hashicorp/terraform-plugin-docs v0.13.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.20.0
	github.com/kube-cloud/terraform-provider-mikrotik/client v0.0.0-00010101000000-000000000000
)

replace github.com/kube-cloud/terraform-provider-mikrotik/client => ./client
