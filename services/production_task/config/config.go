package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string         `yaml:"env" env:"ENV" env-default:"development"`
	GRPCServer GRPCServer     `yaml:"grpc_server"`
	DBConfig   DatabaseConfig `yaml:"database"`
}

type GRPCServer struct {
	Address     string        `yaml:"address" env:"ADDRESS"`
	Port        int           `yaml:"port" env:"PORT"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT"`
}

type DatabaseConfig struct {
	Firebird FirebirdConfig `yaml:"firebird"`
}

type FirebirdConfig struct {
	ConnectionString   string        `yaml:"connection_string" env:"FIREBIRD_CONNECTION_STRING"`
	MaxOpenConnections int           `yaml:"max_open_connections" env:"FIREBIRD_MAX_OPEN_CONNECTIONS" env-default:"10"`
	MaxIdleConnections int           `yaml:"max_idle_connections" env:"FIREBIRD_MAX_IDLE_CONNECTIONS" env-default:"5"`
	ConnMaxLifetime    time.Duration `yaml:"conn_max_lifetime" env:"FIREBIRD_CONN_MAX_LIFETIME" env-default:"5m"`
}

type ServiceConfig struct {
	Address string `yaml:"address" env:"SERVICE_ADDRESS"`
	Port    int    `yaml:"port" env:"SERVICE_PORT"`
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
	if cfg.GRPCServer.Timeout == 0 {
		cfg.GRPCServer.Timeout = 4 * time.Second
	}
	if cfg.GRPCServer.IdleTimeout == 0 {
		cfg.GRPCServer.IdleTimeout = 60 * time.Second
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
