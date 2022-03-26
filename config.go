package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Gitlias struct {
	Alias map[string]Alias
}

type Alias struct {
	User  string `mapstructure:"user"`
	Email string `mapstructure:"email"`
}

func Get(filePath string) (*Gitlias, error) {
	var c *Gitlias

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	_, err = toml.Decode(string(data), &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
