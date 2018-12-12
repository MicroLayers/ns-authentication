package configuration_test

import (
	"testing"

	"ns-auth/configuration"

	"github.com/stretchr/testify/assert"

	yaml "gopkg.in/yaml.v2"
)

func TestParseConfiguration(t *testing.T) {
	configYaml := `
authentication:
  store:
    type: 'mongo'
    connectionString: 'mongo:27017'
    database: 'authentication'
  cache:
    enabled: true
    type: 'redis'
    connectionString: 'redis:6379'
`
	bytes := []byte(configYaml)
	var mapSlice yaml.MapSlice
	yaml.Unmarshal(bytes, &mapSlice)

	config, err := configuration.ReadConfiguration(mapSlice)

	assert.NoError(t, err)
	assert.Equal(t, "mongo", config.Authentication.Store.Type)
	assert.Equal(t, "mongo:27017", config.Authentication.Store.ConnectionString)
	assert.Equal(t, "authentication", config.Authentication.Store.Database)
	assert.Equal(t, true, config.Authentication.Cache.Enabled)
	assert.Equal(t, "redis", config.Authentication.Cache.Type)
	assert.Equal(t, "redis:6379", config.Authentication.Cache.ConnectionString)
}
