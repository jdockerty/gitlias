/*
Coryright © 2022 Jack Dockerty jdockerty19@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/jdockerty/gitlias/gitlias"
	"github.com/spf13/cobra"
)

var (
	addUser  string
	addEmail string
	addAlias string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new alias",
	Long: `Add a new alias your gitlias configuration file.

All flags must be specified with this command, for example:

        gitlias add --alias personal --user "John Smith" --email "john@example.com"
`,
	Run: func(cmd *cobra.Command, args []string) {
		aliases, err := gitlias.Add(configPath, addAlias, addUser, addEmail)
		if err != nil {
			fmt.Printf("Unable to add alias: %s\n", err)
			return
		}

		err = aliases.WriteConfig(configPath)
		if err != nil {
			fmt.Printf("Unable to write additional alias to configuraion: %s\n", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&addUser, "user", "", "The username to make commits as, e.g. 'John Smith'")
	addCmd.Flags().StringVar(&addEmail, "email", "", "The email to make commits as, e.g. 'john.smith@example.com'")
	addCmd.Flags().StringVar(&addAlias, "alias", "", "A meaningful alias to give to this commit author")

	// Flags are required, this makes it far easier to parse the given values.
	addCmd.MarkFlagRequired("user")
	addCmd.MarkFlagRequired("email")
	addCmd.MarkFlagRequired("alias")
}
