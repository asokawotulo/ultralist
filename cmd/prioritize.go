package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/asokawotulo/ultralist/ultralist"
)

var (
	revertPrioritizedTodo bool
	prioritizeCmdDesc     = "Prioritize and unprioritize todos"
	prioritizeCmdExample  = `  ultralist prioritize 33
  Prioritizes todo with id 33.

  ultralist prioritize 33 --revert
  Unprioritizes todo with id 33.`
	prioritizeCmdLongDesc = `Prioritize and unprioritize todos.

Todos with a priority flag will be highlighted on top of a list.`
)

var prioritizeCmd = &cobra.Command{
	Use:     "prioritize [id]",
	Aliases: []string{"p"},
	Example: prioritizeCmdExample,
	Long:    prioritizeCmdLongDesc,
	Short:   prioritizeCmdDesc,
	Run: func(cmd *cobra.Command, args []string) {
		if revertArchivedTodo {
			ultralist.NewApp().UnprioritizeTodo(strings.Join(args, " "))
		} else {
			ultralist.NewApp().PrioritizeTodo(strings.Join(args, " "))
		}
	},
}

var (
	prioritizeListCmdDesc     = "Lists all prioritized todos"
	prioritizeListCmdLongDesc = prioritizeListCmdDesc + "."
)

var prioritizeListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Long:    prioritizeListCmdLongDesc,
	Short:   prioritizeListCmdDesc,
	Run: func(cmd *cobra.Command, args []string) {
		ultralist.NewApp().ListTodos("prioritized")
	},
}

func init() {
	rootCmd.AddCommand(prioritizeCmd)
	prioritizeCmd.Flags().BoolVarP(&revertArchivedTodo, "revert", "", false, "Unprioritizes an prioritized todo")
	prioritizeCmd.AddCommand(prioritizeListCmd)
}
