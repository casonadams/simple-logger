package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
	RED     color = 91
	GRAY    color = 37
	MAGENTA color = 95
	CYAN    color = 96
	BLUE    color = 94
	YELLOW  color = 93
)

// Logger struct
type Logger struct {
	Level    int
	Date     bool
	Color    bool
	Function bool
	UTC      bool
}

// NewLogger creates a new logger
func NewLogger(name string) Logger {
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

	return Logger{
		Level:    logLevel,
		Date:     date,
		Color:    lcolor,
		Function: showFunc,
		UTC:      tzUTC,
	}
}

func (l Logger) setColor(m string, c int32) string {
	if l.Color {
		return fmt.Sprintf("\033[%dm%s\033[0m", c, m)
	}
	return m
}

func (l Logger) prefix(level string) *log.Logger {
	prefix := ""
	if l.Date {
		prefix += fmt.Sprintf("%v ", time.Now().UTC().Format("2006-01-02 15:04:05.000"))
	} else {
		prefix += fmt.Sprintf("%v ", time.Now().UTC().Format("15:04:05.000"))
	}
	switch level {
	case "DEBUG":
		prefix += l.setColor(level, 37)
	case "TRACE":
		prefix += l.setColor(level, 96)
	case "INFO":
		prefix += l.setColor(level, 94)
	case "WARN":
		prefix += l.setColor(level, 93)
	case "ERROR":
		prefix += l.setColor(level, 91)
	case "FATAL":
		prefix += l.setColor(level, 95)
	default:
		prefix += level
	}
	// need a space here
	prefix += " "

	if l.Function {
		_, file, line, _ := runtime.Caller(2)
		prefix += fmt.Sprintf("[%v:%v] ", filepath.Base(file), line)
	}
	ls := log.New(log.Writer(), prefix, 0)
	return ls
}

// Debug logs debug messages
func (l Logger) Debug(msg string) {
	if l.Level >= level["DEBUG"] {
		l.prefix("DEBUG").Printf(fmt.Sprintf(msg))
	}
}

// Debugf logs debug messages
func (l Logger) Debugf(format string, args ...interface{}) {
	if l.Level >= level["DEBUG"] {
		l.prefix("DEBUG").Printf(fmt.Sprintf(format, args...))
	}
}

// Trace logs trace messages
func (l Logger) Trace(msg string) {
	if l.Level >= level["TRACE"] {
		l.prefix("TRACE").Printf(msg)
	}
}

// Tracef logs trace messages
func (l Logger) Tracef(format string, args ...interface{}) {
	if l.Level >= level["TRACE"] {
		l.prefix("TRACE").Printf(fmt.Sprintf(format, args...))
	}
}

// Info logs info messages
func (l Logger) Info(msg string) {
	if l.Level >= level["INFO"] {
		l.prefix("INFO").Printf(msg)
	}
}

// Infof logs imfo messages
func (l Logger) Infof(format string, args ...interface{}) {
	if l.Level >= level["INFO"] {
		l.prefix("INFO").Printf(fmt.Sprintf(format, args...))
	}
}

// Warn logs warn messages
func (l Logger) Warn(msg string) {
	if l.Level >= level["WARN"] {
		l.prefix("WARN").Printf(msg)
	}
}

// Warnf logs wann messages
func (l Logger) Warnf(format string, args ...interface{}) {
	if l.Level >= level["WARN"] {
		l.prefix("WARN").Printf(fmt.Sprintf(format, args...))
	}
}

// Error logs error messages
func (l Logger) Error(msg string) {
	if l.Level >= level["ERROR"] {
		l.prefix("ERROR").Printf(msg)
	}
}

// Errorf logs error messages
func (l Logger) Errorf(format string, args ...interface{}) {
	if l.Level >= level["ERROR"] {
		l.prefix("ERROR").Printf(fmt.Sprintf(format, args...))
	}
}

// Panic logs fatal message and exits (1)
func (l Logger) Fatal(msg string) {
	if l.Level >= level["FATAL"] {
		l.prefix("FATAL").Fatalf(msg)
	}
}

// Panicf logs fatal message and exits (1)
func (l Logger) Fatalf(format string, args ...interface{}) {
	if l.Level >= level["FATAL"] {
		l.prefix("FATAL").Fatalf(fmt.Sprintf(format, args...))
	}
}
