// Package config wraps tewi's configuration
package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config maps the config.toml file directly
type Config struct {
	PostgresConfig PostgresConfig `toml:"postgres"`
	Boards         []string       `toml:"boards"`
	Port           int            `toml:"port"`
}

// PostgresConfig is the config for postgres
type PostgresConfig struct {
	ConnectionString string `toml:"connection_string"`
}

// LoadConfig loads the config.toml file and returns a Config
func LoadConfig() Config {
	configFile := os.Getenv("TEWI_CONFIG")

	if configFile == "" {
		configFile = "./config.toml"
	}

	f, err := os.Open(configFile)

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	var conf Config

	if _, err := toml.NewDecoder(f).Decode(&conf); err != nil {
		log.Fatalln(err)
	}

	return conf
}
