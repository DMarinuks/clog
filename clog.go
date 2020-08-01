package clog

import (
	"io"
	"log"
	"os"
)

const (
	redColor       = "\033[1;31m"
	greenColor     = "\033[1;32m"
	lightBlueColor = "\033[1;36m"
	yellowColor    = "\033[1;33m"
)

type LogLevel int

const (
	ErrorLevel LogLevel = iota
	InfoLevel
	WarnLevel
	DebugLevel
)

type myLogger struct {
	logger *log.Logger
	level  LogLevel
}

func makeFormat(color string, argsCount int) string {
	format := ""
	for i := 0; i < argsCount; i++ {
		format += "%v "
	}
	return color + format + "\033[0m"
}

func Error(args ...interface{}) {
	if l.level >= 0 {
		l.logger.Printf(makeFormat(redColor, len(args)), args...)
	}
}

func Info(args ...interface{}) {
	if l.level >= 1 {
		l.logger.Printf(makeFormat(greenColor, len(args)), args...)
	}
}

func Warn(args ...interface{}) {
	if l.level >= 2 {
		l.logger.Printf(makeFormat(lightBlueColor, len(args)), args...)
	}
}

func Debug(args ...interface{}) {
	if l.level == 3 {
		l.logger.Printf(makeFormat(yellowColor, len(args)), args...)
	}
}

func Fatal(args ...interface{}) {
	l.logger.Printf(makeFormat(redColor, len(args)), args...)
	os.Exit(1)
}

var l myLogger

func InitLogger(writer io.Writer, prefix string, flag int, level LogLevel) {
	l = myLogger{
		logger: log.New(writer, prefix, flag),
		level:  level,
	}
}
