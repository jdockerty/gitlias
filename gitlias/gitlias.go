package gitlias

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5/config"
)

const scope = config.GlobalScope

const generateTemplate = `
[alias]

    [alias.personal]
    user = ""
    email = ""

    [alias.work]
    user = ""
    email = ""
`

func Generate() string {
	return generateTemplate
}

func List(configPath string) ([]string, string) {
	userConfig, err := Get(configPath)
	if err != nil {
		fmt.Printf("Unable to find configuration file: %s\n", err)
		os.Exit(1)
	}

	var aliases []string
	for key := range userConfig.Alias {
		aliases = append(aliases, key)
	}

	return aliases, userConfig.Current
}

func Add(configPath, alias, user, email string) error {
	userConfig, err := Get(configPath)
	if err != nil {
		fmt.Printf("Unable to find configuration file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Before: %+v\n", userConfig.Alias)

	userConfig.Alias[alias] = Alias{User: user, Email: email}

	fmt.Printf("After: %+v\n", userConfig.Alias)

	userConfig.WriteConfig(configPath)

	return nil
}

func switchAlias(confPath string, alias string, gitConf *config.Config, userConf *Gitlias) error {

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
	userConf.Current = alias
	err = userConf.WriteConfig(confPath)
	if err != nil {
		return err
	}

	fmt.Printf("Switched to alias: %s\n", alias)
	return nil
}
