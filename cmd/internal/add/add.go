package add

import (
	"github.com/spf13/cobra"
)

// CmdAdd run project command.
var CmdAdd = &cobra.Command{
	Use:   "add",
	Short: "add capability",
	Long:  "add capability",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	CmdAdd.AddCommand(CmdAddService)
}
