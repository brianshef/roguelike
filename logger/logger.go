package logger

import (
	"io"
	"log"
	"os"
)

const logOpts = log.Ldate | log.Ltime | log.Lshortfile

// Log defines the set of loggers available
type Log struct {
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func (l *Log) init(
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) *Log {

	l.Debug = log.New(debugHandle, "[DEBUG] ", logOpts)
	l.Info = log.New(infoHandle, "[INFO] ", logOpts)
	l.Warning = log.New(warningHandle, "[WARN] ", logOpts)
	l.Error = log.New(errorHandle, "[ERROR] ", logOpts)

	return l
}

// NewLoggers is a factory function which constructions a new Log
func NewLoggers() *Log {
	loggers := new(Log).init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	return loggers
}
