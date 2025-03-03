package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Основная структура конфигурации
type Config struct {
	Env         string        `yaml:"env" env:"ENV" env-default:"development"`           // Режим работы (development, production)
	Address     string        `yaml:"address" env:"ADDRESS" env-default:"0.0.0.0"`       // Адрес сервера
	Port        int           `yaml:"port" env:"PORT" env-default:"5000"`                // Порт сервера
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"5s"`            // Таймаут запроса
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"` // Idle таймаут
	Services    Services      `yaml:"services"`                                          // Настройки микросервисов
}

// Структура для настроек микросервисов
type Services struct {
	Apps               string `yaml:"apps_addr" env:"APPS_ADDR" env-default:"0.0.0.0:5010"`                                 // Адрес сервиса apps
	Locations          string `yaml:"locations_addr" env:"LOCATIONS_ADDR" env-default:"0.0.0.0:5011"`                       // Адрес сервиса locations
	Movements          string `yaml:"movements_addr" env:"MOVEMENTS_ADDR" env-default:"0.0.0.0:5012"`                       // Адрес сервиса movements
	ProductionTasks    string `yaml:"production_tasks_addr" env:"PRODUCTION_TASKS_ADDR" env-default:"0.0.0.0:5013"`         // Адрес сервиса production_tasks
	ProductSK          string `yaml:"product_sk_addr" env:"PRODUCT_SK_ADDR" env-default:"0.0.0.0:5014"`                     // Адрес сервиса product_sk
	ProductsSKStatuses string `yaml:"products_sk_statuses_addr" env:"PRODUCTS_SK_STATUSES_ADDR" env-default:"0.0.0.0:5015"` // Адрес сервиса products_sk_statuses
	Sso                string `yaml:"sso_addr" env:"SSO_ADDR" env-default:"0.0.0.0:5016"`                                   // Адрес сервиса sso
	Statuses           string `yaml:"statuses_addr" env:"STATUSES_ADDR" env-default:"0.0.0.0:5017"`                         // Адрес сервиса statuses
}

var Cfg *Config

func Initialize() {
	Cfg = MustLoad()
}

func MustLoad() *Config {
	configPath := fetchConfigFlag()
	if configPath == "" {
		return loadingDataInEnv()
	}
	return MustLoadByPath(configPath)
}

func MustLoadByPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
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

func loadingDataInEnv() *Config {
	loadEnv()

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		log.Printf("Warning: Invalid PORT value in environment variables, using default value %d.", 5000)
		port = 5000
	}

	return &Config{
		Env:         os.Getenv("ENV"),
		Address:     os.Getenv("ADDRESS"),
		Port:        port,
		Timeout:     parseDuration(os.Getenv("TIMEOUT"), 5*time.Second),
		IdleTimeout: parseDuration(os.Getenv("IDLE_TIMEOUT"), 60*time.Second),
		Services: Services{
			Apps:               os.Getenv("APPS_ADDR"),
			Locations:          os.Getenv("LOCATIONS_ADDR"),
			Movements:          os.Getenv("MOVEMENTS_ADDR"),
			ProductionTasks:    os.Getenv("PRODUCTION_TASKS_ADDR"),
			ProductSK:          os.Getenv("PRODUCT_SK_ADDR"),
			ProductsSKStatuses: os.Getenv("PRODUCTS_SK_STATUSES_ADDR"),
			Sso:                os.Getenv("SSO_ADDR"),
			Statuses:           os.Getenv("STATUSES_ADDR"),
		},
	}
}

func loadEnv() {
	if err := godotenv.Load(".gateway.env"); err != nil {
		log.Println("Warning: .gateway.env file not found, using default values.")
	}
}

func parseDuration(value string, defaultValue time.Duration) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil || duration <= 0 {
		log.Printf("Warning: Invalid TIMEOUT or IDLE_TIMEOUT value in environment variables, using default value %v.", defaultValue)
		return defaultValue
	}
	return duration
}
