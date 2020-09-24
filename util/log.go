package util

import "log"

// This may be dumb, but I want a layer of indirection
// to have freedom to replace logging implementations
// without ripping it all out
type Logger interface {
	Info(...interface{})
	Error(...interface{})
	Fatal(...interface{})
}

func GetLogger() Logger {
	return &logger{}
}

type logger struct{}

func (l *logger) Error(msg ...interface{}) {
	log.Print(msg...)
}

func (l *logger) Fatal(msg ...interface{}) {
	log.Fatal(msg...)
}

func (l *logger) Info(msg ...interface{}) {
	log.Print(msg...)
}
