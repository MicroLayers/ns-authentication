package main

import (
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"ns-auth/configuration"
	"ns-auth/handler"
)

type AuthenticationModule struct {
	configuration configuration.Configuration
	protoHandler  *handler.ProtoHandler
}

func echo(data []byte) []byte {
	// Dumb
	return data
}

func (m *AuthenticationModule) HandleJSON(data []byte) []byte {
	return echo(data)
}

func (m *AuthenticationModule) HandleProto(data []byte) []byte {
	return m.protoHandler.HandleRequest(data)
}

func (m *AuthenticationModule) Init(rawConfig yaml.MapSlice) {
	config, err := configuration.ReadConfiguration(rawConfig)

	if err != nil {
		log.Fatal(err)
	}

	m.configuration = config
	m.protoHandler = handler.GetProtoHandler(&m.configuration)
}

var NetServerModule AuthenticationModule
