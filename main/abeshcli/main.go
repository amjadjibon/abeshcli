package main

import (
    _ "embed"
    "github.com/mkawserm/abesh/cmd"

)

//go:embedded manifest.yaml
var manifestBytes []byte

func main() {
    cmd.Excecute()
}
