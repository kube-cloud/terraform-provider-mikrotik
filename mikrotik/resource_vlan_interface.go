package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

func resourceVlanInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Manages Virtual Local Area Network (VLAN) interfaces.",

		CreateContext: resourceVlanInterfaceCreate,
		ReadContext:   resourceVlanInterfaceRead,
		UpdateContext: resourceVlanInterfaceUpdate,
		DeleteContext: resourceVlanInterfaceDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "*0",
				Description: "Name of physical interface on top of which VLAN will work.",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1500,
				Description: "Layer3 Maximum transmission unit.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Interface name.",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to create the interface in disabled state.",
			},
			"use_service_tag": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "802.1ad compatible Service Tag.",
			},
			"vlan_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Virtual LAN identifier or tag that is used to distinguish VLANs. Must be equal for all computers that belong to the same VLAN.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Virtual LAN Interface Description.",
			},
		},
	}
}

func resourceVlanInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	r := dataToVlanInterface(d)
	record, err := c.AddVlanInterface(r)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(record.Name)

	return resourceVlanInterfaceRead(ctx, d, m)
}

func resourceVlanInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	record, err := c.FindVlanInterface(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return recordVlanInterfaceToData(record, d)
}

func resourceVlanInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	existingRecord, err := c.FindVlanInterface(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	record := dataToVlanInterface(d)
	record.Id = existingRecord.Id
	_, err = c.UpdateVlanInterface(record)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(record.Name)

	return nil
}

func resourceVlanInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	err := c.DeleteVlanInterface(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func dataToVlanInterface(d *schema.ResourceData) *client.VlanInterface {
	return &client.VlanInterface{
		Interface:     d.Get("interface").(string),
		Mtu:           d.Get("mtu").(int),
		Name:          d.Get("name").(string),
		Disabled:      d.Get("disabled").(bool),
		UseServiceTag: d.Get("use_service_tag").(bool),
		VlanId:        d.Get("vlan_id").(int),
		Comment:       d.Get("comment").(string),
	}
}

func recordVlanInterfaceToData(r *client.VlanInterface, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	if err := d.Set("disabled", r.Disabled); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("interface", r.Interface); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("mtu", r.Mtu); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("name", r.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("use_service_tag", r.UseServiceTag); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("vlan_id", r.VlanId); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("comment", r.Comment); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	d.SetId(r.Name)

	return diags
}
