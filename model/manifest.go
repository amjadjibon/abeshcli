package model

type ConfigMap map[string]string

type CapabilityManifest struct {
	ContractId    string    `yaml:"contract_id,omitempty" json:"contract_id"`
	NewContractId string    `yaml:"new_contract_id,omitempty"`
	Values        ConfigMap `yaml:"values,omitempty" json:"values"`
}

type TriggerManifest struct {
	Trigger              string    `yaml:"trigger,omitempty" json:"trigger"`
	TriggerValues        ConfigMap `yaml:"trigger_values,omitempty" json:"trigger_values"`
	Service              string    `yaml:"service,omitempty" json:"service"`
	Authorizer           string    `yaml:"authorizer,omitempty" json:"authorizer"`
	AuthorizerExpression string    `yaml:"authorizer_expression,omitempty" json:"authorizer_expression"`
}

type RPCManifest struct {
	RPC                  string `yaml:"rpc,omitempty" json:"rpc"`
	Method               string `yaml:"method,omitempty" json:"method"`
	Authorizer           string `yaml:"authorizer,omitempty" json:"authorizer"`
	AuthorizerExpression string `yaml:"authorizer_expression,omitempty" json:"authorizer_expression"`
}

type ConsumerManifest struct {
	Source string `yaml:"source,omitempty" json:"source"`
	Sink   string `yaml:"sink,omitempty" json:"sink"`
}

type Manifest struct {
	Version      string                `yaml:"version,omitempty" json:"version"` // 1
	Capabilities []*CapabilityManifest `yaml:"capabilities,omitempty" json:"capabilities"`
	Triggers     []*TriggerManifest    `yaml:"triggers,omitempty" json:"triggers"`
	RPCS         []*RPCManifest        `yaml:"rpcs,omitempty" json:"rpcs"`
	Consumers    []*ConsumerManifest   `yaml:"consumers,omitempty" json:"consumers"`
	Start        []string              `yaml:"start,omitempty" json:"start"`
}
