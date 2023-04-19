package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define Firewall Nat Resource
 */
func resourceFirewallNat() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a Firewall Nat resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createFirewallNat,

		// Read Resource Context Method CallBack
		ReadContext: readFirewallNat,

		// Update Resource Context Method CallBack
		UpdateContext: updateFirewallNat,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteFirewallNat,

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
				Description: "Firewall Nat Chain.",
			},
			"source_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Source Address.",
			},
			"source_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Source Port.",
			},
			"destination_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Destination Address.",
			},
			"destination_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Destination Port.",
			},
			"any_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Any Port.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Protocol.",
			},
			"in_interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat In Interface.",
			},
			"out_interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Out Interface.",
			},
			"in_interface_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat In Interface List.",
			},
			"out_interface_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Out Interface List.",
			},
			"packet_mark": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Packet Mark.",
			},
			"connection_mark": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Connection Mark.",
			},
			"routing_mark": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Routing Mark.",
			},
			"routing_table": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Routing Table.",
			},
			"connection_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Connection Type.",
			},
			"source_address_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Source Address List.",
			},
			"destination_address_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Destination Address List.",
			},
			"layer7_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Layer 7 Protocol.",
			},
			"source_mac_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Source Mac Address.",
			},
			"ipsec_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat IPSec Policy.",
			},
			"in_bridge_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat In Bridge Port.",
			},
			"out_bridge_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Out Bridge Port.",
			},
			"in_bridge_port_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat In Bridge Port List.",
			},
			"out_bridge_port_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Nat Out Bridge Port List.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Default:     "accept",
				Description: "Firewall Nat Action.",
			},
			"log_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Default:     "",
				Description: "Firewall Nat Log Prefix.",
			},
			"log": {
				Type:        schema.TypeBool,
				Optional:    true,
				Required:    false,
				Default:     false,
				Description: "Firewall Nat Log.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Firewall Nat Disabled.",
			},
		},
	}
}

/**
 * Create Firewall Nat from Resource Data
 */
func createFirewallNat(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get Firewall Nat Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to Firewall Nat
	dataStructure := dataToFirewallNat(d)

	// Add Firewall Nat
	firewallNat, err := c.AddFirewallNat(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert Firewall Nat to Resource Data and put it in Resource Pointer
	firewallNatToData(firewallNat, d)

	// Reload Firewall Nat
	return readFirewallNat(ctx, d, m)
}

/**
 * Read Firewall Nat from Resource Data
 */
func readFirewallNat(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Nat Client
	c := m.(*client.Mikrotik)

	// Find Firewall Nat
	firewallNat, err := c.FindFirewallNat(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert Firewall Nat to Resource Data and put it in Resource Pointer
	firewallNatToData(firewallNat, d)

	// Return Diagnistic
	return diags
}

/**
 * Update Firewall Nat from Resource Data
 */
func updateFirewallNat(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Nat Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to Firewall Nat
	dataStructure := dataToFirewallNat(d)

	// Update Firewall Nat
	_, err := c.UpdateFirewallNat(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete Firewall Nat from Resource Data
 */
func deleteFirewallNat(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Nat Client
	c := m.(*client.Mikrotik)

	// Delete Firewall Nat
	err := c.DeleteFirewallNat(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to Firewall Nat
 */
func dataToFirewallNat(d *schema.ResourceData) *client.FirewallNat {

	// Build and Return Firewall Nat
	return &client.FirewallNat{
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
		PacketMark:             d.Get("packet_mark").(string),
		ConnectionMark:         d.Get("connection_mark").(string),
		RoutingMark:            d.Get("routing_mark").(string),
		RoutingTable:           d.Get("routing_table").(string),
		ConnectionType:         d.Get("connection_type").(string),
		SourceAddressList:      d.Get("source_address_list").(string),
		DestinationAddressList: d.Get("destination_address_list").(string),
		Layer7Protocol:         d.Get("layer7_protocol").(string),
		SourceMacAddress:       d.Get("source_mac_address").(string),
		IpSecPolicy:            d.Get("ipsec_policy").(string),
		InBridgePort:           d.Get("in_bridge_port").(string),
		OutBridgePort:          d.Get("out_bridge_port").(string),
		InBridgePortList:       d.Get("in_bridge_port_list").(string),
		OutBridgePortList:      d.Get("out_bridge_port_list").(string),
		Action:                 d.Get("action").(string),
		Log:                    d.Get("log").(bool),
		LogPrefix:              d.Get("log_prefix").(string),
		Disabled:               d.Get("disabled").(bool),
	}
}

/**
 * Function used to Convert Firewall Nat to Resource Data
 */
func firewallNatToData(firewallNat *client.FirewallNat, d *schema.ResourceData) {

	// Initialize Resource ID
	d.SetId(firewallNat.Id)

	// Initialize Fields
	d.Set("chain", firewallNat.Chain)
	d.Set("source_address", firewallNat.SourceAddress)
	d.Set("destination_address", firewallNat.DestinationAddress)
	d.Set("source_port", firewallNat.SourcePort)
	d.Set("destination_port", firewallNat.DestinationPort)
	d.Set("any_port", firewallNat.AnyPort)
	d.Set("protocol", firewallNat.Protocol)
	d.Set("in_interface", firewallNat.InInterface)
	d.Set("out_interface", firewallNat.OutInterface)
	d.Set("in_interface_list", firewallNat.InInterfaceList)
	d.Set("out_interface_list", firewallNat.OutInterfaceList)
	d.Set("packet_mark", firewallNat.PacketMark)
	d.Set("connection_mark", firewallNat.ConnectionMark)
	d.Set("routing_mark", firewallNat.RoutingMark)
	d.Set("routing_table", firewallNat.RoutingTable)
	d.Set("connection_type", firewallNat.ConnectionType)
	d.Set("source_address_list", firewallNat.SourceAddressList)
	d.Set("destination_address_list", firewallNat.DestinationAddressList)
	d.Set("layer7_protocol", firewallNat.Layer7Protocol)
	d.Set("source_mac_address", firewallNat.SourceMacAddress)
	d.Set("ipsec_policy", firewallNat.IpSecPolicy)
	d.Set("in_bridge_port", firewallNat.InBridgePort)
	d.Set("out_bridge_port", firewallNat.OutBridgePort)
	d.Set("in_bridge_port_list", firewallNat.InBridgePortList)
	d.Set("out_bridge_port_list", firewallNat.OutBridgePortList)
	d.Set("action", firewallNat.Action)
	d.Set("log", firewallNat.Log)
	d.Set("log_prefix", firewallNat.LogPrefix)
	d.Set("disabled", firewallNat.Disabled)
}
