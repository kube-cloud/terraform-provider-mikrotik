package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define IPSec Policy Resource
 */
func resourceIpSecPolicy() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a IPSec Policy resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createIpSecPolicy,

		// Read Resource Context Method CallBack
		ReadContext: readIpSecPolicy,

		// Update Resource Context Method CallBack
		UpdateContext: updateIpSecPolicy,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteIpSecPolicy,

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
			"peer": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Description: "IPSec Policy Peer.",
			},
			"tunnel": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "IPSec Policy Tunnel Flag.",
			},
			"source_address": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Description: "IPSec Policy Source Address",
			},
			"source_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "IPSec Policy Source Port.",
			},
			"destination_address": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Description: "IPSec Policy Destination Address.",
			},
			"destination_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "IPSec Policy Destination Port.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "all",
				Description: "IPSec Policy Protocol.",
			},
			"template": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "IPSec Policy is Template.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "encrypt",
				Description: "IPSec Policy Action.",
			},
			"level": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "require",
				Description: "IPSec Policy Level.",
			},
			"ipsec_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "esp",
				Description: "IPSec Policy IPSec Protocol.",
			},
			"proposal": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Description: "IPSec Policy Proposal.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "IPSec Policy is Disabled.",
			},
		},
	}
}

/**
 * Create IPSec Policy from Resource Data
 */
func createIpSecPolicy(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get IPSec Policy Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Policy
	dataStructure := dataToIpSecPolicy(d)

	// Add IPSec Policy
	ipsecPolicy, err := c.AddIpSecPolicy(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Policy to Resource Data and put it in Resource Pointer
	ipsecPolicyToData(ipsecPolicy, d)

	// Reload IPSec Policy
	return readIpSecPolicy(ctx, d, m)
}

/**
 * Read IPSec Policy from Resource Data
 */
func readIpSecPolicy(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Policy Client
	c := m.(*client.Mikrotik)

	// Find IPSec Policy
	ipsecPolicy, err := c.FindIpSecPolicy(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Policy to Resource Data and put it in Resource Pointer
	ipsecPolicyToData(ipsecPolicy, d)

	// Return Diagnistic
	return diags
}

/**
 * Update IPSec Policy from Resource Data
 */
func updateIpSecPolicy(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Policy Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Policy
	dataStructure := dataToIpSecPolicy(d)

	// Update IPSec Policy
	_, err := c.UpdateIpSecPolicy(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete IPSec Policy from Resource Data
 */
func deleteIpSecPolicy(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Policy Client
	c := m.(*client.Mikrotik)

	// Delete IPSec Policy
	err := c.DeleteIpSecPolicy(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to IPSec Policy
 */
func dataToIpSecPolicy(d *schema.ResourceData) *client.IpSecPolicy {

	// Build and Return IPSec Policy
	return &client.IpSecPolicy{
		Id:                 d.Id(),
		Peer:               d.Get("peer").(string),
		Tunnel:             d.Get("tunnel").(bool),
		SourceAddress:      d.Get("source_address").(string),
		SourcePort:         d.Get("source_port").(int),
		DestinationAddress: d.Get("destination_address").(string),
		DestinationPort:    d.Get("destination_port").(int),
		Protocol:           d.Get("protocol").(string),
		Template:           d.Get("template").(bool),
		Action:             d.Get("action").(string),
		Level:              d.Get("level").(string),
		IpSecProtocol:      d.Get("ipsec_protocol").(string),
		Proposal:           d.Get("proposal").(string),
		Disabled:           d.Get("disabled").(bool),
	}
}

/**
 * Function used to Convert IPSec Policy to Resource Data
 */
func ipsecPolicyToData(ipsecPolicy *client.IpSecPolicy, d *schema.ResourceData) {

	// Initialize Resource ID
	d.SetId(ipsecPolicy.Id)

	// Initialize Fields
	d.Set("peer", ipsecPolicy.Peer)
	d.Set("tunnel", ipsecPolicy.Tunnel)
	d.Set("source_address", ipsecPolicy.SourceAddress)
	d.Set("source_port", ipsecPolicy.SourcePort)
	d.Set("destination_address", ipsecPolicy.DestinationAddress)
	d.Set("destination_port", ipsecPolicy.DestinationPort)
	d.Set("protocol", ipsecPolicy.Protocol)
	d.Set("template", ipsecPolicy.Template)
	d.Set("action", ipsecPolicy.Action)
	d.Set("level", ipsecPolicy.Level)
	d.Set("ipsec_protocol", ipsecPolicy.IpSecProtocol)
	d.Set("proposal", ipsecPolicy.Proposal)
	d.Set("disabled", ipsecPolicy.Disabled)
}
