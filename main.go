package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jdockerty/gitlias/gitlias"
)

const helpText = `gitlias

Swap between your configured git aliases.

Examples:

	gitlias <alias>
	gitlias -config /tmp/gitlias.toml <alias>
	gitlias -list
`
const configName = "gitlias.toml"

var configPath string
var listAliases bool

func init() {
	defaultConf := func() string {
		h, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Unable to get user home directory: %s\n", err)
		}
		return fmt.Sprintf("%s/%s", h, configName)
	}()

	flag.StringVar(&configPath, "config", defaultConf, "Configuration TOML file path")
	flag.BoolVar(&listAliases, "list", false, "List current alias names in your configuration file")
	flag.Parse()
}

func main() {

	if listAliases {
		gitlias.List(configPath)
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Printf("%s\n", helpText)
		flag.PrintDefaults()
		return
	}

	gitlias.Run(configPath)
}
