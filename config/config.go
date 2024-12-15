package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Server ServerConfig
	DB     DatabaseConfig
}

type ServerConfig struct {
	Port         int           `mapstructure:"SERVER_PORT"`
	ReadTimeout  time.Duration `mapstructure:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `mapstructure:"SERVER_WRITE_TIMEOUT"`
	IdleTimeout  time.Duration `mapstructure:"SERVER_IDLE_TIMEOUT"`
	Debug        bool          `mapstructure:"SERVER_DEBUG"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

func NewConfig() *Config {
	v := viper.New()

	// Set defaults
	v.SetDefault("SERVER_PORT", 8080)
	v.SetDefault("SERVER_READ_TIMEOUT", "30s")
	v.SetDefault("SERVER_WRITE_TIMEOUT", "30s")
	v.SetDefault("SERVER_IDLE_TIMEOUT", "60s")
	v.SetDefault("SERVER_DEBUG", false)

	// Environment variable binding
	v.AutomaticEnv()

	// Load configuration from file if it exists
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AddConfigPath("./")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
	}

	config.Server.ReadTimeout = parseDuration(v.GetString("SERVER_READ_TIMEOUT"))
	config.Server.WriteTimeout = parseDuration(v.GetString("SERVER_WRITE_TIMEOUT"))
	config.Server.IdleTimeout = parseDuration(v.GetString("SERVER_IDLE_TIMEOUT"))

	return &config
}

func NewDB() *DatabaseConfig {
	v := viper.New()

	v.AutomaticEnv()

	dbConfig := &DatabaseConfig{
		Host:     v.GetString("DB_HOST"),
		Port:     v.GetInt("DB_PORT"),
		User:     v.GetString("DB_USER"),
		Password: v.GetString("DB_PASSWORD"),
		Name:     v.GetString("DB_NAME"),
	}

	return dbConfig
}

func parseDuration(duration string) time.Duration {
	d, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalf("Error parsing duration, %s", err)
	}
	return d
}
