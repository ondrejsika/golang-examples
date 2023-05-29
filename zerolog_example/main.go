package main

import (
	"flag"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")
	pretty := flag.Bool("pretty", false, "sets pretty to console")

	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	if *pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	log.Debug().
		Str("module", "main").Int("instance", 1).
		Msg("Debug message ...")
	log.Info().
		Str("module", "main").Int("instance", 1).
		Msg("Info message ...")
	log.Warn().
		Str("module", "main").Int("instance", 1).
		Msg("Warn message ...")
	log.Error().
		Str("module", "main").Int("instance", 1).
		Msg("Error message ...")
	log.Debug().
		Str("module", "main").Int("instance", 1).
		Msg("Debug message ...")
	log.Info().
		Str("module", "server").Int("instance", 1).
		Str("method", "GET").Str("path", "/").
		Send()
}
