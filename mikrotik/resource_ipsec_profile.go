package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define IPSec Profile Resource
 */
func resourceIpSecProfile() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a IPSec Profile resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createIpSecProfile,

		// Read Resource Context Method CallBack
		ReadContext: readIpSecProfile,

		// Update Resource Context Method CallBack
		UpdateContext: updateIpSecProfile,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteIpSecProfile,

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
				Description: "IPSec Profile Name.",
			},
			"dh_group": {
				Type:     schema.TypeString,
				Optional: true,

				// Possible : ecp256,ecp384,ecp521,ec2n185,ec2n155,modp8192,modp6144,modp4096,modp3072,modp2048,modp1536,modp1024,modp768
				Default:     "modp2048,modp3072,modp1536",
				Description: "IPSec Profile Diffie-Hellman Group.",
			},
			"dpd_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "2m",
				Description: "IPSec Profile DPD Interval.",
			},
			"dpd_max_failure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     5,
				Description: "IPSec Profile Max Failure (in minute).",
			},
			"enc_algorithms": {
				Type:     schema.TypeString,
				Optional: true,

				// Possible : aes-256,camellia-256,aes-192,camellia-192,aes-128,camellia-128,3des,blowfish,des
				Default:     "aes-256,aes-128",
				Description: "IPSec Profile Encryption Algorithms.",
			},
			"hash_algorithm": {
				Type:     schema.TypeString,
				Optional: true,

				// Possible : sha1,sha256,sha384,md5
				Default:     "sha1",
				Description: "IPSec Profile Hash Algorithm.",
			},
			"lifetime": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "1h30m",
				Description: "IPSec Profile Lifetime",
			},
			"nat_traversal": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "IPSec Profile NAT Transversal",
			},
			"proposal_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "obey",
				Description: "IPSec Profile Lifetime",
			},
		},
	}
}

/**
 * Create IPSec Profile from Resource Data
 */
func createIpSecProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get IPSec Profile Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Profile
	dataStructure := dataToIpSecProfile(d)

	// Add IPSec Profile
	ipsecProfile, err := c.AddIpSecProfile(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Profile to Resource Data and put it in Resource Pointer
	ipsecProfileToData(ipsecProfile, d)

	// Reload IPSec Profile
	return readIpSecProfile(ctx, d, m)
}

/**
 * Read IPSec Profile from Resource Data
 */
func readIpSecProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Profile Client
	c := m.(*client.Mikrotik)

	// Find IPSec Profile
	ipsecProfile, err := c.FindIpSecProfile(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert IPSec Profile to Resource Data and put it in Resource Pointer
	ipsecProfileToData(ipsecProfile, d)

	// Return Diagnistic
	return diags
}

/**
 * Update IPSec Profile from Resource Data
 */
func updateIpSecProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Profile Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to IPSec Profile
	dataStructure := dataToIpSecProfile(d)

	// Update IPSec Profile
	_, err := c.UpdateIpSecProfile(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete IPSec Profile from Resource Data
 */
func deleteIpSecProfile(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get IPSec Profile Client
	c := m.(*client.Mikrotik)

	// Delete IPSec Profile
	err := c.DeleteIpSecProfile(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to IPSec Profile
 */
func dataToIpSecProfile(d *schema.ResourceData) *client.IpSecProfile {

	// Build and Return IPSec Profile
	return &client.IpSecProfile{
		Id:            d.Id(),
		Name:          d.Get("name").(string),
		DhGroup:       d.Get("dh_group").(string),
		DpdInterval:   d.Get("dpd_interval").(string),
		DpdMaxFailure: d.Get("dpd_max_failure").(int),
		EncAlgorithms: d.Get("enc_algorithms").(string),
		HashAlgorithm: d.Get("hash_algorithm").(string),
		Lifetime:      d.Get("lifetime").(string),
		NatTraversal:  d.Get("nat_traversal").(bool),
		ProposalCheck: d.Get("proposal_check").(string),
	}
}

/**
 * Function used to Convert IPSec Profile to Resource Data
 */
func ipsecProfileToData(ipsecProfile *client.IpSecProfile, d *schema.ResourceData) {

	// Initialize Resource ID
	d.SetId(ipsecProfile.Id)

	// Initialize Fields
	d.Set("name", ipsecProfile.Name)
	d.Set("dh_group", ipsecProfile.DhGroup)
	d.Set("dpd_interval", ipsecProfile.DpdInterval)
	d.Set("dpd_max_failure", ipsecProfile.DpdMaxFailure)
	d.Set("enc_algorithms", ipsecProfile.EncAlgorithms)
	d.Set("hash_algorithm", ipsecProfile.HashAlgorithm)
	d.Set("lifetime", ipsecProfile.Lifetime)
	d.Set("nat_traversal", ipsecProfile.NatTraversal)
	d.Set("proposal_check", ipsecProfile.ProposalCheck)
}
