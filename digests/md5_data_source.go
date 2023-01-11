package digests

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

func NewMD5DataSource() datasource.DataSource {
	return &md5DataSource{}
}

type md5DataSource struct{}

func (m *md5DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_md5"
}

func (m *md5DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Calculate md5 hash",
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				Description: "Input data to hash",
				Required:    true,
			},
			"base64": schema.StringAttribute{
				Description: "Base64 encoded hash value",
				Computed:    true,
			},
			"hex": schema.StringAttribute{
				Description: "Hexadecimal hash value",
				Computed:    true,
			},
		},
	}
}

func (m *md5DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model md5DataSourceModel
	diag := req.Config.Get(ctx, &model)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := ([]byte)(model.Input.ValueString())
	hash := md5.New()
	hash.Write(input)

	rawHash := hash.Sum([]byte{})
	var base64Hash strings.Builder
	encoder := base64.NewEncoder(base64.StdEncoding, &base64Hash)
	_, err := encoder.Write(rawHash)
	if err != nil {
		resp.Diagnostics.AddError("Failed to encode MD5 hash as base64", "Failed to encode: "+err.Error())
		return
	}
	err = encoder.Close()
	if err != nil {
		resp.Diagnostics.AddError("Failed to finalize MD5 hash as base64", "Failed to finalize: "+err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}

	model.Base64 = types.StringValue(base64Hash.String())

	model.Hex = types.StringValue(fmt.Sprintf("%032x", rawHash))

	resp.State.Set(ctx, &model)

}

type md5DataSourceModel struct {
	Input  types.String `tfsdk:"input"`
	Base64 types.String `tfsdk:"base64"`
	Hex    types.String `tfsdk:"hex"`
}
