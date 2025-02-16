package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/asokawotulo/ultralist/ultralist"
)

var (
	versionCmdDesc     = "Displays the version of ultralist"
	versionCmdLongDesc = versionCmdDesc + "."
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Long:  versionCmdLongDesc,
	Short: versionCmdDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Ultralist v%s.", ultralist.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
