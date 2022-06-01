package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/amjadjibon/abeshcli/cmd/internal/build"
	"github.com/amjadjibon/abeshcli/cmd/internal/project"
	"github.com/amjadjibon/abeshcli/cmd/internal/run"
	"github.com/amjadjibon/abeshcli/constant"
)

var RootCmd = &cobra.Command{
	Use:     "abesh",
	Short:   "abesh: An elegant toolkit for Go microservices",
	Long:    `abesh: An elegant toolkit for Go microservices`,
	Version: constant.Version,
}

func init() {
	RootCmd.AddCommand(project.CmdNew)
	RootCmd.AddCommand(run.CmdRun)
	RootCmd.AddCommand(build.CmdBuild)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
