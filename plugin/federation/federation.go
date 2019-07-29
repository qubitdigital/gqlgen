package federation

import (
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
	"github.com/vektah/gqlparser/ast"
)

type Plugin struct{}

var _ plugin.Plugin = &Plugin{}
var _ plugin.ConfigMutator = &Plugin{}

func New() plugin.Plugin {
	return &Plugin{}
}

func (*Plugin) Name() string {
	return "federation"
}

func (p *Plugin) MutateConfig(cfg *config.Config) error {
	cfg.Init()
	cfg.AddToSchema(federationSchema)
	cfg.Models.Add("_Service", "github.com/99designs/gqlgen/plugin/federation.Service")
	return nil
}

var federationSchema = &ast.Source{
	Name: "federation",
	Input: `
	scalar _FieldSet

	type _Service { sdl: String }
	extend type Query { _service: _Service }

	directive @key(fields: _FieldSet!) on OBJECT | INTERFACE
	`,
	BuiltIn: true,
}
