package cmds

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "app",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}
