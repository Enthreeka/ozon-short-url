package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type (
	Config struct {
		Postgres    Postgres    `json:"postgres"`
		Redis       Redis       `json:"redis"`
		HTTTPServer HTTTPServer `json:"http_server"`
		GRPCServer  GRPCServer  `json:"grpc_server"`
	}

	Postgres struct {
		URL string `json:"url"`
	}

	Redis struct {
		Password     string `json:"password"`
		Host         string `json:"host"`
		DB           int    `json:"db"`
		MinIdleConns int    `json:"min_idle_conns"`
	}

	HTTTPServer struct {
		Hostname   string `json:"hostname"`
		Port       string `json:"port"`
		TypeServer string `json:"type_server"`
	}

	GRPCServer struct {
		Port string `json:"port"`
	}
)

func New() (*Config, error) {
	err := godotenv.Load("configs/config.env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		Postgres: Postgres{
			URL: os.Getenv("POSTGRES_URL"),
		},
		Redis: Redis{
			Password:     os.Getenv("REDIS_PASSWORD"),
			Host:         os.Getenv("REDIS_HOST"),
			DB:           parseEnvInt(os.Getenv("REDIS_DB")),
			MinIdleConns: parseEnvInt(os.Getenv("REDIS_MIN_IDLE_COONS")),
		},
		HTTTPServer: HTTTPServer{
			Hostname:   os.Getenv("HTTP_HOSTNAME"),
			Port:       os.Getenv("HTTP_PORT"),
			TypeServer: os.Getenv("HTTP_TYPE_SERVER"),
		},
		GRPCServer: GRPCServer{
			Port: os.Getenv("GRPC_SERVER_PORTs"),
		},
	}

	return config, nil
}

func parseEnvInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}
