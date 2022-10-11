// Code generated by tools/tfsdk2fw/main.go. Manual editing is required.

package meta

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/fwtypes"
)

func init() {
	registerFWDataSourceFactory(newDataSourceARN)
}

// newDataSourceARN instantiates a new DataSource for the aws_arn data source.
func newDataSourceARN(context.Context) (datasource.DataSource, error) {
	return &dataSourceARN{}, nil
}

type dataSourceARN struct {
	meta any
}

// Metadata should return the full name of the data source, such as
// examplecloud_thing.
func (d *dataSourceARN) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) { // nosemgrep:ci.meta-in-func-name
	response.TypeName = "aws_arn"
}

// GetSchema returns the schema for this data source.
func (d *dataSourceARN) GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics) {
	schema := tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"account": {
				Type:     types.StringType,
				Computed: true,
			},
			"arn": {
				Type:     fwtypes.ARNType,
				Required: true,
			},
			"id": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
			"partition": {
				Type:     types.StringType,
				Computed: true,
			},
			"region": {
				Type:     types.StringType,
				Computed: true,
			},
			"resource": {
				Type:     types.StringType,
				Computed: true,
			},
			"service": {
				Type:     types.StringType,
				Computed: true,
			},
		},
	}

	return schema, nil
}

// Configure enables provider-level data or clients to be set in the
// provider-defined DataSource type. It is separately executed for each
// ReadDataSource RPC.
func (d *dataSourceARN) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) { //nolint:unparam
	d.meta = request.ProviderData
}

// Read is called when the provider must read data source values in order to update state.
// Config values should be read from the ReadRequest and new state values set on the ReadResponse.
func (d *dataSourceARN) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data dataSourceARNData

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	arn := &data.ARN.Value
	id := arn.String()

	data.Account = types.String{Value: arn.AccountID}
	data.ID = types.String{Value: id}
	data.Partition = types.String{Value: arn.Partition}
	data.Region = types.String{Value: arn.Region}
	data.Resource = types.String{Value: arn.Resource}
	data.Service = types.String{Value: arn.Service}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

type dataSourceARNData struct {
	Account   types.String `tfsdk:"account"`
	ARN       fwtypes.ARN  `tfsdk:"arn"`
	ID        types.String `tfsdk:"id"`
	Partition types.String `tfsdk:"partition"`
	Region    types.String `tfsdk:"region"`
	Resource  types.String `tfsdk:"resource"`
	Service   types.String `tfsdk:"service"`
}
