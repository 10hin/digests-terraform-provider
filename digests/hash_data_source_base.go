package digests

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

type hashDataSourceBase struct {
	typeNameSuffix string
	algorithmName  string
	hash           func([]byte) []byte
}

func (b *hashDataSourceBase) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + b.typeNameSuffix
}

func (b *hashDataSourceBase) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: fmt.Sprintf("Calculate %s hash", b.algorithmName),
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

func (b *hashDataSourceBase) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model hashDataSourceModel
	diag := req.Config.Get(ctx, &model)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := ([]byte)(model.Input.ValueString())

	rawHash := b.hash(input)
	var base64Hash strings.Builder
	encoder := base64.NewEncoder(base64.StdEncoding, &base64Hash)
	_, err := encoder.Write(rawHash)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to encode %s hash as base64", b.algorithmName), "Failed to encode: "+err.Error())
		return
	}
	err = encoder.Close()
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Failed to finalize encoding %s hash as base64", b.algorithmName), "Failed to finalize: "+err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}

	model.Base64 = types.StringValue(base64Hash.String())

	format := fmt.Sprintf("%%0%dx", len(rawHash)*2)
	model.Hex = types.StringValue(fmt.Sprintf(format, rawHash))

	resp.State.Set(ctx, &model)

}

type hashDataSourceModel struct {
	Input  types.String `tfsdk:"input"`
	Base64 types.String `tfsdk:"base64"`
	Hex    types.String `tfsdk:"hex"`
}
