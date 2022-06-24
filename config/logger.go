package config

import (
	"strings"

	"github.com/rs/zerolog"
)

func ConfigureLogLevel(level string) {
	level = strings.ToLower(strings.TrimSpace(level))

	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else {
		zerolog.SetGlobalLevel(logLevel)
	}
}
