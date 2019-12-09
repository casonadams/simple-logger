package log_test

import (
	"os"
	"regexp"
	"testing"

	logger "github.com/casonadams/simple-logger"
)

func TestNewLogger(t *testing.T) {
	actual := logger.NewLogger("test")
	expected := &logger.Logger{
		Level:    4,
		Date:     true,
		Color:    true,
		Function: true,
		UTC:      true,
	}

	if expected.Level != actual.Level {
		t.Errorf("expected %v, actual %v", expected.Level, actual.Level)
	}
	if expected.Date != actual.Date {
		t.Errorf("expected %v, actual %v", expected.Date, actual.Date)
	}
	if expected.Color != actual.Color {
		t.Errorf("expected %v, actual %v", expected.Color, actual.Color)
	}
	if expected.Function != actual.Function {
		t.Errorf("expected %v, actual %v", expected.Function, actual.Function)
	}
	if expected.UTC != actual.UTC {
		t.Errorf("expected %v, actual %v", expected.UTC, actual.UTC)
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
		{"panic", 0},
		{"PANIC", 0},
		{"Panic", 0},
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

func TestInfo(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	log := logger.NewLogger("test")
	s := log.Info("info")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+94mINFO\D+0m \W\w+.\w+:\d+\W info`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		t.Errorf("Test expected: %v actual: %v", match, s)
	}
	for _, v := range match {
		if v != s {
			t.Errorf("Test expected: %v, actual %v", v, s)
		}
	}
}

func TestInfof(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	log := logger.NewLogger("test")
	s := log.Infof("%v", "info")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+94mINFO\D+0m \W\w+.\w+:\d+\W info`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		t.Errorf("Test expected: %v actual: %v", match, s)
	}
	for _, v := range match {
		if v != s {
			t.Errorf("Test expected: %v, actual %v", v, s)
		}
	}
}

func TestDebug(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	log := logger.NewLogger("test")
	s := log.Debug("info")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+90mDEBUG\D+0m \W\w+.\w+:\d+\W info`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		t.Errorf("Test expected: %v actual: %v", match, s)
	}
	for _, v := range match {
		if v != s {
			t.Errorf("Test expected: %v, actual %v", v, s)
		}
	}
}

func TestDebugf(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	log := logger.NewLogger("test")
	s := log.Debugf("%v", "info")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+90mDEBUG\D+0m \W\w+.\w+:\d+\W info`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		t.Errorf("Test expected: %v actual: %v", match, s)
	}
	for _, v := range match {
		if v != s {
			t.Errorf("Test expected: %v, actual %v", v, s)
		}
	}
}

func TestTrace(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	log := logger.NewLogger("test")
	s := log.Trace("info")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+96mTRACE\D+0m \W\w+.\w+:\d+\W info`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		t.Errorf("Test expected: %v actual: %v", match, s)
	}
	for _, v := range match {
		if v != s {
			t.Errorf("Test expected: %v, actual %v", v, s)
		}
	}
}

func TestTracef(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	log := logger.NewLogger("test")
	s := log.Tracef("%v", "info")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+96mTRACE\D+0m \W\w+.\w+:\d+\W info`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		t.Errorf("Test expected: %v actual: %v", match, s)
	}
	for _, v := range match {
		if v != s {
			t.Errorf("Test expected: %v, actual %v", v, s)
		}
	}
}

func TestWarn(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	log := logger.NewLogger("test")
	s := log.Warn("info")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+93mWARN\D+0m \W\w+.\w+:\d+\W info`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		t.Errorf("Test expected: %v actual: %v", match, s)
	}
	for _, v := range match {
		if v != s {
			t.Errorf("Test expected: %v, actual %v", v, s)
		}
	}
}

func TestWarnf(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	log := logger.NewLogger("test")
	s := log.Warnf("%v", "info")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+93mWARN\D+0m \W\w+.\w+:\d+\W info`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		t.Errorf("Test expected: %v actual: %v", match, s)
	}
	for _, v := range match {
		if v != s {
			t.Errorf("Test expected: %v, actual %v", v, s)
		}
	}
}

func TestError(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	log := logger.NewLogger("test")
	s := log.Error("info")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+91mERROR\D+0m \W\w+.\w+:\d+\W info`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		t.Errorf("Test expected: %v actual: %v", match, s)
	}
	for _, v := range match {
		if v != s {
			t.Errorf("Test expected: %v, actual %v", v, s)
		}
	}
}

func TestErrorf(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	log := logger.NewLogger("test")
	s := log.Errorf("%v", "info")
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+91mERROR\D+0m \W\w+.\w+:\d+\W info`)
	match := re.FindStringSubmatch(s)
	if len(match) == 0 {
		t.Errorf("Test expected: %v actual: %v", match, s)
	}
	for _, v := range match {
		if v != s {
			t.Errorf("Test expected: %v, actual %v", v, s)
		}
	}
}

func TestFormat(t *testing.T) {
	var tests = []struct {
		cEnv  string
		fEnv  string
		dEnv  string
		lEnv  string
		regex string
	}{
		{"false", "false", "false", "INFO", `\d{2}:\d{2}:\d{2}.\d{3} INFO info`},
		{"false", "true", "false", "INFO", `\d{2}:\d{2}:\d{2}.\d{3} INFO \W\w+.\w+:\d+\W info`},
		{"false", "false", "true", "INFO", `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} INFO info`},
		{"false", "true", "true", "INFO", `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} INFO \W\w+.\w+:\d+\W info`},
		{"true", "false", "false", "INFO", `\d{2}:\d{2}:\d{2}.\d{3} \D+94mINFO\D+0m info`},
		{"true", "false", "true", "INFO", `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+94mINFO\D+0m info`},
		{"true", "true", "false", "INFO", `\d{2}:\d{2}:\d{2}.\d{3} \D+94mINFO\D+0m \W\w+.\w+:\d+\W info`},
		{"true", "true", "true", "INFO", `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3} \D+94mINFO\D+0m \W\w+.\w+:\d+\W info`},
	}

	for i, tt := range tests {
		os.Setenv("LOG_COLOR", tt.cEnv)
		os.Setenv("LOG_FUNC", tt.fEnv)
		os.Setenv("LOG_DATE", tt.dEnv)
		os.Setenv("LOG_LEVEL", tt.lEnv)

		log := logger.NewLogger("test")
		s := log.Infof("%v", "info")
		re := regexp.MustCompile(tt.regex)
		match := re.FindStringSubmatch(s)
		if len(match) == 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
		}
		for _, v := range match {
			if v != s {
				t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
			}
		}
	}
}

func TestDebugLevelOutput(t *testing.T) {
	var tests = []struct {
		logLevel string
	}{
		{"DEBUG"},
		{"TRACE"},
		{"INFO"},
		{"WARN"},
		{"ERROR"},
		{"FATAL"},
		{"PANIC"},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.logLevel)
		os.Setenv("LOG_COLOR", "false")
		os.Setenv("LOG_FUNC", "false")
		os.Setenv("LOG_DATE", "false")

		log := logger.NewLogger("test")
		s := log.Debug("info")
		if tt.logLevel == "DEBUG" {
			re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} DEBUG info`)
			match := re.FindStringSubmatch(s)
			if len(match) == 0 {
				t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
			}
			for _, v := range match {
				if v != s {
					t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
				}
			}
		} else if len(s) != 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, s, "")
		}
	}
}

func TestDebugfLevelOutput(t *testing.T) {
	var tests = []struct {
		logLevel string
	}{
		{"DEBUG"},
		{"TRACE"},
		{"INFO"},
		{"WARN"},
		{"ERROR"},
		{"FATAL"},
		{"PANIC"},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.logLevel)
		os.Setenv("LOG_COLOR", "false")
		os.Setenv("LOG_FUNC", "false")
		os.Setenv("LOG_DATE", "false")

		log := logger.NewLogger("test")
		s := log.Debugf("%v", "info")
		if tt.logLevel == "DEBUG" {
			re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} DEBUG info`)
			match := re.FindStringSubmatch(s)
			if len(match) == 0 {
				t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
			}
			for _, v := range match {
				if v != s {
					t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
				}
			}
		} else if len(s) != 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, s, "")
		}
	}
}

func TestTraceLevelOutput(t *testing.T) {
	var tests = []struct {
		logLevel string
	}{
		{"DEBUG"},
		{"TRACE"},
		{"INFO"},
		{"WARN"},
		{"ERROR"},
		{"FATAL"},
		{"PANIC"},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.logLevel)
		os.Setenv("LOG_COLOR", "false")
		os.Setenv("LOG_FUNC", "false")
		os.Setenv("LOG_DATE", "false")

		log := logger.NewLogger("test")
		s := log.Trace("info")

		if tt.logLevel == "DEBUG" || tt.logLevel == "TRACE" {
			re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} TRACE info`)
			match := re.FindStringSubmatch(s)
			if len(match) == 0 {
				t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
			}
			for _, v := range match {
				if v != s {
					t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
				}
			}
		} else if len(s) != 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, s, "")
		}
	}
}

func TestTracefLevelOutput(t *testing.T) {
	var tests = []struct {
		logLevel string
	}{
		{"DEBUG"},
		{"TRACE"},
		{"INFO"},
		{"WARN"},
		{"ERROR"},
		{"FATAL"},
		{"PANIC"},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.logLevel)
		os.Setenv("LOG_COLOR", "false")
		os.Setenv("LOG_FUNC", "false")
		os.Setenv("LOG_DATE", "false")

		log := logger.NewLogger("test")
		s := log.Tracef("%v", "info")

		if tt.logLevel == "DEBUG" || tt.logLevel == "TRACE" {
			re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} TRACE info`)
			match := re.FindStringSubmatch(s)
			if len(match) == 0 {
				t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
			}
			for _, v := range match {
				if v != s {
					t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
				}
			}
		} else if len(s) != 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, s, "")
		}
	}
}

func TestInfoLevelOutput(t *testing.T) {
	var tests = []struct {
		logLevel string
	}{
		{"DEBUG"},
		{"TRACE"},
		{"INFO"},
		{"WARN"},
		{"ERROR"},
		{"FATAL"},
		{"PANIC"},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.logLevel)
		os.Setenv("LOG_COLOR", "false")
		os.Setenv("LOG_FUNC", "false")
		os.Setenv("LOG_DATE", "false")

		log := logger.NewLogger("test")
		s := log.Info("info")

		if tt.logLevel == "DEBUG" || tt.logLevel == "TRACE" || tt.logLevel == "INFO" {
			re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} INFO info`)
			match := re.FindStringSubmatch(s)
			if len(match) == 0 {
				t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
			}
			for _, v := range match {
				if v != s {
					t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
				}
			}
		} else if len(s) != 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, s, "")
		}
	}
}

func TestInfofLevelOutput(t *testing.T) {
	var tests = []struct {
		logLevel string
	}{
		{"DEBUG"},
		{"TRACE"},
		{"INFO"},
		{"WARN"},
		{"ERROR"},
		{"FATAL"},
		{"PANIC"},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.logLevel)
		os.Setenv("LOG_COLOR", "false")
		os.Setenv("LOG_FUNC", "false")
		os.Setenv("LOG_DATE", "false")

		log := logger.NewLogger("test")
		s := log.Infof("%v", "info")

		if tt.logLevel == "DEBUG" || tt.logLevel == "TRACE" || tt.logLevel == "INFO" {
			re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} INFO info`)
			match := re.FindStringSubmatch(s)
			if len(match) == 0 {
				t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
			}
			for _, v := range match {
				if v != s {
					t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
				}
			}
		} else if len(s) != 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, s, "")
		}
	}
}

func TestWarnfLevelOutput(t *testing.T) {
	var tests = []struct {
		logLevel string
	}{
		{"DEBUG"},
		{"TRACE"},
		{"INFO"},
		{"WARN"},
		{"ERROR"},
		{"FATAL"},
		{"PANIC"},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.logLevel)
		os.Setenv("LOG_COLOR", "false")
		os.Setenv("LOG_FUNC", "false")
		os.Setenv("LOG_DATE", "false")

		log := logger.NewLogger("test")
		s := log.Warnf("%v", "info")

		if tt.logLevel == "DEBUG" || tt.logLevel == "TRACE" || tt.logLevel == "INFO" || tt.logLevel == "WARN" {
			re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} WARN info`)
			match := re.FindStringSubmatch(s)
			if len(match) == 0 {
				t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
			}
			for _, v := range match {
				if v != s {
					t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
				}
			}
		} else if len(s) != 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, s, "")
		}
	}
}

func TestWarnLevelOutput(t *testing.T) {
	var tests = []struct {
		logLevel string
	}{
		{"DEBUG"},
		{"TRACE"},
		{"INFO"},
		{"WARN"},
		{"ERROR"},
		{"FATAL"},
		{"PANIC"},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.logLevel)
		os.Setenv("LOG_COLOR", "false")
		os.Setenv("LOG_FUNC", "false")
		os.Setenv("LOG_DATE", "false")

		log := logger.NewLogger("test")
		s := log.Warn("info")

		if tt.logLevel == "DEBUG" || tt.logLevel == "TRACE" || tt.logLevel == "INFO" || tt.logLevel == "WARN" {
			re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} WARN info`)
			match := re.FindStringSubmatch(s)
			if len(match) == 0 {
				t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
			}
			for _, v := range match {
				if v != s {
					t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
				}
			}
		} else if len(s) != 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, s, "")
		}
	}
}

func TestErrorfLevelOutput(t *testing.T) {
	var tests = []struct {
		logLevel string
	}{
		{"DEBUG"},
		{"TRACE"},
		{"INFO"},
		{"WARN"},
		{"ERROR"},
		{"FATAL"},
		{"PANIC"},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.logLevel)
		os.Setenv("LOG_COLOR", "false")
		os.Setenv("LOG_FUNC", "false")
		os.Setenv("LOG_DATE", "false")

		log := logger.NewLogger("test")
		s := log.Errorf("%v", "info")

		if tt.logLevel == "DEBUG" || tt.logLevel == "TRACE" || tt.logLevel == "INFO" || tt.logLevel == "WARN" || tt.logLevel == "ERROR" {
			re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} ERROR info`)
			match := re.FindStringSubmatch(s)
			if len(match) == 0 {
				t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
			}
			for _, v := range match {
				if v != s {
					t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
				}
			}
		} else if len(s) != 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, s, "")
		}
	}
}

func TestErrorLevelOutput(t *testing.T) {
	var tests = []struct {
		logLevel string
	}{
		{"DEBUG"},
		{"TRACE"},
		{"INFO"},
		{"WARN"},
		{"ERROR"},
		{"FATAL"},
		{"PANIC"},
	}

	for i, tt := range tests {
		os.Setenv("LOG_LEVEL", tt.logLevel)
		os.Setenv("LOG_COLOR", "false")
		os.Setenv("LOG_FUNC", "false")
		os.Setenv("LOG_DATE", "false")

		log := logger.NewLogger("test")
		s := log.Error("info")

		if tt.logLevel == "DEBUG" || tt.logLevel == "TRACE" || tt.logLevel == "INFO" || tt.logLevel == "WARN" || tt.logLevel == "ERROR" {
			re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3} ERROR info`)
			match := re.FindStringSubmatch(s)
			if len(match) == 0 {
				t.Errorf("Test(%d) expected: %v actual: %v", i, match, s)
			}
			for _, v := range match {
				if v != s {
					t.Errorf("Test(%d) expected: %v, actual %v", i, v, s)
				}
			}
		} else if len(s) != 0 {
			t.Errorf("Test(%d) expected: %v actual: %v", i, s, "")
		}
	}
}

func BenchmarkInfoWrite(b *testing.B) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_FUNC", "true")
	os.Setenv("LOG_DATE", "true")
	l := logger.NewLogger("test")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Debug("Debug Message")
	}
}

func BenchmarkNewLogger(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.NewLogger("test")
	}
}
