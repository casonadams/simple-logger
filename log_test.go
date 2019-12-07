package log_test

import (
	"os"
	"testing"

	logger "github.com/casonadams/simple-logger/v2"
)

func TestNewLogger(t *testing.T) {
	actual := logger.NewLogger("test")
	expected := logger.Logger{
		Level:    4,
		Date:     true,
		Color:    true,
		Function: true,
		UTC:      true,
	}
	if expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestNewLoggerLevels(t *testing.T) {
	var tests = []struct {
		in  string
		out int
	}{
		{"debug", 6},
		{"DEBUG", 6},
		{"Debug", 6},
		{"trace", 5},
		{"TRACE", 5},
		{"Trace", 5},
		{"info", 4},
		{"INFO", 4},
		{"Info", 4},
		{"warn", 3},
		{"WARN", 3},
		{"Warn", 3},
		{"error", 2},
		{"ERROR", 2},
		{"Error", 2},
		{"fatal", 1},
		{"FATAL", 1},
		{"Fatal", 1},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.in)
		actual := logger.NewLogger("test")
		expected := logger.Logger{
			Level: tt.out,
		}
		if expected.Level != actual.Level {
			t.Errorf("Test(%v): expected %v, actual %v", i, expected.Level, actual.Level)
		}

	}
}

func TestNewLoggerDate(t *testing.T) {
	var tests = []struct {
		in  string
		out bool
	}{
		{"1", true},
		{"0", false},
		{"TRUE", true},
		{"FALSE", false},
		{"true", true},
		{"false", false},
		{"True", true},
		{"False", false},
	}

	for i, tt := range tests {
		os.Setenv("LOG_DATE", tt.in)
		actual := logger.NewLogger("test")
		expected := logger.Logger{
			Date: tt.out,
		}
		if expected.Date != actual.Date {
			t.Errorf("Test(%v): expected %v, actual %v", i, expected.Date, actual.Date)
		}
	}
}

func TestNewLoggerColor(t *testing.T) {
	var tests = []struct {
		in  string
		out bool
	}{
		{"1", true},
		{"0", false},
		{"TRUE", true},
		{"FALSE", false},
		{"true", true},
		{"false", false},
		{"True", true},
		{"False", false},
	}

	for i, tt := range tests {
		os.Setenv("LOG_COLOR", tt.in)
		actual := logger.NewLogger("test")
		expected := logger.Logger{
			Color: tt.out,
		}
		if expected.Color != actual.Color {
			t.Errorf("Test(%v): expected %v, actual %v", i, expected.Color, actual.Color)
		}
	}
}

func TestNewLoggerFunction(t *testing.T) {
	var tests = []struct {
		in  string
		out bool
	}{
		{"1", true},
		{"0", false},
		{"TRUE", true},
		{"FALSE", false},
		{"true", true},
		{"false", false},
		{"True", true},
		{"False", false},
	}

	for i, tt := range tests {
		os.Setenv("LOG_FUNC", tt.in)
		actual := logger.NewLogger("test")
		expected := logger.Logger{
			Function: tt.out,
		}
		if expected.Function != actual.Function {
			t.Errorf("Test(%v): expected %v, actual %v", i, expected.Function, actual.Function)
		}
	}
}

func TestNewLoggerUTC(t *testing.T) {
	var tests = []struct {
		in  string
		out bool
	}{
		{"1", true},
		{"0", false},
		{"TRUE", true},
		{"FALSE", false},
		{"true", true},
		{"false", false},
		{"True", true},
		{"False", false},
	}

	for i, tt := range tests {
		os.Setenv("LOG_UTC", tt.in)
		actual := logger.NewLogger("test")
		expected := logger.Logger{
			UTC: tt.out,
		}
		if expected.UTC != actual.UTC {
			t.Errorf("Test(%v): expected %v, actual %v", i, expected.UTC, actual.UTC)
		}
	}
}
