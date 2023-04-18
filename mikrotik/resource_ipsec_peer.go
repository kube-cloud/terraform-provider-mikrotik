package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define IPSec Peer Resource
 */
func resourceIpSecPeer() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a IPSec Peer resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createIpSecPeer,

		// Read Resource Context Method CallBack
		ReadContext: readIpSecPeer,

		// Update Resource Context Method CallBack
		UpdateContext: updateIpSecPeer,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteIpSecPeer,

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
				Description: "IPSec Peer Name.",
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Description: "IPSec Peer Address.",
			},
			"profile": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Description: "IPSec Peer Profile.",
			},
			"exchange_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ike2",
				Description: "IPSec Peer Max Failure (in minute).",
			},
			"send_initial_contact": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "IPSec Peer Send Initial Contact.",
			},
			"passive": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "IPSec Peer Passive.",
			},
			"local_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "IPSec Peer Local Address",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "IPSec Peer Port",
			},
		},
	}
}

/**
 * Create IPSec Peer from Resource Data
 */
func createIpSecPeer(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get IPSec Peer Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Peer
	dataStructure := dataToIpSecPeer(d)

	// Add IPSec Peer
	ipsecPeer, err := c.AddIpSecPeer(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Peer to Resource Data and put it in Resource Pointer
	ipsecPeerToData(ipsecPeer, d)

	// Reload IPSec Peer
	return readIpSecPeer(ctx, d, m)
}

/**
 * Read IPSec Peer from Resource Data
 */
func readIpSecPeer(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Peer Client
	c := m.(*client.Mikrotik)

	// Find IPSec Peer
	ipsecPeer, err := c.FindIpSecPeer(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Peer to Resource Data and put it in Resource Pointer
	ipsecPeerToData(ipsecPeer, d)

	// Return Diagnistic
	return diags
}

/**
 * Update IPSec Peer from Resource Data
 */
func updateIpSecPeer(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Peer Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Peer
	dataStructure := dataToIpSecPeer(d)

	// Update IPSec Peer
	_, err := c.UpdateIpSecPeer(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete IPSec Peer from Resource Data
 */
func deleteIpSecPeer(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Peer Client
	c := m.(*client.Mikrotik)

	// Delete IPSec Peer
	err := c.DeleteIpSecPeer(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to IPSec Peer
 */
func dataToIpSecPeer(d *schema.ResourceData) *client.IpSecPeer {

	// Build and Return IPSec Peer
	return &client.IpSecPeer{
		Id:                 d.Id(),
		Name:               d.Get("name").(string),
		Address:            d.Get("address").(string),
		Profile:            d.Get("profile").(string),
		ExchangeMode:       d.Get("exchange_mode").(string),
		SendInitialContact: d.Get("send_initial_contact").(bool),
		Passive:            d.Get("passive").(bool),
		LocalAddress:       d.Get("local_address").(string),
		Port:               d.Get("port").(int),
	}
}

/**
 * Function used to Convert IPSec Peer to Resource Data
 */
func ipsecPeerToData(ipsecPeer *client.IpSecPeer, d *schema.ResourceData) {

	// Initialize Resource ID
	d.SetId(ipsecPeer.Id)

	// Initialize Fields
	d.Set("name", ipsecPeer.Name)
	d.Set("address", ipsecPeer.Address)
	d.Set("profile", ipsecPeer.Profile)
	d.Set("exchange_mode", ipsecPeer.ExchangeMode)
	d.Set("send_initial_contact", ipsecPeer.SendInitialContact)
	d.Set("passive", ipsecPeer.Passive)
	d.Set("local_address", ipsecPeer.LocalAddress)
	d.Set("port", ipsecPeer.Port)
}
