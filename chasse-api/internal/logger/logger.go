package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	name string
}

type LogLevel string

const (
	LOG_ERROR   LogLevel = "ERROR"
	LOG_INFO    LogLevel = "INFO"
	LOG_DEBUG   LogLevel = "DEBUG"
	LOG_WARNING LogLevel = "WARN"
)

var logLevels = map[LogLevel]int{
	LOG_ERROR:   1,
	LOG_WARNING: 2,
	LOG_INFO:    3,
	LOG_DEBUG:   4,
}

const (
	logFormat      = "%s | %7s | %7s | %s"
	DateTimeFormat = "2006/01/02 15:04:05"
)

var (
	// default log level LOG_INFO
	logLevel       = logLevels[LOG_INFO]
	FiberLogFormat = fmt.Sprintf(logFormat, "${time}", LOG_INFO, "FIBER", "(${ip}:${port}) ${method} ${status} - ${path}\n")
)

// name -> module name, will be visible in log message
func New(name string) *Logger {
	return &Logger{name}
}

func SetGlobalLogLevel(level LogLevel) error {
	if v, ok := logLevels[level]; ok {
		logLevel = v
		return nil
	}

	return fmt.Errorf("abort SetGlobalLogLevel -> bad log level type: %s", string(level))
}

func (l *Logger) Infof(format string, v ...any) {
	l.write(LOG_INFO, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(message string) {
	l.write(LOG_INFO, message)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.write(LOG_ERROR, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(message string) {
	l.write(LOG_ERROR, message)
}

func (l *Logger) Warningf(format string, v ...any) {
	l.write(LOG_WARNING, fmt.Sprintf(format, v...))
}

func (l *Logger) Warning(message string) {
	l.write(LOG_WARNING, message)
}

func (l *Logger) Debugf(format string, v ...any) {
	l.write(LOG_DEBUG, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(message string) {
	l.write(LOG_DEBUG, message)
}

func (l *Logger) write(level LogLevel, message string) {
	if v, ok := logLevels[level]; ok && v <= logLevel {
		var (
			len  = len(message)
			date = time.Now().Format(DateTimeFormat)
		)

		if len > 0 && message[len-1] != '\n' {
			message = message + "\n"
		}

		fmt.Printf(logFormat, date, level, l.name, message)
	}
}
