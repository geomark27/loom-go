package helpers

import (
	"fmt"
	"log"
	"os"
)

// Logger interfaz para logging estructurado
type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
}

// LogLevel niveles de log
type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// DefaultLogger implementación por defecto
type DefaultLogger struct {
	logger *log.Logger
	level  LogLevel
}

// NewLogger crea un nuevo logger con nivel Info por defecto
func NewLogger() Logger {
	return &DefaultLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		level:  InfoLevel,
	}
}

// NewLoggerWithLevel crea un logger con nivel específico
func NewLoggerWithLevel(level LogLevel) Logger {
	return &DefaultLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		level:  level,
	}
}

func (l *DefaultLogger) Info(msg string, keysAndValues ...interface{}) {
	if l.level <= InfoLevel {
		l.logger.Printf("[INFO] %s %s", msg, formatKeyValues(keysAndValues...))
	}
}

func (l *DefaultLogger) Error(msg string, keysAndValues ...interface{}) {
	if l.level <= ErrorLevel {
		l.logger.Printf("[ERROR] %s %s", msg, formatKeyValues(keysAndValues...))
	}
}

func (l *DefaultLogger) Debug(msg string, keysAndValues ...interface{}) {
	if l.level <= DebugLevel {
		l.logger.Printf("[DEBUG] %s %s", msg, formatKeyValues(keysAndValues...))
	}
}

func (l *DefaultLogger) Warn(msg string, keysAndValues ...interface{}) {
	if l.level <= WarnLevel {
		l.logger.Printf("[WARN] %s %s", msg, formatKeyValues(keysAndValues...))
	}
}

func (l *DefaultLogger) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalf("[FATAL] %s %s", msg, formatKeyValues(keysAndValues...))
}

// formatKeyValues formatea pares clave-valor para logging
func formatKeyValues(keysAndValues ...interface{}) string {
	if len(keysAndValues) == 0 {
		return ""
	}

	result := ""
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			result += fmt.Sprintf("%v=%v ", keysAndValues[i], keysAndValues[i+1])
		}
	}
	return result
}
