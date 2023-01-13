package digests

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
		Description: "digests provide some more message-digest algorithms",
		Attributes:  map[string]schema.Attribute{},
	}
}

func (d *digestsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	tflog.Info(ctx, "digests provider complete configuration", map[string]any{"success": true})

}

func (d *digestsProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewMD4DataSource,
		NewMD5DataSource,
		NewSHA1DataSource,
		NewSHA224DataSource,
		NewSHA256DataSource,
		NewSHA384DataSource,
		NewSHA512DataSource,
		NewSHA512_224DataSource,
		NewSHA512_256DataSource,
		NewSHA3_224DataSource,
		NewSHA3_256DataSource,
		NewSHA3_384DataSource,
		NewSHA3_512DataSource,
		NewCRC32IEEEDataSource,
		NewCRC32CastagnoliDataSource,
		NewCRC32KoopmanDataSource,
	}
}

func (d *digestsProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
