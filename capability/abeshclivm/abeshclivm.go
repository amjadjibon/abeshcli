package abeshclivm

import (
	"os"
	"strings"

	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/registry"

	"github.com/amjadjibon/abeshcli/constant"
	"github.com/amjadjibon/abeshcli/template"
)

type AbeshCLIVM struct {
}

func (a AbeshCLIVM) Name() string {
	return constant.Name
}

func (a AbeshCLIVM) Version() string {
	return constant.Version
}

func (a AbeshCLIVM) Category() string {
	return constant.Category
}

func (a AbeshCLIVM) ContractId() string {
	return constant.ContractId
}

func (a AbeshCLIVM) New() iface.ICapability {
	return &AbeshCLIVM{}
}

func (a AbeshCLIVM) GenerateService() error {
	fileName := "file.go"
	if !strings.HasSuffix(fileName, ".go") {
		fileName += ".go"
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	template.ServiceTemplate.ExecuteTemplate(f, "service", struct {
		PackageName string
	}{
		PackageName: constant.Name,
	})

	return nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&AbeshCLIVM{})
}
