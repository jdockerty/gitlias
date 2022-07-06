/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jdockerty/gitlias/gitlias"
	"github.com/spf13/cobra"
)

const configName = "gitlias.toml"

var configPath string
var listAliases bool
var generateTemplate bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitlias",
	Short: "Swap between your configured git aliases",
	Long: `gitlias
Swap between your configured git aliases, ensuring that you commit as the correct user.
`,
	Run: func(cmd *cobra.Command, args []string) {
		listAlias, err := cmd.Flags().GetBool("list")
		if err != nil {
			fmt.Println(err)
			return
		}

		if listAlias {
			aliases, current := gitlias.List(configPath)
			s := strings.Join(aliases, "\n")
			fmt.Printf("%s\n\ncurrent: %s\n", s, current)
			return
		}
	},
}

// Execute adds child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	defaultConf := func() string {
		h, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Unable to get user home directory: %s\n", err)
		}
		return fmt.Sprintf("%s/%s", h, configName)
	}()

	rootCmd.PersistentFlags().StringVar(&configPath, "config", defaultConf, "Configuration file path")
	rootCmd.Flags().BoolVar(&listAliases, "list", false, "List current aliases")
}
