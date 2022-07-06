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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

func switchAlias(confPath string, alias string, gitConf *config.Config, userConf *gitlias.Gitlias) error {

	if _, ok := userConf.Alias[alias]; !ok {
		fmt.Println("Invalid key provided, valid keys are:")
		for key, _ := range userConf.Alias {
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

		// with debug: log.Printf("Writing to path: %s\n", p)
		if _, err = os.Stat(p); errors.Is(err, os.ErrNotExist) {
			// with debug: fmt.Printf("Skip writing to %s as it does not exist\n", p)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
