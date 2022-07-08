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

// Remove is used to remove an element from the alias mapping by specifying the alias name.
func Remove(configPath, alias string) (*Gitlias, error) {
	userConfig, err := Get(configPath)
	if err != nil {
		return nil, err
	}

	// Go's builtin for removing an element from a map.
	// This is a no-op when the key doesn't exist, so there is no need to check for existence.
	delete(userConfig.Alias, alias)

	return userConfig, nil
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
