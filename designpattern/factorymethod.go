package main

import (
	"os"
	"strings"
)

// Provider represents a configurable provider.
type Provider interface {
	Name() string
	Configure(config map[string]interface{}) error
}

// SecretReader represents a contract for a store capable of resolving secrets.
type SecretReader interface {
	Provider
	GetSecret(secretName string) (string, bool)
}

type SecretStore interface {
	SecretReader
	GetCache() ([]string, bool)
}

//---------------------------------------------------------------------------------

// EnvironmentProvider represents a security provider which uses environment variables to store secrets.
type EnvironmentProvider struct {
	lookup func(string) (string, bool)
}

// NewEnvironmentProvider creates a new environment security provider.
func NewEnvironmentProvider() *EnvironmentProvider {
	return &EnvironmentProvider{
		lookup: os.LookupEnv,
	}
}

// Name returns the name of the security provider.
func (p *EnvironmentProvider) Name() string {
	return "environment"
}

// Configure configures the security provider.
func (p *EnvironmentProvider) Configure(config map[string]interface{}) (err error) {
	return
}

// GetSecret retrieves a secret from the provider
func (p *EnvironmentProvider) GetSecret(secretName string) (string, bool) {
	name := strings.ToUpper(strings.Replace(secretName, "/", "_", -1))
	return p.lookup(name)
}
