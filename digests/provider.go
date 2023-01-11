package digests

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func New() provider.Provider {
	return &digestsProvider{}
}

type digestsProviderModel struct {
	AttrString    types.String `tfsdk:"attr_string"`
	NumAttr       types.Number `tfsdk:"num_attr"`
	SensitiveAttr types.String `tfsdk:"sensitive_attr"`
}

type digestsProvider struct{}

func (d *digestsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "digests"
}

func (d *digestsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "schema description",
		Attributes: map[string]schema.Attribute{
			"attr_string": schema.StringAttribute{
				Description: "provider attribute with type string",
				Optional:    true,
			},
			"num_attr": schema.NumberAttribute{
				Description: "provider attribute with type number",
				Required:    true,
			},
			"sensitive_attr": schema.StringAttribute{
				Description: "sensitive provider attribute with type string",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func (d *digestsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Hello! Configuring digests client")
	var config digestsProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.AttrString.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("attr_string"),
			"unknown attr_string",
			"provider attribute \"attr_string\" is unknown",
		)
	}

	if config.NumAttr.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("num_attr"),
			"unknown num_attr",
			"provider attribute \"num_attr\" is unknown",
		)
	}

	if config.SensitiveAttr.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("sensitive_attr"),
			"unknown sensitive_attr",
			"provider attribute \"sensitive_attr\" is unknown",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	var attrString string
	var numAttr *big.Float
	var sensitiveAttr string

	if !config.AttrString.IsNull() {
		attrString = config.AttrString.ValueString()
	}
	if !config.NumAttr.IsNull() {
		numAttr = config.NumAttr.ValueBigFloat()
	}
	if !config.SensitiveAttr.IsNull() {
		sensitiveAttr = config.SensitiveAttr.ValueString()
	}

	if attrString == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("attr_string"),
			"missing attr_string",
			"provider attribute \"attr_string\" is missing",
		)
	}

	if numAttr == nil {
		resp.Diagnostics.AddAttributeError(
			path.Root("num_attr"),
			"missing num_attr",
			"provider attribute \"num_attr\" is missing",
		)
	}

	if sensitiveAttr == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("sensitive_attr"),
			"missing sensitive_attr",
			"provider attribute \"sensitive_attr\" is missing",
		)
	}

	ctx = tflog.SetField(ctx, "digests_attr_string", attrString)
	ctx = tflog.SetField(ctx, "digests_num_attr", numAttr)
	ctx = tflog.SetField(ctx, "digests_sensitive_attr", sensitiveAttr)

	tflog.Debug(ctx, "Normal provider may create client instance here, but digests provider has no client.")

	resp.ResourceData = nil   // normal provider will assign client to ResourceData field to use it during resources creation
	resp.DataSourceData = nil // normal provider will assign client to ResourceData field to use it during data-sources reading

	tflog.Info(ctx, "digests provider complete configuration", map[string]any{"success": true})

}

func (d *digestsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewMD5DataSource,
	}
}

func (d *digestsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
