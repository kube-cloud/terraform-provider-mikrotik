package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

func resourceBridgeInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Manages Bridge Network (VLAN) interfaces.",

		CreateContext: resourceBridgeInterfaceCreate,
		ReadContext:   resourceBridgeInterfaceRead,
		UpdateContext: resourceBridgeInterfaceUpdate,
		DeleteContext: resourceBridgeInterfaceDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
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
				Default:     false,
				Description: "Whether to create the interface in disabled state.",
			},
			"auto_mac": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Bridge Interface MAC Auto Selection Flag.",
			},
			"admin_mac": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     "",
				Description: "Bridge Interface Administration MAC.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Bridge Interface Description.",
			},
		},
	}
}

func resourceBridgeInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	record, err := c.FindBridgeInterface(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return recordBridgeInterfaceToData(record, d)
}

func resourceBridgeInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	r := dataToBridgeInterface(d)
	record, err := c.AddBridgeInterface(r)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(record.Name)

	return resourceBridgeInterfaceRead(ctx, d, m)
}

func resourceBridgeInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	existingRecord, err := c.FindBridgeInterface(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	record := dataToBridgeInterface(d)
	record.Id = existingRecord.Id
	_, err = c.UpdateBridgeInterface(record)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(record.Name)

	return nil
}

func resourceBridgeInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	err := c.DeleteBridgeInterface(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func dataToBridgeInterface(d *schema.ResourceData) *client.BridgeInterface {
	return &client.BridgeInterface{
		Mtu:      d.Get("mtu").(int),
		Name:     d.Get("name").(string),
		Disabled: d.Get("disabled").(bool),
		AutoMac:  d.Get("auto_mac").(bool),
		AdminMac: d.Get("admin_mac").(string),
		Comment:  d.Get("comment").(string),
	}
}

func recordBridgeInterfaceToData(r *client.BridgeInterface, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	if err := d.Set("mtu", r.Mtu); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("name", r.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("disabled", r.Disabled); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("auto_mac", r.AutoMac); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("admin_mac", r.AdminMac); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("comment", r.Comment); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	d.SetId(r.Name)

	return diags
}
