package main

import (
	_ "embed"

	"github.com/mkawserm/abesh/cmd"

	_ "github.com/amjadjibon/hello/capability/health"

	_ "github.com/mkawserm/httpserver2/capability/httpserver2"
)

//go:embed manifest.yaml
var manifestBytes []byte

func main() {
	cmd.ManifestBytes = manifestBytes
	cmd.Execute()
}
