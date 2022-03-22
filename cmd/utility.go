package cmd

import "github.com/mkawserm/abesh/iface"

func GetCapability(v map[string]iface.ICapability, contractId string) iface.ICapability {
	return v[contractId]
}
