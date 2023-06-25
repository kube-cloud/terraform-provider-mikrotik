package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

func resourceDhcpServerNetwork() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a DHCP network resource within Mikrotik device.",

		CreateContext: resourceDhcpServerNetworkCreate,
		ReadContext:   resourceDhcpServerNetworkRead,
		UpdateContext: resourceDhcpServerNetworkUpdate,
		DeleteContext: resourceDhcpServerNetworkDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifier of this network.",
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network DHCP server(s) will lease addresses from.",
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The DHCP client will use these as the default DNS servers.",
			},
			"gateway": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "0.0.0.0",
				Description: "The default gateway to be used by DHCP Client.",
			},
			"netmask": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "0",
				Description: "The actual network mask to be used by DHCP client. If set to '0' - netmask from network address will be used.",
			},
			"next_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The actual TFTP Server IP used by PXE Agent to continue Boot Process.",
			},
			"ntp_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The actual NTP Servers IP Addresses (as Comma Separated).",
			},
			"wins_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The actual WINS Servers IP Addresses (as Coma Separated).",
			},
			"boot_file_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The actual TFTP Boot File Name used by PXE Agent to continue Boot Process.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The actual Network Domain.",
			},
			"dhcp_option_set": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The actual DHCP Options Set (as Coma Separated).",
			},
		},
	}
}

func resourceDhcpServerNetworkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	r := dataToDhcpServerNetwork(d)
	record, err := c.AddDhcpServerNetwork(r)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(record.Id)

	return resourceDhcpServerNetworkRead(ctx, d, m)
}

func resourceDhcpServerNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	record, err := c.FindDhcpServerNetwork(d.Id())
	if _, ok := err.(*client.NotFound); ok {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}

	return dhcpServerNetworkToData(record, d)
}

func resourceDhcpServerNetworkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	r := dataToDhcpServerNetwork(d)
	_, err := c.UpdateDhcpServerNetwork(r)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceDhcpServerNetworkRead(ctx, d, m)
}

func resourceDhcpServerNetworkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	if err := c.DeleteDhcpServerNetwork(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func dataToDhcpServerNetwork(d *schema.ResourceData) *client.DhcpServerNetwork {
	r := &client.DhcpServerNetwork{}
	r.Address = d.Get("address").(string)
	r.Comment = d.Get("comment").(string)
	r.DnsServer = d.Get("dns_server").(string)
	r.Gateway = d.Get("gateway").(string)
	r.Netmask = d.Get("netmask").(string)
	r.NextServer = d.Get("next_server").(string)
	r.NtpServer = d.Get("ntp_server").(string)
	r.WinsServer = d.Get("wins_server").(string)
	r.BootFileName = d.Get("boot_file_name").(string)
	r.Domain = d.Get("domain").(string)
	r.DhcpOptionSet = d.Get("dhcp_option_set").(string)
	r.Id = d.Id()

	return r
}

func dhcpServerNetworkToData(r *client.DhcpServerNetwork, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	if err := d.Set("address", r.Address); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("comment", r.Comment); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("dns_server", r.DnsServer); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("gateway", r.Gateway); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("netmask", r.Netmask); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("next_server", r.NextServer); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("ntp_server", r.NtpServer); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("wins_server", r.WinsServer); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("boot_file_name", r.BootFileName); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("domain", r.Domain); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("dhcp_option_set", r.DhcpOptionSet); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	d.SetId(r.Id)

	return diags
}
