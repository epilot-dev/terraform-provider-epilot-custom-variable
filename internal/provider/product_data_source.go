// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ProductDataSource{}
var _ datasource.DataSourceWithConfigure = &ProductDataSource{}

func NewProductDataSource() datasource.DataSource {
	return &ProductDataSource{}
}

// ProductDataSource is the data source implementation.
type ProductDataSource struct {
	client *sdk.SDK
}

// ProductDataSourceModel describes the data model.
type ProductDataSourceModel struct {
	ACL                   ACL                                 `tfsdk:"acl"`
	CreatedAt             types.String                        `tfsdk:"created_at"`
	Org                   types.String                        `tfsdk:"org"`
	Owners                []EntityOwner                       `tfsdk:"owners"`
	Schema                types.String                        `tfsdk:"schema"`
	Tags                  []types.String                      `tfsdk:"tags"`
	Title                 types.String                        `tfsdk:"title"`
	UpdatedAt             types.String                        `tfsdk:"updated_at"`
	AvailabilityFiles     []BaseRelation                      `tfsdk:"availability_files"`
	Code                  types.String                        `tfsdk:"code"`
	CrossSellableProducts *ProductCreateCrossSellableProducts `tfsdk:"cross_sellable_products"`
	Description           types.String                        `tfsdk:"description"`
	Feature               []Feature                           `tfsdk:"feature"`
	ID                    types.String                        `tfsdk:"id"`
	InternalName          types.String                        `tfsdk:"internal_name"`
	Name                  types.String                        `tfsdk:"name"`
	PriceOptions          *BaseRelation                       `tfsdk:"price_options"`
	ProductDownloads      *ProductCreateCrossSellableProducts `tfsdk:"product_downloads"`
	ProductImages         *ProductCreateCrossSellableProducts `tfsdk:"product_images"`
	Type                  types.String                        `tfsdk:"type"`
}

// Metadata returns the data source type name.
func (r *ProductDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_product"
}

// Schema defines the schema for the data source.
func (r *ProductDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Product DataSource",

		Attributes: map[string]schema.Attribute{
			"acl": schema.SingleNestedAttribute{
				Computed:   true,
				Attributes: map[string]schema.Attribute{},
			},
			"created_at": schema.StringAttribute{
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
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"availability_files": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dollar_relation": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"entity_id": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
					},
				},
				MarkdownDescription: `Stores references to the availability files that define where this product is available.` + "\n" +
					`These files are used when interacting with products via epilot Journeys, thought the AvailabilityCheck block.` + "\n" +
					``,
			},
			"code": schema.StringAttribute{
				Computed:    true,
				Description: `The product code`,
			},
			"cross_sellable_products": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"dollar_relation": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"dollar_relation": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"entity_id": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
							},
						},
					},
				},
				Description: `Stores references to products that can be cross sold with the current product.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `A description of the product. Multi-line supported.`,
			},
			"feature": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"tags": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `An arbitrary set of tags attached to a feature`,
						},
						"feature": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The product id`,
			},
			"internal_name": schema.StringAttribute{
				Computed:    true,
				Description: `Not visible to customers, only in internal tables`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The description for the product`,
			},
			"price_options": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"dollar_relation": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"entity_id": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
			"product_downloads": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"dollar_relation": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"dollar_relation": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"entity_id": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
							},
						},
					},
				},
				MarkdownDescription: `Stores references to a set of files downloadable from the product.` + "\n" +
					`e.g: tech specifications, quality control sheets, privacy policy agreements` + "\n" +
					``,
			},
			"product_images": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"dollar_relation": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"dollar_relation": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"entity_id": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
							},
						},
					},
				},
				Description: `Stores references to a set of file images of the product`,
			},
			"type": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: `must be one of ["product", "service"]; Default: "product"` + "\n" +
					`The type of Product:` + "\n" +
					`` + "\n" +
					`| type | description |` + "\n" +
					`|----| ----|` + "\n" +
					`| ` + "`" + `product` + "`" + ` | Represents a physical good |` + "\n" +
					`| ` + "`" + `service` + "`" + ` | Represents a service or virtual product |` + "\n" +
					``,
			},
		},
	}
}

func (r *ProductDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ProductDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ProductDataSourceModel
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

	productID := data.ID.ValueString()
	request := operations.GetProductRequest{
		ProductID: productID,
	}
	res, err := r.client.Product.GetProduct(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Product == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.Product)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
