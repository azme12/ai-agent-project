package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[AI-AGENT] ", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.Printf("[INFO] "+msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.Printf("[ERROR] "+msg, args...)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Printf("[DEBUG] "+msg, args...)
}
