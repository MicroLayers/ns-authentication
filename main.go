package main

import (
	"fmt"
	"os"

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

func main() {
	fmt.Println("This is a plugin, it will now exit")
	os.Exit(1)
}

// nolint: deadcode
var NetServerModule AuthenticationModule
