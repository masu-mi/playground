package main

import (
	"context"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/packer/plugin"
)

// Assume this implements packer.Builder
type Builder struct{}

func (b *Builder) Prepare(_ ...interface{}) ([]string, []string, error) {
	return []string{}, []string{}, nil
}

// Run is where the actual build should take place. It takes a Build and a Ui.
func (b *Builder) Run(ctx context.Context, _ packer.Ui, _ packer.Hook) (packer.Artifact, error) {
	return &packer.MockArtifact{}, nil
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec {
	return map[string]hcldec.Spec{}
}

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(Builder))
	server.Serve()
}
