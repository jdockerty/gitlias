/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/jdockerty/gitlias/gitlias"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more aliases from your configuration.",
	Long: `Remove one or more aliases from your configuration.

Example:

    gitlias rm singleAlias
    gitlias rm aliasOne aliasTwo
    gitlias rm --config /path/to/gitlias.toml testAccount
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("You must provide an alias to remove.")
			return
		} else if len(args) == 1 { // remove single alias
			alias := args[0]

			removed, err := gitlias.Remove(configPath, alias)
			if err != nil {
				fmt.Printf("Cannot remove provided '%s' alias: %s\n", alias, err)
				return
			}
			err = removed.WriteConfig(configPath)
			if err != nil {
				fmt.Printf("Error writing config after removing '%s' alias: %s\n", alias, err)
			}

			fmt.Printf("Removed %s\n", alias)

		} else { // remove multiple aliases

			for _, alias := range args {
				removed, err := gitlias.Remove(configPath, alias)
				if err != nil {
					fmt.Printf("Cannot remove provided '%s' alias: %s\n", alias, err)
					return
				}

				err = removed.WriteConfig(configPath)
				if err != nil {
					fmt.Printf("Error writing config after removing '%s' alias: %s\n", alias, err)
				}
				fmt.Printf("Removed %s\n", alias)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
