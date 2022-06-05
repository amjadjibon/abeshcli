package main

import (
	_ "embed"

	"github.com/mkawserm/abesh/cmd"
	_ "github.com/mkawserm/httpserver2/capability/httpserver2"

	_ "github.com/amjadjibon/example/capability/echo"
	_ "github.com/amjadjibon/example/capability/health"
)

//go:embed manifest.yaml
var manifestBytes []byte

func main() {
	cmd.ManifestBytes = manifestBytes
	cmd.Execute()
}
