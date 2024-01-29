package tests

import (
	provider "eos2git.cec.lab.emc.com/CIRRUS/cirrus-terraform-providers/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the Apex Navigator client is properly configured.
	// It is also possible to use the environment variables instead,
	// such as updating the Makefile and running the testing through that tool.

	sourceConfig = `terraform {
		required_providers {
		  apex = {
			version = "~> 0.0.0"
			source = "dell/apex"
		  }
		}
	  }`

	providerConfig = `
	provider "apex" {
		host = "http://localhost:8080"
	  }
`
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"apex": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)
