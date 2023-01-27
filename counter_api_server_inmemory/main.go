package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Counter  int    `json:"counter"`
	Hostname string `json:"hostname"`
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	HOSTNAME, _ := os.Hostname()
	COUNTER := 0

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(Response{
			Hostname: HOSTNAME,
			Counter:  COUNTER,
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		log.Debug().
			Str("hostname", HOSTNAME).
			Int("counter", COUNTER).
			Msgf("counter=%d", COUNTER)
		COUNTER++
	})

	log.Info().
		Str("hostname", HOSTNAME).
		Msg("Starting server on port :8000, http://127.0.0.1:8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Error().
			Str("hostname", HOSTNAME).
			Msgf("error=%s", err)
	}
}
