package logger

import (
	"fmt"
	"strings"
	"time"
)

type Logger struct {
	name string
}

type LogLevel string

const (
	levelFatal   LogLevel = "FATAL"
	levelError   LogLevel = "ERROR"
	levelWarning LogLevel = "WARN"
	levelInfo    LogLevel = "INFO"
	levelDebug   LogLevel = "DEBUG"
)

var logLevels = map[LogLevel]int{
	levelFatal:   0,
	levelError:   1,
	levelWarning: 2,
	levelInfo:    3,
	levelDebug:   4,
}

const (
	logFormat      = "%s | %7s | %7s | %s"
	DateTimeFormat = "2006/01/02 15:04:05"
)

var (
	// default log level INFO
	logLevel       = logLevels[levelInfo]
	FiberLogFormat = fmt.Sprintf(logFormat, "${time}", levelInfo, "FIBER", "(${ip}:${port}) ${method} ${status} - ${path}\n")
)

// name -> module name, will be visible in log message
func New(name string) *Logger {
	return &Logger{name}
}

func SetGlobalLogLevel(level LogLevel) error {
	if v, ok := logLevels[LogLevel(strings.ToUpper(string(level)))]; ok {
		logLevel = v
		return nil
	}

	return fmt.Errorf("abort SetGlobalLogLevel -> bad log level type: %s", string(level))
}

func (l *Logger) Infof(format string, v ...any) {
	l.write(levelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(message string) {
	l.write(levelInfo, message)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.write(levelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(message string) {
	l.write(levelError, message)
}

func (l *Logger) Warningf(format string, v ...any) {
	l.write(levelWarning, fmt.Sprintf(format, v...))
}

func (l *Logger) Warning(message string) {
	l.write(levelWarning, message)
}

func (l *Logger) Debugf(format string, v ...any) {
	l.write(levelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(message string) {
	l.write(levelDebug, message)
}

func (l *Logger) Fatalf(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	l.write(levelFatal, message)
	panic(message)
}

func (l *Logger) Fatal(message string) {
	l.write(levelFatal, message)
	panic(message)
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
