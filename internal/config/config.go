package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// HTTPServer holds the network configuration settings
type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// Config represents the application's root configuration structure
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

// MustLoad retrieves application configuration or terminates on failure
func MustLoad() *Config {
	var configPath string

	// Attempt to load configuration path from environment variables
	configPath = os.Getenv("CONFIG_PATH")

	// Fallback to command-line flag if environment variable is not set
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags
	}

	// Fallback to default local configuration path if still empty
	if configPath == "" {
		configPath = "config/local.yaml"
	}

	// Verify the configuration file exists before proceeding
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	// Parse the YAML configuration file into the struct
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg
}
