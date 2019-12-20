package log

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

// I get the feeling you're trying to avoid using a third party lib in this library to make this a 'native' library but I would consider using assertions I use testify
// it simplifies the code, and on that note you can use the command `dep ensure` to pull in all dependencies with the Gopkg files
func TestNewLoggerLevelsTestify(t *testing.T) {
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
		{"panic", 0},
		{"PANIC", 0},
		{"Panic", 0},
	}

	for _, tt := range tests {
		actual := getLogger()
		actual.Level = tt.out
		expected := Logger{
			Level: tt.out,
		}
		// TODO: This test isn't super beneficial from the unit test perspective imo so I personally wouldn't include it
		assert.EqualValues(t, expected.Level, actual.Level, "the levels don't match")
	}
}

func TestNewLoggerUTCDateTestify(t *testing.T) {
	logger := getLogger()
	logger.Date = true
	logger.Color = false
	debugOutput := logger.Debug("unit test dis s***")
	reg := regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}\\s+\\d{2}:\\d{2}:\\d{2}") // this does leave off the nano seconds though
	assert.Regexp(t, reg, debugOutput, "date format is missing")
}

// this test is redundant
func TestNewLoggerDateTestify(t *testing.T) {
	logger := getLogger()
	logger.Date = true
	logger.Color = false
	logger.UTC = false
	debugOutput := logger.Debug("unit test dis s***")
	reg := regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}\\s+\\d{2}:\\d{2}:\\d{2}") // this does leave off the nano seconds though
	assert.Regexp(t, reg, debugOutput, "date format is missing")
}

func TestNewLoggerColor(t *testing.T) {
	logger := getLogger()
	debugOutput := logger.Debug("unit test dis s***")
	reg := regexp.MustCompile("\\x1b\\[90mDEBUG\\x1b\\[0m")
	assert.Regexp(t, reg, debugOutput, "date format is missing")
}

func TestNewLoggerFunction(t *testing.T) {
	logger := NewLogger(1, true, true, true, true)
	assert.IsType(t, Logger{}, logger, "NewLogger didn't send back a logger object")
	assert.EqualValues(t, 1, logger.Level)
	assert.EqualValues(t, true, logger.Color)
	assert.EqualValues(t, true, logger.Date)
	assert.EqualValues(t, true, logger.UTC)
	assert.EqualValues(t, true, logger.Function)
}

// Not a huge fan of doing A doesn't equal B but this is the easiest way to verify different outputs for timezone changes
func TestLoggerUTC(t *testing.T) {
	loggerUTC := getLogger()
	logger := getLogger()
	logger.UTC = false
	logUTCOut := loggerUTC.Debug("testing")
	logOut := logger.Debug("testing")
	assert.NotEqual(t, logUTCOut, logOut, "UTC produces the same timestamp")
}

func TestInfo(t *testing.T) {
	log := getLogger()
	s := log.Info("info")
	re := regexp.MustCompile(`info`)
	assert.Regexp(t, re, s)
}

func TestInfof(t *testing.T) {
	log := getLogger()
	s := log.Infof("test %v", "info")
	re := regexp.MustCompile(`test info`)
	assert.Regexp(t, re, s)
}

func TestFormat(t *testing.T) {
	logger := getLogger()
	result := logger.format(infoLogLevel, "testing")
	regex := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} INFO .+ testing`)
	assert.Regexp(t, regex, result, "format didn't match")
}

func TestLevelOutput(t *testing.T) {
	log := getLogger()
	testMsg := "testing"
	log.Level = 4
	infoResult := log.Info(testMsg)
	traceResult := log.Trace(testMsg)
	warnResult := log.Warn(testMsg)
	// I don't like multi-assert tests but i'm getting lazy
	assert.NotEqual(t, "", infoResult, "same level error")
	assert.NotEqual(t, "", warnResult, "lower level error")
	assert.EqualValues(t, "", traceResult, "height level error")
}

func TestFormatLevelOutput(t *testing.T) {
	log := getLogger()
	testMsg := "testing"
	testFormat := "%s"
	log.Level = 4
	infoResult := log.Infof(testFormat, testMsg)
	traceResult := log.Tracef(testFormat, testMsg)
	warnResult := log.Warnf(testFormat, testMsg)
	// I don't like multi-assert tests but i'm getting lazy
	assert.NotEqual(t, "", infoResult, "same level error")
	assert.NotEqual(t, "", warnResult, "lower level error")
	assert.EqualValues(t, "", traceResult, "height level error")
}

// I personally would make a benchmark test file but that's just personal preference
func BenchmarkInfoWrite(b *testing.B) {
	l := getLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Debug("Debug Message")
	}
}

func BenchmarkNewLogger(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getLogger()
	}
}

func getLogger() Logger {
	return NewLogger(debugLevel, true, true, true, true)
}
