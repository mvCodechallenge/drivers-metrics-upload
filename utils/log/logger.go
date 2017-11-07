/*
	Simple logger that enables levels
*/

package log

import (
	"time"
	"fmt"
	"CodeChallenge/DriversMetricsImporter/utils/config"
	"strings"
	"os"
)

/*
	Holds log level name and number
*/
type logLevel struct {
	name  string
	value byte
}

const log_format string = "[%s]\t%s\t%s\n"

/*
	Holds all log levels
*/
var (
	trace_level logLevel = logLevel{"TRACE", 1}
	info_level logLevel = logLevel{"INFO", 2}
	warn_level logLevel = logLevel{"WARN", 3}
	error_level logLevel = logLevel{"ERROR", 4}
	fatal_level logLevel = logLevel{"FATAL", 5}
	none_level logLevel = logLevel{"NONE", 100}
)

var (
	// Current log level, defaulted to INFO level
	currentLogLevel byte = info_level.value
	// Holds map of log levels to fast retrieval
	logLevels map[string]logLevel;
	// Holds app config key-value data
	appConfig map[string]interface{};
)

/*
	Populates all log levels to map
*/
func populateLogLevels() {
	logLevels = map[string]logLevel{};
	logLevels[none_level.name] = none_level;
	logLevels[trace_level.name] = trace_level;
	logLevels[info_level.name] = info_level;
	logLevels[warn_level.name] = warn_level;
	logLevels[fatal_level.name] = fatal_level;
}

/*
	Initialize all log levels types, map of log levels, current log level and app config
*/
func init() {
	populateLogLevels();

	appConfig = config.GetAppConfiguration();
	requiredLogLevelName := appConfig["logLevel"].(string);
	if (len(requiredLogLevelName) < 1) {
		return;
	}

	currentLogLevel = logLevels[strings.ToUpper(requiredLogLevelName)].value
}

/*
	Gets formatted time for log
*/
func getTime() string {
	return time.Now().Format(time.RFC3339)
}

/*
	Prepend log message to include level and time
*/
func adjustLog(level string, msg string) string {
	return fmt.Sprintf(log_format, level, getTime(), msg)
}

/*
	Trace level log sending
*/
func Trace(msg string) {
	if (currentLogLevel > trace_level.value) {
		return;
	}

	fmt.Fprintf(os.Stdout, adjustLog(trace_level.name, msg))
}

/*
	Error level log sending
*/
func Error(msg string) {
	if (currentLogLevel > error_level.value) {
		return;
	}

	fmt.Fprintf(os.Stderr, adjustLog(error_level.name, msg))
}

/*
	Info level log sending
*/
func Info(msg string) {
	if (currentLogLevel > info_level.value) {
		return;
	}

	fmt.Fprintf(os.Stdout, adjustLog(info_level.name, msg))
}

/*
	Warn level log sending
*/
func Warn(msg string) {
	if (currentLogLevel > warn_level.value) {
		return;
	}

	fmt.Println(adjustLog(warn_level.name, msg))
}

/*
	Fatal level log sending
*/
func Fatal(msg string) {
	if (currentLogLevel > fatal_level.value) {
		return;
	}

	fmt.Fprintf(os.Stderr, adjustLog(fatal_level.name, msg))
}


