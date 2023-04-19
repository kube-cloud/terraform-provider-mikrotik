package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define Firewall Rule Resource
 */
func resourceFirewallRule() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a Firewall Rule resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createFirewallRule,

		// Read Resource Context Method CallBack
		ReadContext: readFirewallRule,

		// Update Resource Context Method CallBack
		UpdateContext: updateFirewallRule,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteFirewallRule,

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
				Description: "Firewall Rule Chain.",
			},
			"source_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Source Address.",
			},
			"source_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Source Port.",
			},
			"destination_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Destination Address.",
			},
			"destination_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Destination Port.",
			},
			"any_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Any Port.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Protocol.",
			},
			"in_interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule In Interface.",
			},
			"out_interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Out Interface.",
			},
			"in_interface_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule In Interface List.",
			},
			"out_interface_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Out Interface List.",
			},
			"packet_mark": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Packet Mark.",
			},
			"connection_mark": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Connection Mark.",
			},
			"routing_mark": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Routing Mark.",
			},
			"routing_table": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Routing Table.",
			},
			"connection_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Connection Type.",
			},
			"source_address_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Source Address List.",
			},
			"destination_address_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Destination Address List.",
			},
			"layer7_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Layer 7 Protocol.",
			},
			"source_mac_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Source Mac Address.",
			},
			"ipsec_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule IPSec Policy.",
			},
			"in_bridge_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule In Bridge Port.",
			},
			"out_bridge_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Out Bridge Port.",
			},
			"in_bridge_port_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule In Bridge Port List.",
			},
			"out_bridge_port_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Rule Out Bridge Port List.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Default:     "accept",
				Description: "Firewall Rule Action.",
			},
			"log_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Default:     "",
				Description: "Firewall Rule Log Prefix.",
			},
			"log": {
				Type:        schema.TypeBool,
				Optional:    true,
				Required:    false,
				Default:     false,
				Description: "Firewall Rule Log.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Firewall Rule Disabled.",
			},
		},
	}
}

/**
 * Create Firewall Rule from Resource Data
 */
func createFirewallRule(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get Firewall Rule Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to Firewall Rule
	dataStructure := dataToFirewallRule(d)

	// Add Firewall Rule
	firewallRule, err := c.AddFirewallRule(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert Firewall Rule to Resource Data and put it in Resource Pointer
	firewallRuleToData(firewallRule, d)

	// Reload Firewall Rule
	return readFirewallRule(ctx, d, m)
}

/**
 * Read Firewall Rule from Resource Data
 */
func readFirewallRule(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Rule Client
	c := m.(*client.Mikrotik)

	// Find Firewall Rule
	firewallRule, err := c.FindFirewallRule(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert Firewall Rule to Resource Data and put it in Resource Pointer
	firewallRuleToData(firewallRule, d)

	// Return Diagnistic
	return diags
}

/**
 * Update Firewall Rule from Resource Data
 */
func updateFirewallRule(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Rule Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to Firewall Rule
	dataStructure := dataToFirewallRule(d)

	// Update Firewall Rule
	_, err := c.UpdateFirewallRule(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete Firewall Rule from Resource Data
 */
func deleteFirewallRule(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Rule Client
	c := m.(*client.Mikrotik)

	// Delete Firewall Rule
	err := c.DeleteFirewallRule(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to Firewall Rule
 */
func dataToFirewallRule(d *schema.ResourceData) *client.FirewallRule {

	// Build and Return Firewall Rule
	return &client.FirewallRule{
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
 * Function used to Convert Firewall Rule to Resource Data
 */
func firewallRuleToData(firewallRule *client.FirewallRule, d *schema.ResourceData) {

	// Initialize Resource ID
	d.SetId(firewallRule.Id)

	// Initialize Fields
	d.Set("chain", firewallRule.Chain)
	d.Set("source_address", firewallRule.SourceAddress)
	d.Set("destination_address", firewallRule.DestinationAddress)
	d.Set("source_port", firewallRule.SourcePort)
	d.Set("destination_port", firewallRule.DestinationPort)
	d.Set("any_port", firewallRule.AnyPort)
	d.Set("protocol", firewallRule.Protocol)
	d.Set("in_interface", firewallRule.InInterface)
	d.Set("out_interface", firewallRule.OutInterface)
	d.Set("in_interface_list", firewallRule.InInterfaceList)
	d.Set("out_interface_list", firewallRule.OutInterfaceList)
	d.Set("packet_mark", firewallRule.PacketMark)
	d.Set("connection_mark", firewallRule.ConnectionMark)
	d.Set("routing_mark", firewallRule.RoutingMark)
	d.Set("routing_table", firewallRule.RoutingTable)
	d.Set("connection_type", firewallRule.ConnectionType)
	d.Set("source_address_list", firewallRule.SourceAddressList)
	d.Set("destination_address_list", firewallRule.DestinationAddressList)
	d.Set("layer7_protocol", firewallRule.Layer7Protocol)
	d.Set("source_mac_address", firewallRule.SourceMacAddress)
	d.Set("ipsec_policy", firewallRule.IpSecPolicy)
	d.Set("in_bridge_port", firewallRule.InBridgePort)
	d.Set("out_bridge_port", firewallRule.OutBridgePort)
	d.Set("in_bridge_port_list", firewallRule.InBridgePortList)
	d.Set("out_bridge_port_list", firewallRule.OutBridgePortList)
	d.Set("action", firewallRule.Action)
	d.Set("log", firewallRule.Log)
	d.Set("log_prefix", firewallRule.LogPrefix)
	d.Set("disabled", firewallRule.Disabled)
}
