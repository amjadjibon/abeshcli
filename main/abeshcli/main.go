package main

import (
	_ "embed"

	"github.com/mkawserm/abesh/cmd"

	abeshcliCMD "github.com/amjadjibon/abeshcli/cmd"
)

//go:embedded manifest.yaml
var manifestBytes []byte

func main() {
	cmd.ManifestBytes = manifestBytes
	abeshcliCMD.Execute()
}
