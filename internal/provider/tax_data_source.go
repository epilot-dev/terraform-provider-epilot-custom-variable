// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-product/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &TaxDataSource{}
var _ datasource.DataSourceWithConfigure = &TaxDataSource{}

func NewTaxDataSource() datasource.DataSource {
	return &TaxDataSource{}
}

// TaxDataSource is the data source implementation.
type TaxDataSource struct {
	client *sdk.SDK
}

// TaxDataSourceModel describes the data model.
type TaxDataSourceModel struct {
	ACL         tfTypes.BaseEntityACL     `tfsdk:"acl"`
	Active      types.Bool                `tfsdk:"active"`
	CreatedAt   types.String              `tfsdk:"created_at"`
	Description types.String              `tfsdk:"description"`
	Hydrate     types.Bool                `tfsdk:"hydrate"`
	ID          types.String              `tfsdk:"id"`
	Org         types.String              `tfsdk:"org"`
	Owners      []tfTypes.BaseEntityOwner `tfsdk:"owners"`
	Rate        types.String              `tfsdk:"rate"`
	Region      types.String              `tfsdk:"region"`
	Schema      types.String              `tfsdk:"schema"`
	Tags        []types.String            `tfsdk:"tags"`
	Title       types.String              `tfsdk:"title"`
	Type        types.String              `tfsdk:"type"`
	UpdatedAt   types.String              `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *TaxDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tax"
}

// Schema defines the schema for the data source.
func (r *TaxDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tax DataSource",

		Attributes: map[string]schema.Attribute{
			"acl": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"delete": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"edit": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"view": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
				Description: `Access control list (ACL) for an entity. Defines sharing access to external orgs or users.`,
			},
			"active": schema.BoolAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"hydrate": schema.BoolAttribute{
				Optional:    true,
				Description: `Hydrates entities in relations when passed true`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"org": schema.StringAttribute{
				Computed:    true,
				Description: `Organization Id the entity belongs to`,
			},
			"owners": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"org_id": schema.StringAttribute{
							Computed: true,
						},
						"user_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"rate": schema.StringAttribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["DE", "AT", "CH"]`,
			},
			"schema": schema.StringAttribute{
				Computed: true,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"title": schema.StringAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["VAT", "Custom"]`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *TaxDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *TaxDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TaxDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	hydrate := new(bool)
	if !data.Hydrate.IsUnknown() && !data.Hydrate.IsNull() {
		*hydrate = data.Hydrate.ValueBool()
	} else {
		hydrate = nil
	}
	var taxID string
	taxID = data.ID.ValueString()

	request := operations.GetTaxRequest{
		Hydrate: hydrate,
		TaxID:   taxID,
	}
	res, err := r.client.Tax.GetTax(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Tax != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTax(res.Tax)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
