package gitlias

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Gitlias struct {
	Alias   map[string]Alias `toml:"alias"`
	Current string           `toml:"current"`
}

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

type Alias struct {
	User  string `toml:"user"`
	Email string `toml:"email"`
}

func Get(filePath string) (*Gitlias, error) {
	var c *Gitlias

	_, err := toml.DecodeFile(filePath, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
