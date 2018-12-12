package main

import (
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"ns-auth/configuration"
)

type AuthenticationModule struct {
	configuration configuration.Configuration
}

func echo(data []byte) []byte {
	// Dumb
	return data
}

func (m *AuthenticationModule) HandleJSON(data []byte) []byte {
	return echo(data)
}

func (m *AuthenticationModule) HandleProto(data []byte) []byte {
	return echo(data)
}

func (m *AuthenticationModule) Init(rawConfig yaml.MapSlice) {
	config, err := configuration.ReadConfiguration(rawConfig)

	if err != nil {
		log.Fatal(err)
	}

	m.configuration = config
}

var NetServerModule AuthenticationModule
