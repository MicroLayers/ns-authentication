package main

import (
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"ns-auth/configuration"
)

type AuthenticationModule struct {
	configuration configuration.Configuration
}

func echo(rawConfig yaml.MapSlice, data []byte) []byte {
	log.WithField("data", data).Info("")
	return data
}

func (m *AuthenticationModule) HandleJSON(rawConfig yaml.MapSlice, data []byte) []byte {
	return echo(rawConfig, data)
}

func (m *AuthenticationModule) HandleProto(rawConfig yaml.MapSlice, data []byte) []byte {
	return echo(rawConfig, data)
}

func (m *AuthenticationModule) Init(rawConfig yaml.MapSlice) {

}

var NetServerModule AuthenticationModule
