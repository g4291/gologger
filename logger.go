// TODO
// multiwriter: use os.stdout and io.writer

package main

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLoger  *log.Logger

	flags = log.Lmsgprefix | log.Ldate | log.Lmicroseconds | log.Lshortfile
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO ", flags)
	warnLogger = log.New(os.Stdout, "WARN ", flags)
	errorLogger = log.New(os.Stdout, "ERROR ", flags)
	fatalLoger = log.New(os.Stdout, "FATAL ", flags)
}

func Info(v ...interface{}) {
	v = append(v, "\n")
	infoLogger.Output(2, fmt.Sprintln(v...))
}

func Warn(v ...interface{}) {
	v = append(v, "\n")
	warnLogger.Output(2, fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	v = append(v, "\n")
	errorLogger.Output(2, fmt.Sprintln(v...))
}

func Fatal(v ...interface{}) {
	v = append(v, "\n")
	fatalLoger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}
