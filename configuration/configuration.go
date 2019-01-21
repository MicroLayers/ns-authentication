package configuration

import (
	yaml "gopkg.in/yaml.v2"
)

// Configuration the plugin's configuration
type Configuration struct {
	Authentication struct {
		Hasher struct {
			Type string `yaml:"type"`
			Salt string `yaml:"salt"`
		} `yaml:"hasher"`
		Store struct {
			Type             string `yaml:"type"`
			ConnectionString string `yaml:"connectionString"`
			Database         string `yaml:"database"`
		} `yaml:"store"`
		Cache struct {
			Enabled          bool   `yaml:"enabled"`
			Type             string `yaml:"type"`
			ConnectionString string `yaml:"connectionString"`
		} `yaml:"cache"`
	} `yaml:"authentication"`
}

// ReadConfiguration read the configuration, starting from a raw YML configuration
func ReadConfiguration(rawConfig yaml.MapSlice) (Configuration, error) {
	var config Configuration

	marshaled, err := yaml.Marshal(&rawConfig)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(marshaled, &config)

	return config, err
}
