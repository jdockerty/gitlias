/*
Copyright © 2022 Jack Dockerty jdockerty19@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jdockerty/gitlias/gitlias"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise gitlias",
	Long: `Initialise gitlias

Generate a template configuration file, writing to the config file path, ideal for your first time running the application.
`,
	Run: func(cmd *cobra.Command, args []string) {
		tmpl := gitlias.Init()

		err := os.WriteFile(configPath, []byte(tmpl), 0644)
		if err != nil {
			fmt.Printf("Unable to write init file: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
