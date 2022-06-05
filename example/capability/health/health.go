package health

import (
	"context"

	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/model"
	"github.com/mkawserm/abesh/registry"

	"github.com/amjadjibon/example/constant"
)

type Health struct {
	mCM model.ConfigMap
}

func (h Health) SetConfigMap(cm model.ConfigMap) error {
	h.mCM = cm
	return nil
}

func (h Health) GetConfigMap() model.ConfigMap {
	return h.mCM
}

func (h Health) Name() string {
	return Name
}

func (h Health) Version() string {
	return constant.Version
}

func (h Health) Category() string {
	return Category
}

func (h Health) ContractId() string {
	return ContractId
}

func (h Health) New() iface.ICapability {
	return &Health{}
}

func (h Health) Serve(ctx context.Context, event *model.Event) (*model.Event, error) {
	return model.GenerateOutputEvent(
		event.Metadata,
		h.ContractId(),
		"OK",
		200,
		"application/json",
		[]byte("{}"),
	), nil
}

func init() {
	registry.GlobalRegistry().AddCapability(&Health{})
}
