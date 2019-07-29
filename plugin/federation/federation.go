package federation

import (
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin"
	"github.com/vektah/gqlparser/ast"
)

type Plugin struct{}

var _ plugin.Plugin = &Plugin{}
var _ plugin.ConfigMutator = &Plugin{}
var _ config.SchemaMutator = &Plugin{}

func New() plugin.Plugin {
	return &Plugin{}
}

func (*Plugin) Name() string {
	return "federation"
}

func (p *Plugin) MutateConfig(cfg *config.Config) error {
	cfg.Init()
	cfg.AddMutator(p)
	cfg.AddToSchema(&ast.Source{Name: "federation", Input: federationSchema})
	cfg.Models.Add("_Service", "github.com/99designs/gqlgen/plugin/federation.Service")
	cfg.Models.Add("_Entity", "github.com/99designs/gqlgen/plugin/federation.Entity")
	return nil
}

func (p *Plugin) MutateSchema(s *ast.Schema) error {
	logger.Print("Mutating schema")
	entunion := s.Types["_Entity"]
	for _, typ := range s.Types {
		switch typ.Kind {
		case ast.Object:
			keys := typ.Directives.ForName("key")
			if keys != nil {
				logger.Printf("Adding %s to _Entity", typ.Name)
				entunion.Types = append(entunion.Types, typ.Name)
				s.AddPossibleType("_Entity", typ)
				s.AddImplements(typ.Name, entunion)
			}
		}
	}
	return nil
}

const federationSchema = `# federation schema add-on
	scalar _FieldSet

	type _Service { sdl: String }
	extend type Query { _service: _Service }

union _Entity

	directive @key(fields: _FieldSet!) on OBJECT | INTERFACE
`
