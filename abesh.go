package main

import (
	"log"

	"github.com/amjadjibon/abeshcli/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
