package mikrotik

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

func resourceRecord() *schema.Resource {
	return &schema.Resource{
		Description: "Creates a DNS record on the MikroTik device.",

		CreateContext: resourceServerCreate,
		ReadContext:   resourceServerRead,
		UpdateContext: resourceServerUpdate,
		DeleteContext: resourceServerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the DNS hostname to be created.",
			},
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The A record to be returend from the DNS hostname.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The comment text associated with the DNS record.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The ttl of the DNS record.",
			},
		},
	}
}

func resourceServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	record := prepareDnsRecord(d)

	c := m.(*client.Mikrotik)

	dnsRecord, err := c.AddDnsRecord(record)
	if err != nil {
		return diag.FromErr(err)
	}

	return recordToData(dnsRecord, d)
}

func resourceServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	record, err := c.FindDnsRecord(d.Id())

	if _, ok := err.(*client.NotFound); ok {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}

	return recordToData(record, d)
}

func resourceServerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	currentRecord, err := c.FindDnsRecord(d.Id())
	record := prepareDnsRecord(d)
	record.Id = currentRecord.Id

	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] About to update dns record with %v", record)
	dnsRecord, err := c.UpdateDnsRecord(record)
	if err != nil {
		return diag.FromErr(err)
	}

	return recordToData(dnsRecord, d)
}

func resourceServerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Id()

	c := m.(*client.Mikrotik)

	record, err := c.FindDnsRecord(name)

	if err != nil {
		return diag.FromErr(err)
	}
	err = c.DeleteDnsRecord(record.Id)

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}

func recordToData(record *client.DnsRecord, d *schema.ResourceData) diag.Diagnostics {
	values := map[string]interface{}{
		"name":    record.Name,
		"address": record.Address,
		"ttl":     record.Ttl,
	}

	d.SetId(record.Name)

	var diags diag.Diagnostics

	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			diags = append(diags, diag.Errorf("failed to set %s: %v", key, err)...)
		}
	}

	return diags
}

func prepareDnsRecord(d *schema.ResourceData) *client.DnsRecord {
	dnsRecord := new(client.DnsRecord)

	dnsRecord.Name = d.Get("name").(string)
	dnsRecord.Ttl = d.Get("ttl").(int)
	dnsRecord.Address = d.Get("address").(string)
	dnsRecord.Comment = d.Get("comment").(string)

	return dnsRecord
}
