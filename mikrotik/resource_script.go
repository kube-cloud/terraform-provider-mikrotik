package mikrotik

import (
	"context"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScript() *schema.Resource {
	return &schema.Resource{
		Description: "Creates a MikroTik script.",

		CreateContext: resourceScriptCreate,
		ReadContext:   resourceScriptRead,
		UpdateContext: resourceScriptUpdate,
		DeleteContext: resourceScriptDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of script.",
			},
			"owner": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The owner of the script.",
			},
			"source": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source code of the script. See the [MikroTik docs](https://wiki.mikrotik.com/wiki/Manual:Scripting) for the scripting language.",
			},
			"policy": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "What permissions the script has. This must be one of the following: ftp, reboot, read, write, policy, test, password, sniff, sensitive, romon.",
			},
			"dont_require_permissions": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If the script requires permissions or not.",
			},
		},
	}
}

func resourceScriptCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	owner := d.Get("owner").(string)
	source := d.Get("source").(string)
	policy := d.Get("policy").([]interface{})
	policies := []string{}
	for _, p := range policy {
		policies = append(policies, p.(string))
	}
	dontReqPerms := d.Get("dont_require_permissions").(bool)

	c := m.(*client.Mikrotik)

	script, err := c.CreateScript(
		name,
		owner,
		source,
		policies,
		dontReqPerms,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	return scriptToData(script, d)
}

func scriptToData(s *client.Script, d *schema.ResourceData) diag.Diagnostics {
	values := map[string]interface{}{
		"name":                     s.Name,
		"owner":                    s.Owner,
		"source":                   s.Source,
		"policy":                   s.Policy(),
		"dont_require_permissions": s.DontRequirePermissions,
	}

	d.SetId(s.Name)

	var diags diag.Diagnostics

	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			diags = append(diags, diag.Errorf("failed to set %s: %v", key, err)...)
		}
	}

	return diags
}

func resourceScriptRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Mikrotik)

	script, err := c.FindScript(d.Id())

	if _, ok := err.(*client.NotFound); ok {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}

	return scriptToData(script, d)
}
func resourceScriptUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	owner := d.Get("owner").(string)
	source := d.Get("source").(string)
	dontReqPerms := d.Get("dont_require_permissions").(bool)
	policy := d.Get("policy").([]interface{})
	policies := []string{}
	for _, p := range policy {
		str, ok := p.(string)
		if ok {
			policies = append(policies, str)
		}
	}

	c := m.(*client.Mikrotik)

	script, err := c.UpdateScript(name, owner, source, policies, dontReqPerms)
	if err != nil {
		return diag.FromErr(err)
	}

	return scriptToData(script, d)
}
func resourceScriptDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Id()

	c := m.(*client.Mikrotik)

	err := c.DeleteScript(name)

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
