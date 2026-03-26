package config

import (
	"os"
	"time"
)

func getenvAny(keys ...string) string {
	for _, key := range keys {
		v := os.Getenv(key)
		if v != "" {
			return v
		}
	}
	return ""
}

// Config holds all configuration for the bot
type Config struct {
	// Discord
	DiscordToken string

	// Database
	DatabaseURL string

	// Logging
	LogLevel string

	// Bot behavior
	Timezone   string
	Environment string

	// Timeouts
	DatabaseTimeout time.Duration
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	timezone := getenvAny("TIMEZONE", "DEFAULT_TIMEZONE")
	if timezone == "" {
		timezone = "Asia/Tokyo"
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return &Config{
		DiscordToken:    getenvAny("DISCORD_TOKEN", "DISCORD_BOT_TOKEN"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		LogLevel:        logLevel,
		Timezone:        timezone,
		Environment:     env,
		DatabaseTimeout: 30 * time.Second,
	}
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.DiscordToken == "" {
		return ErrMissingDiscordToken
	}
	if c.DatabaseURL == "" {
		return ErrMissingDatabaseURL
	}
	return nil
}
