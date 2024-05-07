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

package provider

import (
	"context"
	"net/url"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &myProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &myProvider{
			version: version,
		}
	}
}

// myProviderModel maps provider schema data to a Go type.
type myProviderModel struct {
	Host        types.String `tfsdk:"host"`
	Token       types.String `tfsdk:"token"`
	JMSEndpoint types.String `tfsdk:"jms_endpoint"`
	Insecure    types.Bool   `tfsdk:"insecure"`
}

// myProvider is the provider implementation.
type myProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *myProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "apex"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *myProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
			},
			"token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"jms_endpoint": schema.StringAttribute{
				Optional: true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Boolean variable to specify whether to validate SSL certificate or not.",
				Description:         "Boolean variable to specify whether to validate SSL certificate or not.",
				Optional:            true,
			},
		},
	}
}

// Configure is the method responsible for configuring the provider.
func (p *myProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) { //nolint:funlen
	// Retrieve provider data from configuration
	var config myProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	token := config.Token.ValueString()
	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Apex Navigator API Host",
			"The provider cannot create the Apex Navigator API client as there is an unknown configuration value for the Apex Navigator API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CIRRUS_HOST environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.
	jmsEndpoint := os.Getenv("JMS_ENDPOINT")
	host := os.Getenv("APEX_NAVIGATOR_HOST")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.JMSEndpoint.IsNull() {
		jmsEndpoint = config.JMSEndpoint.ValueString()
	} else if jmsEndpoint == "" {
		jmsEndpoint = config.Host.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.
	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Apex Navigator API Host",
			"The provider cannot create the Apex Navigator API client as there is a missing or empty value for the Apex Navigator API host. "+
				"Set the host value in the configuration or use the APEX_NAVIGATOR_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	hostURL, err := url.Parse(host)
	if err != nil {
		resp.Diagnostics.AddError(
			"Could not parse host URL: "+host,
			"Could not parse host URL: "+host,
		)
	}
	jmsURL, err := url.Parse(jmsEndpoint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Could not parse job host URL: "+jmsEndpoint,
			"Could not parse job host URL: "+jmsEndpoint,
		)
	}
	if hostURL.Scheme == "" {
		resp.Diagnostics.AddError(
			"Missing scheme in host URL",
			"Missing scheme in host URL: "+host,
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create the API and JMS client
	apiClients, err := NewApexJmsClient(ctx, *hostURL, *jmsURL, token, config.Insecure.ValueBool())
	if err != nil {
		resp.Diagnostics.AddError(
			"Could not create client",
			"Could not create client: "+err.Error(),
		)
		return
	}

	// Make the cirrus client available during DataSource and Resource
	// Make the Apex Navigator client available during DataSource and Resource
	resp.DataSourceData = apiClients.APIClient
	resp.ResourceData = Clients{APIClient: apiClients.APIClient, JMSClient: apiClients.JMSClient}
}

// DataSources defines the data sources implemented in the provider.
func (p *myProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewHostsDataSource,
		NewVolumesDataSource,
		NewStorageProductsDataSource,
		NewPoolsDataSource,
		NewClonesDataSource,
		NewStoragesDataSource,
		NewMobilityGroupsDataSource,
		NewMobilityTargetsDataSource,
		NewAwsAccountsDataSource,
		NewAwsPermissionsDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *myProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewClonesResource,
		NewBlockStorageResource,
		NewFileStorageResource,
		NewMobilityGroupsResource,
		NewMobilityTargetsResource,
		NewMobilityGroupsCopyResource,
		NewClonesMapResource,
		NewClonesUnmapResource,
		NewClonesRefreshResource,
		NewAwsAccountResource,
		NewTrustPolicyGenerateResource,
	}
}
