/*
Coryright © 2022 Jack Dockerty jdockerty19@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5/config"
	"github.com/jdockerty/gitlias/gitlias"
	"github.com/spf13/cobra"
)

const scope = config.GlobalScope

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch between different aliases",
	Long: `Switch between your configured aliases.

Example:

    gitlias switch personal
    gitlias --config /path/to/gitlias.toml switch work
`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("You must provide an alias to switch to.")
			return
		}

		alias := args[0]

		gitConfig, err := config.LoadConfig(scope)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		userConfig, err := gitlias.Get(configPath)
		if err != nil {
			fmt.Printf("Unable to find configuration file: %s\n", err)
			os.Exit(1)
		}

		err = switchAlias(configPath, alias, gitConfig, userConfig)
		if err != nil {
			fmt.Printf("Unable to switch to alias: %s\n", err)
			os.Exit(1)
		}
	},
}

// switchAlias switches aliases by writing to the global git configuration.
func switchAlias(confPath string, alias string, gitConf *config.Config, userConf *gitlias.Gitlias) error {

	if _, ok := userConf.Alias[alias]; !ok {
		fmt.Println("Invalid key provided, valid keys are:")
		for key := range userConf.Alias {
			fmt.Printf("\t%s\n", key)
		}
		return errors.New("the alias you want to switch to must exist")
	}
	a := userConf.Alias[alias]

	gitConf.User.Name = a.User
	gitConf.User.Email = a.Email

	data, err := gitConf.Marshal()
	if err != nil {
		return err
	}

	paths, err := config.Paths(scope)
	if err != nil {
		return err
	}

	for _, p := range paths {

		if _, err = os.Stat(p); errors.Is(err, os.ErrNotExist) {
			continue
		}
		err = os.WriteFile(p, data, os.ModeAppend)
		if err != nil {
			return err
		}
	}
	userConf.Current = alias
	err = userConf.WriteConfig(confPath)
	if err != nil {
		return err
	}

	fmt.Printf("Switched to alias: %s\n", alias)
	return nil
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
