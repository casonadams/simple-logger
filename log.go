package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var level map[string]int = map[string]int{
	"DEBUG": 6,
	"TRACE": 5,
	"INFO":  4,
	"WARN":  3,
	"ERROR": 2,
	"FATAL": 1,
}

type color int

// Colors
const (
	DMAGENTA color = 35
	GRAY     color = 90
	RED      color = 91
	YELLOW   color = 93
	BLUE     color = 94
	MAGENTA  color = 95
	CYAN     color = 96
)

// Logger struct
type Logger struct {
	mu       sync.Mutex
	Level    int
	Date     bool
	Color    bool
	Function bool
	UTC      bool
}

// NewLogger creates a new logger
func NewLogger(name string) *Logger {
	envLevel := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	envDate := strings.ToLower(os.Getenv("LOG_DATE"))
	envColor := strings.ToLower(os.Getenv("LOG_COLOR"))
	envFunc := strings.ToLower(os.Getenv("LOG_FUNC"))
	envUTC := strings.ToLower(os.Getenv("LOG_UTC"))

	var logLevel int = 4
	if len(envLevel) > 0 {
		logLevel = level[envLevel]
	}

	var date bool = true
	if envDate == "false" || envDate == "0" {
		date = false
	}
	var lcolor bool = true
	if envColor == "false" || envColor == "0" {
		lcolor = false
	}
	var showFunc bool = true
	if envFunc == "false" || envFunc == "0" {
		showFunc = false
	}
	var tzUTC bool = true
	if envUTC == "false" || envUTC == "0" {
		tzUTC = false
	}

	return &Logger{
		Level:    logLevel,
		Date:     date,
		Color:    lcolor,
		Function: showFunc,
		UTC:      tzUTC,
	}
}

func (l *Logger) color(m string, c color) string {
	if l.Color {
		return fmt.Sprintf("\033[%dm%s\033[0m", c, m)
	}
	return m
}

func (l *Logger) format(logLevel string, msg string) string {
	prefix := ""
	if l.Level >= level[logLevel] {
		if l.Date {
			prefix += fmt.Sprintf("%v ", time.Now().UTC().Format("2006-01-02 15:04:05.000"))
		} else {
			prefix += fmt.Sprintf("%v ", time.Now().UTC().Format("15:04:05.000"))
		}

		switch logLevel {
		case "DEBUG":
			prefix += l.color(logLevel, GRAY)
		case "TRACE":
			prefix += l.color(logLevel, CYAN)
		case "INFO":
			prefix += l.color(logLevel, BLUE)
		case "WARN":
			prefix += l.color(logLevel, YELLOW)
		case "ERROR":
			prefix += l.color(logLevel, RED)
		case "FATAL":
			prefix += l.color(logLevel, MAGENTA)
		case "PANIC":
			prefix += l.color(logLevel, DMAGENTA)
		default:
			prefix += logLevel
		}

		// need a space here
		prefix += " "

		if l.Function {
			_, file, line, _ := runtime.Caller(2)
			prefix += fmt.Sprintf("[%v:%v] ", filepath.Base(file), line)
		}

		return (prefix + " " + msg)
	}
	return ""
}

// Debug logs debug messages
func (l *Logger) Debug(msg string) {
	s := l.format("DEBUG", msg)
	if s != "" {
		fmt.Println(s)
	}
}

// Debugf logs debug messages
func (l *Logger) Debugf(format string, args ...interface{}) {
	s := l.format("DEBUG", fmt.Sprintf(format, args...))
	if s != "" {
		fmt.Println(s)
	}
}

// Trace logs trace messages
func (l *Logger) Trace(msg string) {
	s := l.format("TRACE", msg)
	if s != "" {
		fmt.Println(s)
	}
}

// Tracef logs trace messages
func (l *Logger) Tracef(format string, args ...interface{}) {
	s := l.format("TRACE", fmt.Sprintf(format, args...))
	if s != "" {
		fmt.Println(s)
	}
}

// Info logs info messages
func (l *Logger) Info(msg string) {
	s := l.format("INFO", msg)
	if s != "" {
		fmt.Println(s)
	}
}

// Infof logs imfo messages
func (l *Logger) Infof(format string, args ...interface{}) {
	s := l.format("INFO", fmt.Sprintf(format, args...))
	if s != "" {
		fmt.Println(s)
	}
}

// Warn logs warn messages
func (l *Logger) Warn(msg string) {
	s := l.format("WARN", msg)
	if s != "" {
		fmt.Println(s)
	}
}

// Warnf logs wann messages
func (l *Logger) Warnf(format string, args ...interface{}) {
	s := l.format("WARN", fmt.Sprintf(format, args...))
	if s != "" {
		fmt.Println(s)
	}
}

// Error logs error messages
func (l *Logger) Error(msg string) {
	s := l.format("ERROR", msg)
	if s != "" {
		fmt.Println(s)
	}
}

// Errorf logs error messages
func (l *Logger) Errorf(format string, args ...interface{}) {
	s := l.format("ERROR", fmt.Sprintf(format, args...))
	if s != "" {
		fmt.Println(s)
	}
}

// Fatal logs fatal message and exits (1)
func (l *Logger) Fatal(msg string) {
	s := l.format("FATAL", msg)
	fmt.Println(s)
	os.Exit(1)
}

// Fatalf logs fatal message and exits (1)
func (l *Logger) Fatalf(format string, args ...interface{}) {
	s := l.format("FATAL", fmt.Sprintf(format, args...))
	fmt.Println(s)
	os.Exit(1)
}

// Panic logs fatal message and exits (1)
func (l *Logger) Panic(msg string) {
	s := l.format("PANIC", msg)
	panic(s)
}

// Panicf logs fatal message and exits (1)
func (l *Logger) Panicf(format string, args ...interface{}) {
	s := l.format("PANIC", fmt.Sprintf(format, args...))
	panic(s)
}
