package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define Firewall Mangle Resource
 */
func resourceFirewallMangle() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a Firewall Mangle resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createFirewallMangle,

		// Read Resource Context Method CallBack
		ReadContext: readFirewallMangle,

		// Update Resource Context Method CallBack
		UpdateContext: updateFirewallMangle,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteFirewallMangle,

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
				Description: "Firewall Mangle Chain.",
			},
			"source_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Source Address.",
			},
			"source_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Source Port.",
			},
			"destination_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Destination Address.",
			},
			"destination_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Destination Port.",
			},
			"any_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Any Port.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Protocol.",
			},
			"in_interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle In Interface.",
			},
			"out_interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Out Interface.",
			},
			"in_interface_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle In Interface List.",
			},
			"out_interface_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Out Interface List.",
			},
			"packet_mark": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Packet Mark.",
			},
			"connection_mark": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Connection Mark.",
			},
			"routing_mark": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Routing Mark.",
			},
			"routing_table": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Routing Table.",
			},
			"connection_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Connection Type.",
			},
			"source_address_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Source Address List.",
			},
			"destination_address_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Destination Address List.",
			},
			"layer7_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Layer 7 Protocol.",
			},
			"source_mac_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Source Mac Address.",
			},
			"ipsec_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle IPSec Policy.",
			},
			"in_bridge_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle In Bridge Port.",
			},
			"out_bridge_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Out Bridge Port.",
			},
			"in_bridge_port_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle In Bridge Port List.",
			},
			"out_bridge_port_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Out Bridge Port List.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Default:     "accept",
				Description: "Firewall Mangle Action.",
			},
			"log_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Default:     "",
				Description: "Firewall Mangle Log Prefix.",
			},
			"log": {
				Type:        schema.TypeBool,
				Optional:    true,
				Required:    false,
				Default:     false,
				Description: "Firewall Mangle Log.",
			},
			"connection_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Connection State.",
			},
			"connection_nat_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle Connection NAT State.",
			},
			"tcp_flags": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Description: "Firewall Mangle TCP Flags.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Firewall Mangle Disabled.",
			},
		},
	}
}

/**
 * Create Firewall Mangle from Resource Data
 */
func createFirewallMangle(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get Firewall Mangle Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to Firewall Mangle
	dataStructure := dataToFirewallMangle(d)

	// Add Firewall Mangle
	firewallMangle, err := c.AddFirewallMangle(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert Firewall Mangle to Resource Data and put it in Resource Pointer
	firewallMangleToData(firewallMangle, d)

	// Reload Firewall Mangle
	return readFirewallMangle(ctx, d, m)
}

/**
 * Read Firewall Mangle from Resource Data
 */
func readFirewallMangle(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Mangle Client
	c := m.(*client.Mikrotik)

	// Find Firewall Mangle
	firewallMangle, err := c.FindFirewallMangle(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert Firewall Mangle to Resource Data and put it in Resource Pointer
	firewallMangleToData(firewallMangle, d)

	// Return Diagnistic
	return diags
}

/**
 * Update Firewall Mangle from Resource Data
 */
func updateFirewallMangle(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Mangle Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to Firewall Mangle
	dataStructure := dataToFirewallMangle(d)

	// Update Firewall Mangle
	_, err := c.UpdateFirewallMangle(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete Firewall Mangle from Resource Data
 */
func deleteFirewallMangle(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get Firewall Mangle Client
	c := m.(*client.Mikrotik)

	// Delete Firewall Mangle
	err := c.DeleteFirewallMangle(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to Firewall Mangle
 */
func dataToFirewallMangle(d *schema.ResourceData) *client.FirewallMangle {

	// Build and Return Firewall Mangle
	return &client.FirewallMangle{
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
		ConnectionState:        d.Get("connection_state").(string),
		ConnectionNatState:     d.Get("connection_nat_state").(string),
		TcpFlags:               d.Get("tcp_flags").(string),
		Action:                 d.Get("action").(string),
		Log:                    d.Get("log").(bool),
		LogPrefix:              d.Get("log_prefix").(string),
		Disabled:               d.Get("disabled").(bool),
	}
}

/**
 * Function used to Convert Firewall Mangle to Resource Data
 */
func firewallMangleToData(firewallMangle *client.FirewallMangle, d *schema.ResourceData) {

	// Initialize Resource ID
	d.SetId(firewallMangle.Id)

	// Initialize Fields
	d.Set("chain", firewallMangle.Chain)
	d.Set("source_address", firewallMangle.SourceAddress)
	d.Set("destination_address", firewallMangle.DestinationAddress)
	d.Set("source_port", firewallMangle.SourcePort)
	d.Set("destination_port", firewallMangle.DestinationPort)
	d.Set("any_port", firewallMangle.AnyPort)
	d.Set("protocol", firewallMangle.Protocol)
	d.Set("in_interface", firewallMangle.InInterface)
	d.Set("out_interface", firewallMangle.OutInterface)
	d.Set("in_interface_list", firewallMangle.InInterfaceList)
	d.Set("out_interface_list", firewallMangle.OutInterfaceList)
	d.Set("packet_mark", firewallMangle.PacketMark)
	d.Set("connection_mark", firewallMangle.ConnectionMark)
	d.Set("routing_mark", firewallMangle.RoutingMark)
	d.Set("routing_table", firewallMangle.RoutingTable)
	d.Set("connection_type", firewallMangle.ConnectionType)
	d.Set("source_address_list", firewallMangle.SourceAddressList)
	d.Set("destination_address_list", firewallMangle.DestinationAddressList)
	d.Set("layer7_protocol", firewallMangle.Layer7Protocol)
	d.Set("source_mac_address", firewallMangle.SourceMacAddress)
	d.Set("ipsec_policy", firewallMangle.IpSecPolicy)
	d.Set("in_bridge_port", firewallMangle.InBridgePort)
	d.Set("out_bridge_port", firewallMangle.OutBridgePort)
	d.Set("in_bridge_port_list", firewallMangle.InBridgePortList)
	d.Set("out_bridge_port_list", firewallMangle.OutBridgePortList)
	d.Set("connection_nat_state", firewallMangle.ConnectionNatState)
	d.Set("connection_state", firewallMangle.ConnectionState)
	d.Set("tcp_flags", firewallMangle.TcpFlags)
	d.Set("action", firewallMangle.Action)
	d.Set("log", firewallMangle.Log)
	d.Set("log_prefix", firewallMangle.LogPrefix)
	d.Set("disabled", firewallMangle.Disabled)
}
