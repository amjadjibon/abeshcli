package template

import (
	"context"
	"text/template"

	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"
)

type Service struct {
}

func (s Service) Name() string {
	return Name
}

func (s Service) Version() string {
	return Version
}

func (s Service) Category() string {
	return Category
}

func (s Service) ContractId() string {
	return ContractId
}

func (s Service) New() iface.ICapability {
	return &Service{}
}

func (s Service) Serve(_ context.Context, event *model.Event) (*model.Event, error) {
	return model.GenerateOutputEvent(
		event.Metadata,
		s.ContractId(),
		"OK",
		200,
		"application/json",
		[]byte("{}"),
	), nil

}

func init() {
	registry.GlobalRegistry().AddCapability(&Service{})
}

var ServiceTemplate = template.Must(template.New("service").Parse(`
package {{.PackageName}}

import (
	"context"

	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"
)

type {{.StructName}} struct {
}

func (s {{.StructName}}) Name() string {
	return Name
}

func (s {{.StructName}}) Version() string {
	return Version
}

func (s {{.StructName}}) Category() string {
	return Category
}

func (s {{.StructName}}) ContractId() string {
	return ContractId
}

func (s {{.StructName}}) New() iface.ICapability {
	return &Service{}
}

func (s {{.StructName}}) Serve(_ context.Context, event *model.Event) (*model.Event, error) {
	return model.GenerateOutputEvent(
		event.Metadata,
		s.ContractId(),
		"OK",
		200,
		"application/json",
		[]byte("{}"),
	), nil

}

func init() {
	registry.GlobalRegistry().AddCapability(&{{.StructName}}{})
}
`))
