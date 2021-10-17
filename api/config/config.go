package config


import (
	"os"
	"strconv"
)

type Config struct {
	HttpPort int
	ConnectionString string
}

func NewConfig() (Config, error) {
	var cfg Config
	httpPort, err := strconv.Atoi(getEnv("HTTP_PORT", "8080"))
	if err != nil {
		return cfg, err
	}
	return Config{
		HttpPort: httpPort,
		ConnectionString: getEnv("CONNECTION_STRING", "./urls.db"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}