package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define IPSec Proposal Resource
 */
func resourceIpSecProposal() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a IPSec Proposal resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createIpSecProposal,

		// Read Resource Context Method CallBack
		ReadContext: readIpSecProposal,

		// Update Resource Context Method CallBack
		UpdateContext: updateIpSecProposal,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteIpSecProposal,

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
				Description: "IPSec Proposal Name.",
			},
			"auth_algorithms": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "sha512,sha256,sha1,md5",
				Description: "IPSec Proposal Authentication Algorithms List.",
			},
			"enc_algorithms": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "aes-256-cbc,aes-256-ctr,aes-256-gcm,camellia-256,aes-192-cbc,aes-192-ctr,aes-192-gcm,camellia-192,aes-128-cbc,aes-128-ctr,aes-128-gcm,camellia-128,3des,blowfish,twofish,des",
				Description: "IPSec Proposal Encryption Algorithms List.",
			},
			"lifetime": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "30m",
				Description: "IPSec Proposal Lifetime.",
			},
			"pfs_group": {
				Type:     schema.TypeString,
				Optional: true,

				// Possible : none, modp1024, modp1536, modp2048, modp3072, modp4096, modp6144, modp8192, ecp521, ecp384, ecp256
				Default:     "modp1024",
				Description: "IPSec Proposal PSF Group.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "IPSec Proposal Desabled.",
			},
		},
	}
}

/**
 * Create IPSec Proposal from Resource Data
 */
func createIpSecProposal(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get IPSec Proposal Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Proposal
	dataStructure := dataToIpSecProposal(d)

	// Add IPSec Proposal
	ipsecProposal, err := c.AddIpSecProposal(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Proposal to Resource Data and put it in Resource Pointer
	ipsecProposalToData(ipsecProposal, d)

	// Reload IPSec Proposal
	return readIpSecProposal(ctx, d, m)
}

/**
 * Read IPSec Proposal from Resource Data
 */
func readIpSecProposal(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Proposal Client
	c := m.(*client.Mikrotik)

	// Find IPSec Proposal
	ipsecProposal, err := c.FindIpSecProposal(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Proposal to Resource Data and put it in Resource Pointer
	ipsecProposalToData(ipsecProposal, d)

	// Return Diagnistic
	return diags
}

/**
 * Update IPSec Proposal from Resource Data
 */
func updateIpSecProposal(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Proposal Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Proposal
	dataStructure := dataToIpSecProposal(d)

	// Update IPSec Proposal
	_, err := c.UpdateIpSecProposal(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete IPSec Proposal from Resource Data
 */
func deleteIpSecProposal(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Proposal Client
	c := m.(*client.Mikrotik)

	// Delete IPSec Proposal
	err := c.DeleteIpSecProposal(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to IPSec Proposal
 */
func dataToIpSecProposal(d *schema.ResourceData) *client.IpSecProposal {

	//Build and Return IPSec Proposal
	return &client.IpSecProposal{
		Id:             d.Id(),
		Name:           d.Get("name").(string),
		AuthAlgorithms: d.Get("auth_algorithms").(string),
		EncAlgorithms:  d.Get("enc_algorithms").(string),
		Lifetime:       d.Get("lifetime").(string),
		PfsGroup:       d.Get("pfs_group").(string),
		Disabled:       d.Get("disabled").(bool),
	}
}

/**
 * Function used to Convert IPSec Proposal to Resource Data
 */
func ipsecProposalToData(ipsecProposal *client.IpSecProposal, d *schema.ResourceData) {

	// Initialize Fields
	d.Set("name", ipsecProposal.Name)
	d.Set("auth_algorithms", ipsecProposal.AuthAlgorithms)
	d.Set("enc_algorithms", ipsecProposal.EncAlgorithms)
	d.Set("lifetime", ipsecProposal.Lifetime)
	d.Set("pfs_group", ipsecProposal.PfsGroup)
	d.Set("disabled", ipsecProposal.Disabled)
}
