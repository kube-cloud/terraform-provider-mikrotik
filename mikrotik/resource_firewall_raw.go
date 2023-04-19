package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define Firewall Raw Resource
 */
func resourceFirewallRaw() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a Firewall Raw resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createFirewallRaw,

		// Read Resource Context Method CallBack
		ReadContext: readFirewallRaw,

		// Update Resource Context Method CallBack
		UpdateContext: updateFirewallRaw,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteFirewallRaw,

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
			"chain": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Description: "Firewall Raw Chain.",
			},
			"source_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Source Address.",
			},
			"source_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Source Port.",
			},
			"destination_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Destination Address.",
			},
			"destination_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Destination Port.",
			},
			"any_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Any Port.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Protocol.",
			},
			"in_interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw In Interface.",
			},
			"out_interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Out Interface.",
			},
			"in_interface_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw In Interface List.",
			},
			"out_interface_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Out Interface List.",
			},
			"source_address_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Source Address List.",
			},
			"destination_address_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Destination Address List.",
			},
			"source_mac_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw Source Mac Address.",
			},
			"ipsec_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Raw IPSec Policy.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Default:     "accept",
				Description: "Firewall Raw Action.",
			},
			"log_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Default:     "",
				Description: "Firewall Raw Log Prefix.",
			},
			"log": {
				Type:        schema.TypeBool,
				Optional:    true,
				Required:    false,
				Default:     false,
				Description: "Firewall Raw Log.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Firewall Raw Disabled.",
			},
		},
	}
}

/**
 * Create Firewall Raw from Resource Data
 */
func createFirewallRaw(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get Firewall Raw Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to Firewall Raw
	dataStructure := dataToFirewallRaw(d)

	// Add Firewall Raw
	firewallRaw, err := c.AddFirewallRaw(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert Firewall Raw to Resource Data and put it in Resource Pointer
	firewallRawToData(firewallRaw, d)

	// Reload Firewall Raw
	return readFirewallRaw(ctx, d, m)
}

/**
 * Read Firewall Raw from Resource Data
 */
func readFirewallRaw(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Raw Client
	c := m.(*client.Mikrotik)

	// Find Firewall Raw
	firewallRaw, err := c.FindFirewallRaw(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert Firewall Raw to Resource Data and put it in Resource Pointer
	firewallRawToData(firewallRaw, d)

	// Return Diagnistic
	return diags
}

/**
 * Update Firewall Raw from Resource Data
 */
func updateFirewallRaw(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Raw Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to Firewall Raw
	dataStructure := dataToFirewallRaw(d)

	// Update Firewall Raw
	_, err := c.UpdateFirewallRaw(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete Firewall Raw from Resource Data
 */
func deleteFirewallRaw(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Raw Client
	c := m.(*client.Mikrotik)

	// Delete Firewall Raw
	err := c.DeleteFirewallRaw(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to Firewall Raw
 */
func dataToFirewallRaw(d *schema.ResourceData) *client.FirewallRaw {

	// Build and Return Firewall Raw
	return &client.FirewallRaw{
		Id:                     d.Id(),
		Chain:                  d.Get("chain").(string),
		SourceAddress:          d.Get("source_address").(string),
		DestinationAddress:     d.Get("destination_address").(string),
		SourcePort:             d.Get("source_port").(int),
		DestinationPort:        d.Get("destination_port").(int),
		AnyPort:                d.Get("any_port").(int),
		Protocol:               d.Get("protocol").(string),
		InInterface:            d.Get("in_interface").(string),
		OutInterface:           d.Get("out_interface").(string),
		InInterfaceList:        d.Get("in_interface_list").(string),
		OutInterfaceList:       d.Get("out_interface_list").(string),
		SourceAddressList:      d.Get("source_address_list").(string),
		DestinationAddressList: d.Get("destination_address_list").(string),
		SourceMacAddress:       d.Get("source_mac_address").(string),
		IpSecPolicy:            d.Get("ipsec_policy").(string),
		Action:                 d.Get("action").(string),
		Log:                    d.Get("log").(bool),
		LogPrefix:              d.Get("log_prefix").(string),
		Disabled:               d.Get("disabled").(bool),
	}
}

/**
 * Function used to Convert Firewall Raw to Resource Data
 */
func firewallRawToData(firewallRaw *client.FirewallRaw, d *schema.ResourceData) {

	// Initialize Resource ID
	d.SetId(firewallRaw.Id)

	// Initialize Fields
	d.Set("chain", firewallRaw.Chain)
	d.Set("source_address", firewallRaw.SourceAddress)
	d.Set("destination_address", firewallRaw.DestinationAddress)
	d.Set("source_port", firewallRaw.SourcePort)
	d.Set("destination_port", firewallRaw.DestinationPort)
	d.Set("any_port", firewallRaw.AnyPort)
	d.Set("protocol", firewallRaw.Protocol)
	d.Set("in_interface", firewallRaw.InInterface)
	d.Set("out_interface", firewallRaw.OutInterface)
	d.Set("in_interface_list", firewallRaw.InInterfaceList)
	d.Set("out_interface_list", firewallRaw.OutInterfaceList)
	d.Set("source_address_list", firewallRaw.SourceAddressList)
	d.Set("destination_address_list", firewallRaw.DestinationAddressList)
	d.Set("source_mac_address", firewallRaw.SourceMacAddress)
	d.Set("ipsec_policy", firewallRaw.IpSecPolicy)
	d.Set("action", firewallRaw.Action)
	d.Set("log", firewallRaw.Log)
	d.Set("log_prefix", firewallRaw.LogPrefix)
	d.Set("disabled", firewallRaw.Disabled)
}
