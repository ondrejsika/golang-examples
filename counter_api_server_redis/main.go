package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Counter  int    `json:"counter"`
	Hostname string `json:"hostname"`
}

func getCouterFromRedis() int {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "127.0.0.1"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":6379",
		Password: "",
		DB:       0,
	})
	result, _ := rdb.Incr("counter").Result()
	return int(result)
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	HOSTNAME, _ := os.Hostname()

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter := getCouterFromRedis()
		data, _ := json.Marshal(Response{
			Hostname: HOSTNAME,
			Counter:  counter,
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		log.Debug().
			Str("hostname", HOSTNAME).
			Int("counter", counter).
			Msgf("counter=%d", counter)
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
