package configuration

import (
	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	Authentication struct {
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

func ReadConfiguration(rawConfig yaml.MapSlice) (Configuration, error) {
	var config Configuration

	marshaled, err := yaml.Marshal(&rawConfig)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(marshaled, &config)

	return config, err
}
