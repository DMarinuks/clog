package clog

import (
	"io"
	"log"
	"os"
	"time"
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
	logger 		*log.Logger
	level  		LogLevel
	showTime 	bool
	timePattern string
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
		if l.showTime {
			args = append([]interface{}{time.Now().Format(l.timePattern)}, args...)
		}
		l.logger.Printf(makeFormat(redColor, len(args)), args...)
	}
}

func Info(args ...interface{}) {
	if l.level >= 1 {
		if l.showTime {
			args = append([]interface{}{time.Now().Format(l.timePattern)}, args...)
		}
		l.logger.Printf(makeFormat(greenColor, len(args)), args...)
	}
}

func Warn(args ...interface{}) {
	if l.level >= 2 {
		if l.showTime {
			args = append([]interface{}{time.Now().Format(l.timePattern)}, args...)
		}
		l.logger.Printf(makeFormat(lightBlueColor, len(args)), args...)
	}
}

func Debug(args ...interface{}) {
	if l.level == 3 {
		if l.showTime {
			args = append([]interface{}{time.Now().Format(l.timePattern)}, args...)
		}
		l.logger.Printf(makeFormat(yellowColor, len(args)), args...)
	}
}

func Fatal(args ...interface{}) {
	if l.showTime {
		args = append([]interface{}{time.Now().Format(l.timePattern)}, args...)
	}
	l.logger.Printf(makeFormat(redColor, len(args)), args...)
	os.Exit(1)
}

var l myLogger

func ShowTime(showTime bool) {
	l.showTime = showTime
}

func SetTimeFormat(format string) {
	l.timePattern = format
}

func InitLogger(writer io.Writer, prefix string, flag int, level LogLevel) {
	l = myLogger{
		logger: log.New(writer, prefix, flag),
		level:  level,
		showTime: true,
		timePattern: "15:04:05 02-01-2006",
	}
}
