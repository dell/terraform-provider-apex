/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package provider handles resource and data sources.
package provider

import (
	"context"
	"fmt"
	"net/http"

	client "github.com/dell/terraform-provider-aex/pkg/gen/apex/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &storageProductOptionsDataSource{}
	_ datasource.DataSourceWithConfigure = &storageProductOptionsDataSource{}
)

// NewStorageProductOptionsDataSource returns a new storage products data source instance.
func NewStorageProductOptionsDataSource() datasource.DataSource {
	return &storageProductOptionsDataSource{}
}

// storageProductOptionsDataSource is the data source implementation.
type storageProductOptionsDataSource struct {
	client *client.APIClient
}

// storageProductOptionsDataSourceModel maps the data source schema data.
type storageProductOptionsDataSourceModel struct {
	StorageProductOptions []offerTemplateModel `tfsdk:"storage_product_options"`
	ID                    types.String         `tfsdk:"id"`
}

// storageProductOptionsModel maps hosts schema data.
type offerTemplateModel struct {
	ID                    types.String              `tfsdk:"id"`
	SystemType            client.StorageProductEnum `tfsdk:"system_type"`
	StorageProductVersion types.String              `tfsdk:"storage_product_version"`
	CloudType             client.CloudProviderEnum  `tfsdk:"cloud_type"`
	SupportedTierInfos    []tierInfoModel           `tfsdk:"supported_tier_info"`
}

// tierInfoModel maps supportMap schema data
type tierInfoModel struct {
	ID             types.String      `tfsdk:"id"`
	Name           types.String      `tfsdk:"name"`
	TierType       client.TierEnum   `tfsdk:"tier_type"`
	Description    types.String      `tfsdk:"description"`
	StorageOptions []tierOptionModel `tfsdk:"storage_options"`
	CloudOptions   []tierOptionModel `tfsdk:"cloud_options"`
}

// tierOptionModel maps supportMap schema data
type tierOptionModel struct {
	ID    types.String             `tfsdk:"id"`
	Key   types.String             `tfsdk:"key"`
	Order types.Int64              `tfsdk:"order"`
	Value []client.TierOptionValue `tfsdk:"value"`
	// The Value field is declared as a slice even though it represents an object in the models.
	// This allows it to be read in the read function.
	// It may be necessary to revisit this implementation later to find an alternative to using a slice.
}

func (d *storageProductOptionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_storage_product_options"
}

// Schema defines the acceptable configuration and state attribute names and types.
func (d *storageProductOptionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) { //nolint funlen
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"storage_product_options": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "",
							Optional:            true,
						},
						"system_type": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"storage_product_version": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"cloud_type": schema.StringAttribute{
							MarkdownDescription: " ",
							Optional:            true,
						},
						"supported_tier_info": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Optional: true,
									},
									"name": schema.StringAttribute{
										Optional: true,
									},
									"tier_type": schema.StringAttribute{
										Optional: true,
									},
									"description": schema.StringAttribute{
										Optional: true,
									},
									"storage_options": schema.ListNestedAttribute{ //nolint: dupl
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Optional: true,
												},
												"key": schema.StringAttribute{
													Optional: true,
												},
												"order": schema.Int64Attribute{
													Optional: true,
												},
												"value": schema.ListNestedAttribute{ // tentative
													Optional: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{

															"boolean_option": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Optional: true,
																		},
																		"description": schema.StringAttribute{
																			Optional: true,
																		},
																		"default_value": schema.Int64Attribute{
																			Optional: true,
																		}, "is_configurable": schema.BoolAttribute{
																			Optional: true,
																		},
																	},
																},
															},
															"enum_option": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Optional: true,
																		},
																		"description": schema.StringAttribute{
																			Optional: true,
																		},
																		"available_options": schema.ListAttribute{
																			ElementType: types.StringType,
																			Optional:    true,
																		},
																		"default_value": &schema.StringAttribute{
																			Optional: true,
																		},
																		"is_configurable": schema.BoolAttribute{
																			Optional: true,
																		}, "type": schema.StringAttribute{
																			Optional: true,
																		},
																	},
																},
															},
															"free_text_option": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Optional: true,
																		},
																		"description": &schema.StringAttribute{
																			Optional: true,
																		},
																		"is_configurable": schema.StringAttribute{
																			Optional: true,
																		},
																		"type": schema.StringAttribute{
																			Optional: true,
																		},
																	},
																},
															},
															"multiselect_freetext_option": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Optional: true,
																		},
																		"description": schema.StringAttribute{
																			Optional: true,
																		},
																		"available_options": schema.ListAttribute{
																			ElementType: types.StringType,
																			Optional:    true,
																		},
																		"is_configurable": schema.BoolAttribute{
																			Optional: true,
																		}, "type": schema.BoolAttribute{
																			Optional: true,
																		},
																	},
																},
															}, "range_option": schema.ListNestedAttribute{
																Optional: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{

																		"name": schema.StringAttribute{
																			Optional: true,
																		},
																		"description": &schema.StringAttribute{
																			Optional: true,
																		},
																		"min_value": schema.Int64Attribute{
																			Optional: true,
																		},
																		"max_value": schema.Int64Attribute{
																			Optional: true,
																		},
																		"linear_interval": schema.Int64Attribute{
																			Optional: true,
																		},
																		"unity": schema.StringAttribute{
																			Optional: true,
																		},
																		"default_value": &schema.Int64Attribute{
																			Optional: true,
																		},
																		"is_configurable": schema.BoolAttribute{
																			Optional: true,
																		},
																		"type": schema.StringAttribute{
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"cloud_options": schema.ListNestedAttribute{ //nolint: dupl
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Optional: true,
												},
												"key": schema.StringAttribute{
													Optional: true,
												},
												"order": schema.Int64Attribute{
													Optional: true,
												},
												"value": schema.SingleNestedAttribute{
													Optional: true,
													Attributes: map[string]schema.Attribute{

														"boolean_option": schema.SingleNestedAttribute{
															Optional: true,
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Optional: true,
																},
																"description": schema.StringAttribute{
																	Optional: true,
																},
																"default_value": schema.Int64Attribute{
																	Optional: true,
																}, "is_configurable": schema.BoolAttribute{
																	Optional: true,
																},
															},
														},
														"enum_option": schema.SingleNestedAttribute{
															Optional: true,
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Optional: true,
																},
																"description": schema.StringAttribute{
																	Optional: true,
																},
																"available_options": schema.ListAttribute{
																	ElementType: types.StringType,
																	Optional:    true,
																},
																"default_value": &schema.StringAttribute{
																	Optional: true,
																},
																"is_configurable": schema.BoolAttribute{
																	Optional: true,
																}, "type": schema.StringAttribute{
																	Optional: true,
																},
															},
														},
														"free_text_option": schema.SingleNestedAttribute{
															Optional: true,
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Optional: true,
																},
																"description": &schema.StringAttribute{
																	Optional: true,
																},
																"is_configurable": schema.StringAttribute{
																	Optional: true,
																},
																"type": schema.StringAttribute{
																	Optional: true,
																},
															},
														}, "multiselect_freetext_option": schema.SingleNestedAttribute{
															Optional: true,
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Optional: true,
																},
																"description": schema.StringAttribute{
																	Optional: true,
																},
																"available_options": schema.ListAttribute{
																	ElementType: types.StringType,
																	Optional:    true,
																},
																"is_configurable": schema.BoolAttribute{
																	Optional: true,
																}, "type": schema.BoolAttribute{
																	Optional: true,
																},
															},
														}, "range_option": schema.SingleNestedAttribute{
															Optional: true,
															Attributes: map[string]schema.Attribute{

																"name": schema.StringAttribute{
																	Optional: true,
																},
																"description": &schema.StringAttribute{
																	Optional: true,
																},
																"min_value": schema.Int64Attribute{
																	Optional: true,
																},
																"max_value": schema.Int64Attribute{
																	Optional: true,
																},
																"linear_interval": schema.Int64Attribute{
																	Optional: true,
																},
																"unity": schema.StringAttribute{
																	Optional: true,
																},
																"default_value": &schema.Int64Attribute{
																	Optional: true,
																}, "is_configurable": schema.BoolAttribute{
																	Optional: true,
																},
																"type": schema.StringAttribute{
																	Optional: true,
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Read method is used to refresh the Terraform state based on the schema data.
func (d *storageProductOptionsDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state storageProductOptionsDataSourceModel

	storageProductOptions, status, err := d.client.StorageProductOptionsAPI.StorageProductOptionsCollection(context.Background()).Limit(500).Execute()
	if (err != nil) && (status.StatusCode != http.StatusOK) {
		resp.Diagnostics.AddError(
			"Unable Read to Storage Product Volume",
			err.Error(),
		)
		return
	}
	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}

	// Map response body to model
	for _, storageProductOption := range storageProductOptions.Results {
		storageProductOptionState := offerTemplateModel{
			ID:                    types.StringValue(*storageProductOption.Id),
			SystemType:            storageProductOption.SystemType,
			StorageProductVersion: types.StringValue(storageProductOption.StorageProductVersion),
			CloudType:             storageProductOption.CloudType,
		}

		for _, supportedtierinfo := range storageProductOption.SupportedTierInfo {
			storageProductOptionState.SupportedTierInfos = append(storageProductOptionState.SupportedTierInfos, tierInfoModel{
				ID:          types.StringValue(supportedtierinfo.Id),
				Name:        types.StringValue(supportedtierinfo.Name),
				TierType:    supportedtierinfo.TierType,
				Description: types.StringValue(supportedtierinfo.Description),
			})

			for _, storageoption := range supportedtierinfo.StorageOptions {
				storageProductOptionState.SupportedTierInfos[len(storageProductOptionState.SupportedTierInfos)-1].StorageOptions = append(storageProductOptionState.SupportedTierInfos[len(storageProductOptionState.SupportedTierInfos)-1].StorageOptions, tierOptionModel{
					ID:    types.StringValue(storageoption.Id),
					Key:   types.StringValue(storageoption.Key),
					Order: types.Int64Value(int64(storageoption.Order)),
					Value: []client.TierOptionValue{},
				})
			}

			for _, cloudoption := range supportedtierinfo.CloudOptions {
				storageProductOptionState.SupportedTierInfos[len(storageProductOptionState.SupportedTierInfos)-1].CloudOptions = append(storageProductOptionState.SupportedTierInfos[len(storageProductOptionState.SupportedTierInfos)-1].CloudOptions, tierOptionModel{
					ID:    types.StringValue(cloudoption.Id),
					Key:   types.StringValue(cloudoption.Key),
					Order: types.Int64Value(int64(cloudoption.Order)),
					Value: []client.TierOptionValue{},
				})
			}
		}

		state.StorageProductOptions = append(state.StorageProductOptions, storageProductOptionState)
	}

	state.ID = types.StringValue("placeholder")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *storageProductOptionsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *openapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

//
