package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jdockerty/gitlias/gitlias"
)

const configName = "gitlias.toml"

var configPath string

func init() {
	defaultConf := func() string {
		h, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Unable to get user home directory: %s\n", err)
		}
		return fmt.Sprintf("%s/%s", h, configName)
	}()

	flag.StringVar(&configPath, "config", defaultConf, "Configuration TOML file path")
	flag.Parse()
}

func main() {
	gitlias.Run(configPath)
}
