package logger

import (
	"context"
	"io"
	"log"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

// Logger ...
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
}

// NewLogger ...
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// WithLevel ...
func (l *Logger) WithLevel(lvl Level) *Logger {
	ll := l.clone()
	ll.level = lvl
	return ll
}

func (l *Logger) WithFields(lvl Level) *Logger {
}
