package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLoger  *log.Logger

	flags = log.Lmsgprefix | log.Ldate | log.Lmicroseconds | log.Lshortfile
)

func init() {
	output := os.Stdout
	debugLogger = log.New(output, "DEBUG ", flags)
	infoLogger = log.New(output, "INFO ", flags)
	warnLogger = log.New(output, "WARN ", flags)
	errorLogger = log.New(output, "ERROR ", flags)
	fatalLoger = log.New(output, "FATAL ", flags)
}

// returns file handler for writing log
func FileOutput(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ErrorCheck(err, true, nil)

	return f
}

// returns os. Stdout
func StdOutput() io.Writer {
	return os.Stdout
}

// set log output
func SetOutput(writers ...io.Writer) {
	output := io.MultiWriter(writers...)
	infoLogger.SetOutput(output)
	warnLogger.SetOutput(output)
	errorLogger.SetOutput(output)
	fatalLoger.SetOutput(output)
}

func Debug(v ...interface{}) {
	debugLogger.Output(2, fmt.Sprintln(v...))
}

func Info(v ...interface{}) {
	infoLogger.Output(2, fmt.Sprintln(v...))
}

func Warn(v ...interface{}) {
	warnLogger.Output(2, fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	errorLogger.Output(2, fmt.Sprintln(v...))
}

func Fatal(v ...interface{}) {
	fatalLoger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func ErrorCheck(err error, fatal bool, cleanupFn func()) bool {
	if err == nil {
		return false
	}

	if fatal {
		fatalLoger.Output(2, fmt.Sprintln(err))
	} else {
		errorLogger.Output(2, fmt.Sprintln(err))
	}

	if cleanupFn != nil {
		cleanupFn()
	}

	if fatal {
		os.Exit(1)
	}

	return true
}
