package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

/**
 * Define TFTP Resource
 */
func resourceTftp() *schema.Resource {

	// Build and Return Resource
	return &schema.Resource{

		// Resource Description
		Description: "Manages a TFTP resource within MikroTik device.",

		// Create Resource Context Method CallBack
		CreateContext: createTftp,

		// Read Resource Context Method CallBack
		ReadContext: readTftp,

		// Update Resource Context Method CallBack
		UpdateContext: updateTftp,

		// Delete Resource Context Method CallBack
		DeleteContext: deleteTftp,

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
			"ip_addresses": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Description: "TFTP Server Addresses List (comma separated).",
			},
			"request_file_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     ".*",
				Description: "File name pattern requested by PXE Clients (bios or EFI) supported by current Mikrotik TFTP Configuration (.* ==> TFTP Config Listen All Requested File name).",
			},
			"real_file_name": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Description: "Name of (boot) file actually requested by the Mikrotik router from the TFTP server (eg. pxelinux.0).",
			},
			"allow": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "TFTP Allow Flag.",
			},
			"read_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "TFTP ReadOnly Flag.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "TFTP Disabled Flag.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "TFTP Server Comment.",
			},
		},
	}
}

/**
 * Create TFTP from Resource Data
 */
func createTftp(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Get TFTP Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to TFTP
	dataStructure := dataToTftp(d)

	// Add TFTP
	tftp, err := c.AddTftp(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert TFTP to Resource Data and put it in Resource Pointer
	tftpToData(tftp, d)

	// Reload TFTP
	return readTftp(ctx, d, m)
}

/**
 * Read TFTP from Resource Data
 */
func readTftp(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get TFTP Client
	c := m.(*client.Mikrotik)

	// Find TFTP
	tftp, err := c.FindTftp(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Convert TFTP to Resource Data and put it in Resource Pointer
	tftpToData(tftp, d)

	// Return Diagnistic
	return diags
}

/**
 * Update TFTP from Resource Data
 */
func updateTftp(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get TFTP Client
	c := m.(*client.Mikrotik)

	// Convert Resource Data to TFTP
	dataStructure := dataToTftp(d)

	// Update TFTP
	_, err := c.UpdateTftp(dataStructure)

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Delete TFTP from Resource Data
 */
func deleteTftp(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Define Diagnostic variable
	var diags diag.Diagnostics

	// Get TFTP Client
	c := m.(*client.Mikrotik)

	// Delete TFTP
	err := c.DeleteTftp(d.Id())

	// If there is Error
	if err != nil {

		// Return Error
		return diag.FromErr(err)
	}

	// Return Diagnistic
	return diags
}

/**
 * Function used to Convert Resource Data to TFTP
 */
func dataToTftp(d *schema.ResourceData) *client.Tftp {

	// Build and Return TFTP
	return &client.Tftp{
		Id:              d.Id(),
		IpAddresses:     d.Get("ip_addresses").(string),
		RequestFileName: d.Get("request_file_name").(string),
		RealFileName:    d.Get("real_file_name").(string),
		Allow:           d.Get("allow").(bool),
		ReadOnly:        d.Get("read_only").(bool),
		Disabled:        d.Get("disabled").(bool),
		Comment:         d.Get("comment").(string),
	}
}

/**
 * Function used to Convert TFTP to Resource Data
 */
func tftpToData(tftp *client.Tftp, d *schema.ResourceData) {

	// Initialize Resource ID
	d.SetId(tftp.Id)

	// Initialize Fields
	d.Set("ip_addresses", tftp.IpAddresses)
	d.Set("request_file_name", tftp.RequestFileName)
	d.Set("real_file_name", tftp.RealFileName)
	d.Set("allow", tftp.Allow)
	d.Set("read_only", tftp.ReadOnly)
	d.Set("disabled", tftp.Disabled)
	d.Set("comment", tftp.Comment)
}
