package mlog

import (
	"io"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

//Items type
type Items map[string]interface{}

var (
	appName       string
	version       string
	containerName string
	//Setting default log level as INFO
	logLevel        = INFO
	printStackTrace = true
)

//Logging levels
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

type Config struct {
	ContainerName   string
	Level           string
	Format          string
	PrintStackTrace bool
}

func Initialize(config Config) {
	initializeContainerName(config)
	initializeLoggingLevel(config)
	initializeLoggingFormat(config)
}

func initializeContainerName(config Config) {
	if config.ContainerName != "" {
		containerName = config.ContainerName
	} else {
		containerName = os.Getenv("LOGGER_NAME")
	}
}

func initializeLoggingLevel(config Config) {
	if config.Level == "" {
		return
	}
	switch config.Level {
	case "DEBUG":
		logLevel = DEBUG
		logrus.SetLevel(logrus.DebugLevel)
	case "INFO":
		logLevel = INFO
		logrus.SetLevel(logrus.InfoLevel)
	case "WARN":
		logLevel = WARN
		logrus.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logLevel = ERROR
		logrus.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		logLevel = FATAL
		logrus.SetLevel(logrus.FatalLevel)
	case "PANIC":
		logLevel = PANIC
		logrus.SetLevel(logrus.PanicLevel)
	}
	printStackTrace = config.PrintStackTrace
}

func initializeLoggingFormat(config Config) {
	if config.Format == "" {
		return
	}
	switch config.Format {
	case "JSON":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "TEXT":
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}

func getContext() *logrus.Entry {
	pc, fileName, line, _ := runtime.Caller(3)
	funcName := runtime.FuncForPC(pc).Name()
	return logrus.WithFields(logrus.Fields{
		"appName":   appName,
		"version":   version,
		"container": containerName,
		"file":      fileName,
		"func":      funcName,
		"line":      line,
	})
}

func prepareContext(items []Items) *logrus.Entry {
	contextLogger := getContext()
	for _, item := range items {
		for key, val := range item {
			contextLogger = contextLogger.WithField(key, val)
		}
	}
	return contextLogger
}

func Debug(str string, items ...Items) {
	if logLevel > DEBUG {
		return
	}
	contextLogger := prepareContext(items)
	contextLogger.Debug(str)
}

func Info(str string, items ...Items) {
	if logLevel > INFO {
		return
	}
	contextLogger := prepareContext(items)
	contextLogger.Info(str)
}

func Warn(str string, items ...Items) {
	if logLevel > WARN {
		return
	}
	contextLogger := prepareContext(items)
	contextLogger.Warn(str)
}

func Error(str string, items ...Items) {
	if logLevel > ERROR {
		return
	}
	contextLogger := prepareContext(items)
	contextLogger.Error(str)
	if printStackTrace {
		debug.PrintStack()
	}
}

func Fatal(str string, items ...Items) {
	if logLevel > FATAL {
		return
	}
	contextLogger := prepareContext(items)
	contextLogger.Fatal(str)
}

func Panic(str string, items ...Items) {
	if logLevel > PANIC {
		return
	}
	contextLogger := prepareContext(items)
	contextLogger.Panic(str)
}

// SetOutput sets the Logrus standard output.
func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

// SetFormatter sets the Logrus standard formatter.
func SetFormatter(formatter logrus.Formatter) {
	logrus.SetFormatter(formatter)
}

func SetPrintStackTrace(to bool) {
	printStackTrace = to
}

// GetLevel retrieves the current log level.
func GetLevel() int {
	return logLevel
}

//SetLevel set the current log level
func SetLevel(in int) {
	logLevel = in
}
