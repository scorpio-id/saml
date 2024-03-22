package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config provides a template for marshalling .yml configuration files
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	PKI struct {
		Enabled bool     `yaml:"enabled"`
		URL     string   `yaml:"url"`
		SANs    []string `yaml:"sans"`
	} `yaml:"pki"`
	OAuth struct {
		Enabled  bool   `yaml:"enabled"`
		Issuer   string `yaml:"issuer"`
		ClientID string `yaml:"client_id"`
	} `yaml:"oauth"`
}

// NewConfig takes a .yml filename from the same /config directory, and returns a populated configuration
func NewConfig(s string) Config {
	f, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}