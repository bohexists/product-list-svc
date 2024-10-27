package config

import "os"

type Config struct {
	FetcherServiceAddr string
}

func LoadConfig() Config {
	return Config{
		FetcherServiceAddr: os.Getenv("FETCHER_SERVICE_ADDR"),
	}
}
