package clog

import (
	"bytes"
	"os"
	"regexp"
	"testing"
)

var got bytes.Buffer

func TestMain(m *testing.M) {
	InitLogger(&got, "", 0, DebugLevel)
	ShowTime(false)
	os.Exit(m.Run())
}

func TestError(t *testing.T) {
	want := "\u001B[1;31mERROR: test \u001B[0m\n"
	Error("test")
	if got.String() != want {
		t.Errorf("Error() = %v, want %v", got.String(), want)
	}
}

func TestInfo(t *testing.T) {
	got.Reset()
	want := "\033[1;32mINFO: test \u001B[0m\n"
	Info("test")
	if got.String() != want {
		t.Errorf("Info() = %v, want %v", got.String(), want)
	}
}

func TestWarn(t *testing.T) {
	got.Reset()
	want := "\u001B[1;36mWARNING: test \u001B[0m\n"
	Warn("test")
	if got.String() != want {
		t.Errorf("Warn() = %v, want %v", got.String(), want)
	}
}

func TestDebug(t *testing.T) {
	got.Reset()
	want := "\u001B[1;33mDEBUG: test \u001B[0m\n"
	Debug("test")
	if got.String() != want {
		t.Errorf("Debug() = %v, want %v", got.String(), want)
	}
}

func TestShowTime(t *testing.T) {
	timeRE := regexp.MustCompile(`(\d{4}-\d{1,2}-\d{1,2}\s\d{1,2}:\d{1,2}:\d{1,2})`)
	ShowTime(true)
	Debug("test")
	regResults := timeRE.FindStringSubmatch(got.String())
	if len(regResults) != 2 {
		t.Errorf("ShowTime() = %v, want time in output", got.String())
	}
}

func TestSetTimeFormat(t *testing.T) {
	timeRE := regexp.MustCompile(`(\d{1,2}:\d{1,2}:\d{1,2})`)
	ShowTime(true)
	SetTimeFormat("15:04:05")
	Debug("test")
	regResults := timeRE.FindStringSubmatch(got.String())
	if len(regResults) != 2 {
		t.Errorf("ShowTime() = %v, want time in output", got.String())
	}
}