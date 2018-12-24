package handler_test

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	// Prevent log spam, enabling debug messages only
	log.SetLevel(log.DebugLevel)
}
