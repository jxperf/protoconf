package terraformimporter

import (
	"log"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestBuilder(t *testing.T) {
	log.Println("starting")
	provider := "google"
	version := "3.13.0"
	// meta, err := findPlugin("provider", "random", "2.2.1")
	meta, err := findPlugin("provider", provider, version)
	assert.NoError(t, err)
	config := newGRPCClientConfig(meta)
	client, err := NewGRPCClient(config)
	assert.NoError(t, err)
	defer client.Close()
	g := NewGenerator("", "")
	p, err := NewProviderImporter(*meta, client, g.Importer)
	assert.NoError(t, err)
	Print(p.importer.MasterFile)
	// Print(p.Resources)
	// Print(p.Datasources)
	// Print(p.Provider)
	// t.Fail()
}
