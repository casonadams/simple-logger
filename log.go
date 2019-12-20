package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

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

	dateTimeFormat = "2006-01-02 15:04:05.000"
	timeFormat     = "15:04:05.000"

	debugLogLevel = "DEBUG"
	traceLogLevel = "TRACE"
	infoLogLevel  = "INFO"
	warnLogLevel  = "WARN"
	errorLogLevel = "ERROR"
	fatalLogLevel = "FATAL"
	debugLevel    = 6
	traceLevel    = 5
	infoLevel     = 4
	warnLevel     = 3
	errorLevel    = 2
)

//Logger struct
type Logger struct {
	// if you were to add log to file support I would add the mutex back in but if you are just going to do stdout then it isn't needed
	Level    int
	Date     bool // nit-pick this is a confusing name I personally would like to see something more along the lines of useDate, also this var doesn't need to be exported
	Color    bool // useColor, doesn't need to be exported
	Function bool // useFileCaller, *
	UTC      bool // useUTC, *
	// var names can be changed but I like the name to hint what type the var is(if possible) use/is tends to be related to booleans(in my mind at least)
	// useDate bool
	// useColor bool
	// useFileCaller bool
	// useUTC bool
}

// NeLogger creates a new logger
func NewLogger(logLevel int, logDate, callerLocation, useUTC, color bool) Logger {
	// if this is a library that will be used in multiple services/applications then env vars shouldn't be used at this
	// level they should be passed in from the application using it.
	return Logger{
		Level:    logLevel,
		Date:     logDate,
		Color:    color,
		Function: callerLocation,
		UTC:      useUTC,
	}
}

func (logger Logger) color(m string, c color) string {
	if logger.Color {
		return fmt.Sprintf("\033[%dm%s\033[0m", c, m)
	}
	return m
}

// personal preference here but l is a rough var name imo I prefer to be explicit even if it seems a bit redundant
// logger.Date makes more sense than logger.Date
func (logger Logger) format(logLevel string, msg string) string {
	// NOTE: I'm all about pointer receivers but if it isn't mutating the base struct it can be avoided
	var currentTime string
	var callerData string

	// Setup timestamp
	if logger.UTC {
		if logger.Date {
			currentTime = time.Now().UTC().Format(dateTimeFormat)
		} else {
			currentTime = time.Now().UTC().Format(timeFormat)
		}
	} else {
		if logger.Date {
			currentTime = time.Now().Format(dateTimeFormat)
		} else {
			currentTime = time.Now().Format(timeFormat)
		}
	}

	// Caller location
	if logger.Function {
		_, file, line, _ := runtime.Caller(2)
		callerData = fmt.Sprintf("[%v:%v]", filepath.Base(file), line)
	}

	//avoid string concatenation if possible the overhead of the additional vars is negligible so readability wins
	return fmt.Sprintf("%s %s %s %s", currentTime, logLevel, callerData, msg)
}

// Debug logs debug messages
func (logger Logger) Debug(msg string) string {
	if logger.Level >= debugLevel {
		s := logger.format(logger.color(debugLogLevel, GRAY), msg)
		fmt.Println(s)
		return s
	}
	return ""
}

// Debugf logs debug messages
func (logger Logger) Debugf(format string, args ...interface{}) string {
	if logger.Level >= debugLevel {
		s := logger.format(logger.color(debugLogLevel, GRAY), fmt.Sprintf(format, args...))
		fmt.Println(s)
		return s
	}
	return ""
}

// Trace logs trace messages
func (logger Logger) Trace(msg string) string {
	if logger.Level >= traceLevel {
		s := logger.format(logger.color(traceLogLevel, CYAN), msg)
		fmt.Println(s)
		return s
	}
	return ""
}

// Tracef logs trace messages
func (logger Logger) Tracef(format string, args ...interface{}) string {
	if logger.Level >= traceLevel {
		s := logger.format(logger.color(traceLogLevel, CYAN), fmt.Sprintf(format, args...))
		fmt.Println(s)
		return s
	}
	return ""
}

// Info logs info messages
func (logger Logger) Info(msg string) string {
	if logger.Level >= infoLevel {
		s := logger.format(logger.color(infoLogLevel, BLUE), msg)
		fmt.Println(s)
		return s
	}
	return ""
}

// Infof logs imfo messages
func (logger Logger) Infof(format string, args ...interface{}) string {
	if logger.Level >= infoLevel {
		s := logger.format(logger.color(infoLogLevel, BLUE), fmt.Sprintf(format, args...))
		fmt.Println(s)
		return s
	}
	return ""
}

// Warn logs warn messages
func (logger Logger) Warn(msg string) string {
	if logger.Level >= warnLevel {
		s := logger.format(logger.color(warnLogLevel, YELLOW), msg)
		fmt.Println(s)
		return s
	}
	return ""
}

// Warnf logs wann messages
func (logger Logger) Warnf(format string, args ...interface{}) string {
	if logger.Level >= warnLevel {
		s := logger.format(logger.color(warnLogLevel, YELLOW), fmt.Sprintf(format, args...))
		fmt.Println(s)
		return s
	}
	return ""
}

// Error logs error messages
func (logger Logger) Error(msg string) string {
	if logger.Level >= errorLevel {
		s := logger.format(logger.color(errorLogLevel, RED), msg)
		fmt.Println(s)
		return s
	}
	return ""
}

// Errorf logs error messages
func (logger Logger) Errorf(format string, args ...interface{}) string {
	if logger.Level >= errorLevel {
		s := logger.format(logger.color(errorLogLevel, RED), fmt.Sprintf(format, args...))
		fmt.Println(s)
		return s
	}
	return ""
}

// Fatal logs fatal message and exits (1)
func (logger Logger) Fatal(msg string) string {
	s := logger.format(logger.color(fatalLogLevel, MAGENTA), msg)
	fmt.Println(s)
	defer os.Exit(1) // do we want a configurable code???
	return s
}

//func (logger Logger) Fatal(msg string, code int) string {
//	s := logger.format(logger.color("FATAL", MAGENTA), msg)
//	fmt.Println(s)
//	defer os.Exit(code)
//	return s
//}

// Fatalf logs fatal message and exits (1)
func (logger Logger) Fatalf(format string, args ...interface{}) string {
	s := logger.format(logger.color(fatalLogLevel, MAGENTA), fmt.Sprintf(format, args...))
	fmt.Println(s)
	defer os.Exit(1)
	return s
}

// When you use panic it will log the core behavior
// Panic logs fatal message and exits (1)
func (logger Logger) Panic(msg string) {
	s := logger.format(logger.color("PANIC", DMAGENTA), msg)
	fmt.Println(s)
	defer os.Exit(1)
}

// Panicf logs fatal message and exits (1)
func (logger Logger) Panicf(format string, args ...interface{}) {
	s := logger.format(logger.color("PANIC", DMAGENTA), fmt.Sprintf(format, args...))
	fmt.Println(s)
	defer os.Exit(1)
}
