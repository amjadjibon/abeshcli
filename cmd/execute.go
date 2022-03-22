package cmd

import (
	abeshCMD "github.com/mkawserm/abesh/cmd"
	"github.com/mkawserm/abesh/platform"
	"github.com/spf13/cobra"
)

func Execute() {
	abeshCMD.DefaultCMDHandler = func(c *cobra.Command) {
		abeshCMD.DefaultProject = &Project{}
		abeshCMD.DefaultPlatform = &platform.One{}

		c.Use = abeshCMD.DefaultProject.Name()
		c.Short = abeshCMD.DefaultProject.ShortDescription()
		c.Long = abeshCMD.DefaultProject.LongDescription()
	}

	abeshCMD.Execute()
}
