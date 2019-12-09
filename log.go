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

	// Setup timesampe
	if l.UTC {
		if l.Date {
			prefix += fmt.Sprintf("%v ", time.Now().UTC().Format("2006-01-02 15:04:05.000"))
		} else {
			prefix += fmt.Sprintf("%v ", time.Now().UTC().Format("15:04:05.000"))
		}
	} else {
		if l.Date {
			prefix += fmt.Sprintf("%v ", time.Now().Format("2006-01-02 15:04:05.000"))
		} else {
			prefix += fmt.Sprintf("%v ", time.Now().Format("15:04:05.000"))
		}
	}

	// Logging level
	prefix += logLevel + " "

	// Caller location
	if l.Function {
		_, file, line, _ := runtime.Caller(2)
		prefix += fmt.Sprintf("[%v:%v] ", filepath.Base(file), line)
	}

	return (prefix + msg)
}

// Debug logs debug messages
func (l *Logger) Debug(msg string) string {
	if l.Level >= level["DEBUG"] {
		s := l.format(l.color("DEBUG", GRAY), msg)
		fmt.Println(s)
		return s
	}
	return ""
}

// Debugf logs debug messages
func (l *Logger) Debugf(format string, args ...interface{}) string {
	if l.Level >= level["DEBUG"] {
		s := l.format(l.color("DEBUG", GRAY), fmt.Sprintf(format, args...))
		fmt.Println(s)
		return s
	}
	return ""
}

// Trace logs trace messages
func (l *Logger) Trace(msg string) string {
	if l.Level >= level["TRACE"] {
		s := l.format(l.color("TRACE", CYAN), msg)
		fmt.Println(s)
		return s
	}
	return ""
}

// Tracef logs trace messages
func (l *Logger) Tracef(format string, args ...interface{}) string {
	if l.Level >= level["TRACE"] {
		s := l.format(l.color("TRACE", CYAN), fmt.Sprintf(format, args...))
		fmt.Println(s)
		return s
	}
	return ""
}

// Info logs info messages
func (l *Logger) Info(msg string) string {
	if l.Level >= level["INFO"] {
		s := l.format(l.color("INFO", BLUE), msg)
		fmt.Println(s)
		return s
	}
	return ""
}

// Infof logs imfo messages
func (l *Logger) Infof(format string, args ...interface{}) string {
	if l.Level >= level["INFO"] {
		s := l.format(l.color("INFO", BLUE), fmt.Sprintf(format, args...))
		fmt.Println(s)
		return s
	}
	return ""
}

// Warn logs warn messages
func (l *Logger) Warn(msg string) string {
	if l.Level >= level["WARN"] {
		s := l.format(l.color("WARN", YELLOW), msg)
		fmt.Println(s)
		return s
	}
	return ""
}

// Warnf logs wann messages
func (l *Logger) Warnf(format string, args ...interface{}) string {
	if l.Level >= level["WARN"] {
		s := l.format(l.color("WARN", YELLOW), fmt.Sprintf(format, args...))
		fmt.Println(s)
		return s
	}
	return ""
}

// Error logs error messages
func (l *Logger) Error(msg string) string {
	if l.Level >= level["ERROR"] {
		s := l.format(l.color("ERROR", RED), msg)
		fmt.Println(s)
		return s
	}
	return ""
}

// Errorf logs error messages
func (l *Logger) Errorf(format string, args ...interface{}) string {
	if l.Level >= level["ERROR"] {
		s := l.format(l.color("ERROR", RED), fmt.Sprintf(format, args...))
		fmt.Println(s)
		return s
	}
	return ""
}

// Fatal logs fatal message and exits (1)
func (l *Logger) Fatal(msg string) string {
	s := l.format(l.color("FATAL", MAGENTA), msg)
	fmt.Println(s)
	defer os.Exit(1)
	return s
}

// Fatalf logs fatal message and exits (1)
func (l *Logger) Fatalf(format string, args ...interface{}) string {
	s := l.format(l.color("FATAL", MAGENTA), fmt.Sprintf(format, args...))
	fmt.Println(s)
	defer os.Exit(1)
	return s
}

// Panic logs fatal message and exits (1)
func (l *Logger) Panic(msg string) string {
	s := l.format(l.color("PANIC", DMAGENTA), msg)
	defer panic(s)
	return s
}

// Panicf logs fatal message and exits (1)
func (l *Logger) Panicf(format string, args ...interface{}) string {
	s := l.format(l.color("PANIC", DMAGENTA), fmt.Sprintf(format, args...))
	defer panic(s)
	return s
}
