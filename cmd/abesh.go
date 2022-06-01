package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/amjadjibon/abeshcli/cmd/internal/build"
	"github.com/amjadjibon/abeshcli/cmd/internal/project"
	"github.com/amjadjibon/abeshcli/cmd/internal/run"
	"github.com/amjadjibon/abeshcli/constant"
)

var rootCmd = &cobra.Command{
	Use:     "abesh",
	Short:   "abesh: An elegant toolkit for Go microservices",
	Long:    `abesh: An elegant toolkit for Go microservices`,
	Version: constant.Version,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(run.CmdRun)
	rootCmd.AddCommand(build.CmdBuild)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
