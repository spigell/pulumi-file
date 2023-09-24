package main

import (
	"fmt"
	"os"
	"strings"

	p "github.com/pulumi/pulumi-go-provider"

	file "github.com/spigell/pulumi-file/provider/pkg/provider"
	"github.com/spigell/pulumi-file/provider/pkg/version"
)

// A provider is a program that listens for requests from the Pulumi engine
// to interact with cloud providers using a CRUD-based model.
func main() {
	version := strings.TrimPrefix(version.Version, "v")

	// This method defines the provider implemented in this repository.
	provider := file.NewProvider()

	// This method starts serving requests using the Command provider.
	err := p.RunProvider("file", version, provider)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}