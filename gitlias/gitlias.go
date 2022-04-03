package gitlias

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5/config"
)

const scope = config.GlobalScope
const configName = "gitlias.toml"

func List(configPath string) []string {
	userConfig, err := Get(configPath)
	if err != nil {
		fmt.Printf("Unable to find configuration file: %s\n", err)
		os.Exit(1)
	}
	var aliases []string
	for key, _ := range userConfig.Alias {
		aliases = append(aliases, key)
	}

	return aliases
}

func Run(configPath string) {
	gitConfig, err := config.LoadConfig(scope)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	userConfig, err := Get(configPath)
	if err != nil {
		fmt.Printf("Unable to find configuration file: %s\n", err)
		os.Exit(1)
	}

	alias := func() string {
		a := flag.Args()
		if len(a) == 0 {
			fmt.Println("You must provide an alias to switch to")
			os.Exit(1)
		}
		return a[0]
	}()

	err = switchAlias(alias, gitConfig, userConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func switchAlias(alias string, gitConf *config.Config, userConf *Gitlias) error {

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

		log.Printf("Writing to path: %s\n", p)
		if _, err = os.Stat(p); errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Skip writing to %s as it does not exist\n", p)
			continue
		}
		err = os.WriteFile(p, data, os.ModeAppend)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Switched to alias: %s\n", alias)
	return nil
}
