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
	os.Exit(m.Run())
}

func TestError(t *testing.T) {
	want := "\u001B[1;31mtest \u001B[0m\n"
	Error("test")
	if got.String() != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestInfo(t *testing.T) {
	got.Reset()
	want := "\033[1;32mtest \u001B[0m\n"
	Info("test")
	if got.String() != want {
		t.Errorf("Info() = %v, want %v", got, want)
	}
}

func TestWarn(t *testing.T) {
	got.Reset()
	want := "\u001B[1;36mtest \u001B[0m\n"
	Warn("test")
	if got.String() != want {
		t.Errorf("Warn() = %v, want %v", got, want)
	}
}

func TestDebug(t *testing.T) {
	got.Reset()
	want := "\u001B[1;33mtest \u001B[0m\n"
	Debug("test")
	if got.String() != want {
		t.Errorf("Debug() = %v, want %v", got, want)
	}
}

func TestShowTime(t *testing.T) {
	timeRE := regexp.MustCompile(`(\d{1,2}:\d{1,2}:\d{1,2}\s\d{1,2}-\d{1,2}-\d{4})`)
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