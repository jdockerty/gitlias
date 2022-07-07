package gitlias

import (
	"os"

	"github.com/BurntSushi/toml"
)

// Gitlias wraps the current configured git alias and the map of aliases together.
type Gitlias struct {
	Alias   map[string]Alias `toml:"alias"`
	Current string           `toml:"current"`
}

// WriteConfig is a wrapper for writing to the configuration file.
func (g *Gitlias) WriteConfig(filePath string) error {

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := toml.NewEncoder(f)

	err = enc.Encode(&g)
	if err != nil {
		return err
	}

	return nil
}

// Add is a shorthand for adding a new alias to the Alias struct.
func (g *Gitlias) Add(alias, user, email string) {
	g.Alias[alias] = Alias{User: user, Email: email}
}

// Alias represents a git author, this has a user and associated email address.
type Alias struct {
	User  string `toml:"user"`
	Email string `toml:"email"`
}

// Get is used to get the current gitlias configuration.
func Get(filePath string) (*Gitlias, error) {
	var c *Gitlias

	_, err := toml.DecodeFile(filePath, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
