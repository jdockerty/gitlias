/*
Copyright © 2022 Jack Dockerty jdockerty19@gmail.com
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/jdockerty/gitlias/gitlias"
	"github.com/spf13/cobra"
)

var current bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured aliases",
	Long: `List all configured aliases within your configuration file.

This can be modified using the --current flag to display the current active alias alone.

For example:

    gitlias list
    gitlias list --current
`,
	Run: func(cmd *cobra.Command, args []string) {
		aliases, currentAlias := gitlias.List(configPath)
		if current {
			fmt.Println(currentAlias)
			return
		}
		s := strings.Join(aliases, "\n")
		fmt.Println(s)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&current, "current", false, "show only the active alias")
}
