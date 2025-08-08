package config

import (
	"os"
	"strconv"
)

type Config struct {
	// API Keys
	GoogleCalendarAPIKey string
	SendGridAPIKey       string
	GeminiAPIKey         string

	// Server Configuration
	ServerPort string
	LogLevel   string

	// API URLs and Endpoints
	GoogleCalendarURL string
	SendGridURL       string
	GeminiURL         string

	// Email Configuration
	FromEmail string
	FromName  string
	UserEmail string // User's email for receiving notifications

	// Calendar Configuration
	CalendarID string
	TimeZone   string

	// Scheduler Configuration
	DailyReminderTime      string
	MeetingReminderMinutes int
}

func Load() (*Config, error) {
	return &Config{
		// API Keys
		GoogleCalendarAPIKey: getEnv("GOOGLE_CALENDAR_API_KEY", ""),
		SendGridAPIKey:       getEnv("SENDGRID_API_KEY", ""),
		GeminiAPIKey:         getEnv("GEMINI_API_KEY", ""),

		// Server Configuration
		ServerPort: getEnv("SERVER_PORT", "8080"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),

		// API URLs and Endpoints
		GoogleCalendarURL: getEnv("GOOGLE_CALENDAR_URL", "https://www.googleapis.com/calendar/v3"),
		SendGridURL:       getEnv("SENDGRID_URL", "https://api.sendgrid.com/v3"),
		GeminiURL:         getEnv("GEMINI_URL", "https://generativelanguage.googleapis.com/v1beta"),

		// Email Configuration
		FromEmail: getEnv("FROM_EMAIL", "azmetefera07@gmail.com"),
		FromName:  getEnv("FROM_NAME", "AI Assistant"),
		UserEmail: getEnv("USER_EMAIL", "azmetefera07@gmail.com"),

		// Calendar Configuration
		CalendarID: getEnv("CALENDAR_ID", "primary"),
		TimeZone:   getEnv("TIMEZONE", "UTC"),

		// Scheduler Configuration
		DailyReminderTime:      getEnv("DAILY_REMINDER_TIME", "09:00"),
		MeetingReminderMinutes: getEnvAsInt("MEETING_REMINDER_MINUTES", 15),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
