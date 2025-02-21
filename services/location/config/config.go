package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Env        string         `yaml:"env" env:"ENV" env-default:"development"`
	GRPCServer GRPCServer     `yaml:"grpc_server"`
	DBConfig   DatabaseConfig `yaml:"database"`
}

type GRPCServer struct {
	Address     string        `yaml:"address" env:"ADDRESS" env-default:"0.0.0.0"`
	Port        int           `yaml:"port" env:"PORT" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"10s"`
}

type DatabaseConfig struct {
	MSSQL MSSQLConfig `yaml:"mssql"`
}

type MSSQLConfig struct {
	ConnectionString   string        `yaml:"connection_string" env:"MSSQL_CONNECTION_STRING"`
	MaxOpenConnections int           `yaml:"max_open_connections" env:"MSSQL_MAX_OPEN_CONNECTIONS" env-default:"10"`
	MaxIdleConnections int           `yaml:"max_idle_connections" env:"MSSQL_MAX_IDLE_CONNECTIONS" env-default:"5"`
	ConnMaxLifetime    time.Duration `yaml:"conn_max_lifetime" env:"MSSQL_CONN_MAX_LIFETIME" env-default:"5m"`
}

var Cfg *Config

func Initialize() {
	Cfg = MustLoad()
}

func MustLoad() *Config {
	configPath := fetchConfigFlag()
	if configPath != "" {
		return loadFromYAML(configPath)
	}
	return loadFromEnv()
}

func loadFromYAML(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("Failed to parse YAML config: %v", err)
	}

	return &cfg
}

func loadFromEnv() *Config {
	loadEnv()

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Failed to read environment variables: %v", err)
	}

	return &cfg
}

func fetchConfigFlag() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}

func loadEnv() {
	if err := godotenv.Load(".sso.env"); err != nil {
		log.Println("Warning: .sso.env file not found, using default values.")
	}
}
