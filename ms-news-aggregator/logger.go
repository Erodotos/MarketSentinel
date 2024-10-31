package main

import (
	"os"

	"github.com/rs/zerolog"
)

type LoggerInstance struct {
	zerolog.Logger
}

var logger *LoggerInstance

var LOG_LEVEL = os.Getenv("LOG_LEVEL")

func Logger() LoggerInstance {

	// Reuse Logger if already created
	if logger != nil {
		return *logger
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	switch LOG_LEVEL {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	logger = &LoggerInstance{
		zerolog.New(os.Stderr).With().Timestamp().Logger(),
	}

	return *logger
}
