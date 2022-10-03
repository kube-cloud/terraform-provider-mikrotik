package mikrotik

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kube-cloud/terraform-provider-mikrotik/client"
)

func resourceBridgeInterfacePort() *schema.Resource {
	return &schema.Resource{
		Description: "Manages Bridge Network Interfaces Ports.",

		CreateContext: resourceBridgeInterfacePortCreate,
		ReadContext:   resourceBridgeInterfacePortRead,
		UpdateContext: resourceBridgeInterfacePortUpdate,
		DeleteContext: resourceBridgeInterfacePortDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"bridge": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bridge Interface.",
			},
			"interface": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Network Interface Name.",
			},
			"horizon": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "none", // Other Values : 0, 1, 2, etc...
				Description: "Bridge Port  Horizon (Values : none|Integer).",
			},
			"learn": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "auto", // Other Values : yes, no
				Description: "Bridge Port  Learn (Values : auto|yes|no).",
			},
			"unknown_multicast_flood": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Bridge Port Unknown Multicast Flood.",
			},
			"unknown_unicast_flood": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Bridge Port Unknown Unicast Flood.",
			},
			"broadcast_flood": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Bridge Port Boradcast Flood.",
			},
			"trusted": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Bridge Port Trusted.",
			},
			"hardware_offload": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Bridge Port Hardware Offload.",
			},
			"auto_isolate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Bridge Port Auto Isolate.",
			},
			"restricted_role": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Bridge Port Restricted Role.",
			},
			"restricted_tcn": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Bridge Port Restricted TCN.",
			},
			"bpdu_guard": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Bridge Port BPDU Guard.",
			},
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0x80,
				Description: "Bridge Port Priority (In Hexadecimal).",
			},
			"path_cost": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
				Description: "Bridge Port Path Cost.",
			},
			"internal_path_cost": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
				Description: "Bridge Port Internal Path Cost.",
			},
			"edge": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "auto",
				Description: "Bridge Port Edge (Values : auto|no|no-discover|yes|yes-discover).",
			},
			"point_to_point": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "auto",
				Description: "Bridge Port Point To Point (Values : auto|yes|no).",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Bridge Port Disabled.",
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

func resourceBridgeInterfacePortRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	record, err := c.FindBridgeInterfacePort(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return recordBridgeInterfacePortToData(record, d)
}

func resourceBridgeInterfacePortCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	c := m.(*client.Mikrotik)
	r := dataToBridgeInterfacePort(d)
	record, err := c.AddBridgeInterfacePort(r)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(record.Interface)

	return resourceBridgeInterfacePortRead(ctx, d, m)
}

func resourceBridgeInterfacePortUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	existingRecord, err := c.FindBridgeInterfacePort(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	record := dataToBridgeInterfacePort(d)
	record.Id = existingRecord.Id
	_, err = c.UpdateBridgeInterfacePort(record)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(record.Interface)

	return nil
}

func resourceBridgeInterfacePortDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)
	err := c.DeleteBridgeInterfacePort(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func dataToBridgeInterfacePort(d *schema.ResourceData) *client.BridgeInterfacePort {
	return &client.BridgeInterfacePort{
		Bridge:                d.Get("bridge").(string),
		Interface:             d.Get("interface").(string),
		Horizon:               d.Get("horizon").(string),
		Learn:                 d.Get("learn").(string),
		UnknownMulticastFlood: d.Get("unknown_multicast_flood").(bool),
		UnknownUnicastFlood:   d.Get("unknown_unicast_flood").(bool),
		BroadcastFlood:        d.Get("broadcast_flood").(bool),
		Trusted:               d.Get("trusted").(bool),
		HardwareOffload:       d.Get("hardware_offload").(bool),
		AutoIsolate:           d.Get("auto_isolate").(bool),
		RestrictedRole:        d.Get("restricted_role").(bool),
		RestrictedTcn:         d.Get("restricted_tcn").(bool),
		BpduGuard:             d.Get("bpdu_guard").(bool),
		Priority:              d.Get("priority").(int),
		PathCost:              d.Get("path_cost").(int),
		InternalPathCost:      d.Get("internal_path_cost").(int),
		Edge:                  d.Get("edge").(string),
		PointToPoint:          d.Get("point_to_point").(string),
		Disabled:              d.Get("disabled").(bool),
		Comment:               d.Get("comment").(string),
	}
}

func recordBridgeInterfacePortToData(r *client.BridgeInterfacePort, d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	if err := d.Set("bridge", r.Bridge); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("interface", r.Interface); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("horizon", r.Horizon); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("learn", r.Learn); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("unknown_multicast_flood", r.UnknownMulticastFlood); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("unknown_unicast_flood", r.UnknownUnicastFlood); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("broadcast_flood", r.BroadcastFlood); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("trusted", r.Trusted); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("hardware_offload", r.HardwareOffload); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("auto_isolate", r.AutoIsolate); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("restricted_role", r.RestrictedRole); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("restricted_tcn", r.RestrictedTcn); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("bpdu_guard", r.BpduGuard); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("priority", r.Priority); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("path_cost", r.PathCost); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("internal_path_cost", r.InternalPathCost); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("edge", r.Edge); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("point_to_point", r.PointToPoint); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("disabled", r.Disabled); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if err := d.Set("comment", r.Comment); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	d.SetId(r.Interface)

	return diags
}
