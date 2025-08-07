package config

import (
	"os"
)

type Config struct {
	GoogleCalendarAPIKey string
	SendGridAPIKey       string
	GeminiAPIKey         string
	ServerPort           string
	LogLevel             string
}

func Load() (*Config, error) {
	return &Config{
		GoogleCalendarAPIKey: getEnv("GOOGLE_CALENDAR_API_KEY", ""),
		SendGridAPIKey:       getEnv("SENDGRID_API_KEY", ""),
		GeminiAPIKey:         getEnv("GEMINI_API_KEY", ""),
		ServerPort:           getEnv("SERVER_PORT", "8080"),
		LogLevel:             getEnv("LOG_LEVEL", "info"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
