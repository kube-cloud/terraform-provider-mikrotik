package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define IPSec Policy Group Resource
 */
func resourceIpSecPolicyGroup() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a IPSec Policy Group resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createIpSecPolicyGroup,

		// Read Resource Context Method CallBack
		ReadContext: readIpSecPolicyGroup,

		// Update Resource Context Method CallBack
		UpdateContext: updateIpSecPolicyGroup,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteIpSecPolicyGroup,

		// Define Resource State Context Importer
		Importer: &schema.ResourceImporter{

			// Define State Context
			StateContext: schema.ImportStatePassthroughContext,
		},

		// Define Resource Schema
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Description: "IPSec Policy Group Name.",
			},
		},
	}
}

/**
 * Create IPSec Policy Group from Resource Data
 */
func createIpSecPolicyGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get IPSec Policy Group Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Policy Group
	dataStructure := dataToIpSecPolicyGroup(d)

	// Add IPSec Policy Group
	ipsecPolicyGroup, err := c.AddIpSecPolicyGroup(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Policy Group to Resource Data and put it in Resource Pointer
	ipsecPolicyGroupToData(ipsecPolicyGroup, d)

	// Reload IPSec Policy Group
	return readIpSecPolicyGroup(ctx, d, m)
}

/**
 * Read IPSec Policy Group from Resource Data
 */
func readIpSecPolicyGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Policy Group Client
	c := m.(*client.Mikrotik)

	// Find IPSec Policy Group
	ipsecPolicyGroup, err := c.FindIpSecPolicyGroup(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Policy Group to Resource Data and put it in Resource Pointer
	ipsecPolicyGroupToData(ipsecPolicyGroup, d)

	// Return Diagnistic
	return diags
}

/**
 * Update IPSec Policy Group from Resource Data
 */
func updateIpSecPolicyGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Policy Group Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Policy Group
	dataStructure := dataToIpSecPolicyGroup(d)

	// Update IPSec Policy Group
	_, err := c.UpdateIpSecPolicyGroup(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete IPSec Policy Group from Resource Data
 */
func deleteIpSecPolicyGroup(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Policy Group Client
	c := m.(*client.Mikrotik)

	// Delete IPSec Policy Group
	err := c.DeleteIpSecPolicyGroup(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to IPSec Policy Group
 */
func dataToIpSecPolicyGroup(d *schema.ResourceData) *client.IpSecPolicyGroup {

	// Build and Return IPSec Policy Group
	return &client.IpSecPolicyGroup{
		Id:   d.Id(),
		Name: d.Get("name").(string),
	}
}

/**
 * Function used to Convert IPSec Policy Group to Resource Data
 */
func ipsecPolicyGroupToData(ipsecPolicyGroup *client.IpSecPolicyGroup, d *schema.ResourceData) {

	// Initialize Resource ID
	d.SetId(ipsecPolicyGroup.Id)

	// Initialize Fields
	d.Set("name", ipsecPolicyGroup.Name)
}
