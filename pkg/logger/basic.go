package logger

import (
	"fmt"
	"os"
)

// BasicLogger logs all messages to stdout
type BasicLogger struct{}

var _ LoggerImpl = &BasicLogger{}

func (log *BasicLogger) print(level string, args ...any) {
	fmt.Print(level + ": ")
	fmt.Println(args...)
}

func (log *BasicLogger) printf(level string, format string, args ...any) {
	fmt.Printf(level+": "+format+"\n", args...)
}

func (log *BasicLogger) Progressf(format string, args ...any) {
	log.printf("Progress", format, args...)
}

func (log *BasicLogger) Trace(args ...any) {
	log.print("Trace", args...)
}

func (log *BasicLogger) Tracef(format string, args ...any) {
	log.printf("Trace", format, args...)
}

func (log *BasicLogger) TraceFunc(fn func() (string, []any)) {
	format, args := fn()
	log.printf("Trace", format, args...)
}

func (log *BasicLogger) Debug(args ...any) {
	log.print("Debug", args...)
}

func (log *BasicLogger) Debugf(format string, args ...any) {
	log.printf("Debug", format, args...)
}

func (log *BasicLogger) DebugFunc(fn func() (string, []any)) {
	format, args := fn()
	log.printf("Debug", format, args...)
}

func (log *BasicLogger) Info(args ...any) {
	log.print("Info", args...)
}

func (log *BasicLogger) Infof(format string, args ...any) {
	log.printf("Info", format, args...)
}

func (log *BasicLogger) InfoFunc(fn func() (string, []any)) {
	format, args := fn()
	log.printf("Info", format, args...)
}

func (log *BasicLogger) Warn(args ...any) {
	log.print("Warn", args...)
}

func (log *BasicLogger) Warnf(format string, args ...any) {
	log.printf("Warn", format, args...)
}

func (log *BasicLogger) WarnFunc(fn func() (string, []any)) {
	format, args := fn()
	log.printf("Warn", format, args...)
}

func (log *BasicLogger) Error(args ...any) {
	log.print("Error", args...)
}

func (log *BasicLogger) Errorf(format string, args ...any) {
	log.printf("Error", format, args...)
}

func (log *BasicLogger) ErrorFunc(fn func() (string, []any)) {
	format, args := fn()
	log.printf("Error", format, args...)
}

func (log *BasicLogger) Fatal(args ...any) {
	log.print("Fatal", args...)
	os.Exit(1)
}

func (log *BasicLogger) Fatalf(format string, args ...any) {
	log.printf("Fatal", format, args...)
	os.Exit(1)
}
