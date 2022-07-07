package gitlias

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5/config"
)

const scope = config.GlobalScope

const generateTemplate = `
current = ""

[alias]

    [alias.personal]
    user = ""
    email = ""

    [alias.work]
    user = ""
    email = ""
`

// Init is used to initialise/generate a template configuration file.
func Init() string {
	return generateTemplate
}

// List is used to list all aliases.
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

// Add is a wrapper for writing a new alias into the configuration file.
func Add(configPath, alias, user, email string) (*Gitlias, error) {
	userConfig, err := Get(configPath)
	if err != nil {
		return nil, err
	}
	userConfig.Alias[alias] = Alias{User: user, Email: email}

	return userConfig, nil
}
