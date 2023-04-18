package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define IPSec Identity Resource
 */
func resourceIpSecIdentity() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a IPSec Identity resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createIpSecIdentity,

		// Read Resource Context Method CallBack
		ReadContext: readIpSecIdentity,

		// Update Resource Context Method CallBack
		UpdateContext: updateIpSecIdentity,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteIpSecIdentity,

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
				Description: "IPSec Identity Peer.",
			},
			"auth_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "pre-shared-key",
				Description: "IPSec Identity Diffie-Hellman Group.",
			},
			"secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity Secret.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity Username.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity Password.",
			},
			"eap_methods": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity EAP Methods.",
			},
			"certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity Certificate",
			},
			"remote_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity Remote Certificate",
			},
			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity Key",
			},
			"remote_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity Remote Key",
			},
			"policy_template_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity Policy Template Group",
			},
			"no_track_chain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity No Track Chain",
			},
			"my_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "auto",
				Description: "IPSec Identity My ID",
			},
			"remote_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "auto",
				Description: "IPSec Identity Remote ID",
			},
			"match_by": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "IPSec Identity Match By",
			},
			"mode_config": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPSec Identity Mode Config",
			},
			"generate_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "no",
				Description: "IPSec Identity Generate Policy",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "IPSec Identity Comment",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "IPSec Identity Disabled",
			},
		},
	}
}

/**
 * Create IPSec Identity from Resource Data
 */
func createIpSecIdentity(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get IPSec Identity Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Identity
	dataStructure := dataToIpSecIdentity(d)

	// Add IPSec Identity
	ipsecIdentity, err := c.AddIpSecIdentity(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Identity to Resource Data and put it in Resource Pointer
	ipsecIdentityToData(ipsecIdentity, d)

	// Reload IPSec Identity
	return readIpSecIdentity(ctx, d, m)
}

/**
 * Read IPSec Identity from Resource Data
 */
func readIpSecIdentity(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Identity Client
	c := m.(*client.Mikrotik)

	// Find IPSec Identity
	ipsecIdentity, err := c.FindIpSecIdentity(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Identity to Resource Data and put it in Resource Pointer
	ipsecIdentityToData(ipsecIdentity, d)

	// Return Diagnistic
	return diags
}

/**
 * Update IPSec Identity from Resource Data
 */
func updateIpSecIdentity(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Identity Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Identity
	dataStructure := dataToIpSecIdentity(d)

	// Update IPSec Identity
	_, err := c.UpdateIpSecIdentity(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete IPSec Identity from Resource Data
 */
func deleteIpSecIdentity(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Identity Client
	c := m.(*client.Mikrotik)

	// Delete IPSec Identity
	err := c.DeleteIpSecIdentity(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to IPSec Identity
 */
func dataToIpSecIdentity(d *schema.ResourceData) *client.IpSecIdentity {

	// Build and Return IPSec Identity
	return &client.IpSecIdentity{
		Id:                  d.Id(),
		Peer:                d.Get("peer").(string),
		AuthMethod:          d.Get("auth_method").(string),
		Secret:              d.Get("secret").(string),
		Username:            d.Get("username").(string),
		Password:            d.Get("password").(string),
		EapMethods:          d.Get("eap_methods").(string),
		Certificate:         d.Get("certificate").(string),
		RemoteCertificate:   d.Get("remote_certificate").(string),
		Key:                 d.Get("key").(string),
		RemoteKey:           d.Get("remote_key").(string),
		PolicyTemplateGroup: d.Get("policy_template_group").(string),
		NoTrackChain:        d.Get("no_track_chain").(string),
		MyId:                d.Get("my_id").(string),
		RemoteId:            d.Get("remote_id").(string),
		MatchBy:             d.Get("match_by").(string),
		ModeConfig:          d.Get("mode_config").(string),
		GeneratePolicy:      d.Get("generate_policy").(string),
		Comment:             d.Get("comment").(string),
		Disabled:            d.Get("disabled").(bool),
	}
}

/**
 * Function used to Convert IPSec Identity to Resource Data
 */
func ipsecIdentityToData(ipsecIdentity *client.IpSecIdentity, d *schema.ResourceData) {

	// Initialize Resource ID
	d.SetId(ipsecIdentity.Id)

	// Initialize Fields
	d.Set("peer", ipsecIdentity.Peer)
	d.Set("auth_method", ipsecIdentity.AuthMethod)
	d.Set("secret", ipsecIdentity.Secret)
	d.Set("username", ipsecIdentity.Username)
	d.Set("password", ipsecIdentity.Password)
	d.Set("eap_methods", ipsecIdentity.EapMethods)
	d.Set("certificate", ipsecIdentity.Certificate)
	d.Set("remote_certificate", ipsecIdentity.RemoteCertificate)
	d.Set("key", ipsecIdentity.Key)
	d.Set("remote_key", ipsecIdentity.RemoteKey)
	d.Set("policy_template_group", ipsecIdentity.PolicyTemplateGroup)
	d.Set("no_track_chain", ipsecIdentity.NoTrackChain)
	d.Set("my_id", ipsecIdentity.MyId)
	d.Set("remote_id", ipsecIdentity.RemoteId)
	d.Set("match_by", ipsecIdentity.MatchBy)
	d.Set("mode_config", ipsecIdentity.ModeConfig)
	d.Set("generate_policy", ipsecIdentity.GeneratePolicy)
	d.Set("comment", ipsecIdentity.Comment)
	d.Set("disabled", ipsecIdentity.Disabled)
}
