package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSAMLPropertyMapping() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSAMLPropertyMappingRead,
		Description: "Get SAML Property mappings",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expression": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSAMLPropertyMappingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*APIClient)

	req := c.client.PropertymappingsApi.PropertymappingsSamlList(ctx)
	if n, ok := d.GetOk("name"); ok {
		req = req.Name(n.(string))
	}
	if m, ok := d.GetOk("managed"); ok {
		req = req.Managed(m.(string))
	}
	if m, ok := d.GetOk("saml_name"); ok {
		req = req.SamlName(m.(string))
	}
	if m, ok := d.GetOk("friendly_name"); ok {
		req = req.FriendlyName(m.(string))
	}

	res, hr, err := req.Execute()
	if err != nil {
		return httpToDiag(hr, err)
	}

	if len(res.Results) < 1 {
		return diag.Errorf("No matching mappings found")
	}
	f := res.Results[0]
	d.SetId(f.Pk)
	d.Set("name", f.Name)
	d.Set("expression", f.Expression)
	d.Set("saml_name", f.SamlName)
	if f.FriendlyName.IsSet() {
		d.Set("friendly_name", f.FriendlyName.Get())
	}
	return diags
}
