package log

import (
	"fmt"
	"github.com/andrewmolyuk/pixar/exitor"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Ensure defaultLogger implements ILog interface
var _ ILog = (*defaultLogger)(nil)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
)

var colors = map[string]string{
	"DEBUG": "", "INFO": colorCyan, "WARN": colorYellow, "ERROR": colorRed,
}

type defaultLogger struct {
	secrets   []string
	debugMode bool
	lock      sync.Mutex
	isColored bool
	exitor    exitor.IExitor
}

func (l *defaultLogger) print(level string, args ...interface{}) {
	s := ""

	if len(args) == 1 {
		s = fmt.Sprint(args...)
	} else {
		f := fmt.Sprintf("%s", args[0])
		s = fmt.Sprintf(f, args[1:]...)
	}

	if level == "DEBUG" {
		skip := 0
		_, file, no, ok := runtime.Caller(skip)
		for ok {
			if !strings.HasSuffix(file, "/log/default.go") {
				file = strings.Split(file, "/")[len(strings.Split(file, "/"))-1]
				s = fmt.Sprintf("(%s:%d) %s", file, no, s)
				break
			}
			skip++
			_, file, no, ok = runtime.Caller(skip)
		}
	}

	s = fmt.Sprintf("%s [%s] %s", time.Now().Format("2006/01/02 15:04:05.000"), level, s)

	for _, secret := range l.secrets {
		s = strings.Replace(s, secret, "*****", -1)
	}

	if l.isColored {
		s = fmt.Sprint(colors[level], s, colorReset)
	}

	l.lock.Lock()
	defer l.lock.Unlock()
	fmt.Println(s)
}

func (l *defaultLogger) Debug(args ...interface{}) {
	if l.debugMode {
		l.print("DEBUG", args...)
	}
}

func (l *defaultLogger) Info(args ...interface{}) {
	l.print("INFO", args...)
}

func (l *defaultLogger) Warn(args ...interface{}) {
	l.print("WARN", args...)
}

func (l *defaultLogger) Error(args ...interface{}) {
	l.print("ERROR", args...)
	l.exitor.Exit(1)
}

// New creates a new instance of defaultLogger implementing ILog interface
func New(debugMode bool, secrets []string, isColored bool, e exitor.IExitor) ILog {
	if e == nil {
		e = exitor.Default()
	}
	return &defaultLogger{
		secrets:   secrets,
		debugMode: debugMode,
		isColored: isColored,
		exitor:    e,
	}
}

var logger ILog

// Default creates a new instance of defaultLogger with default values that can be used as a singleton
func Default() ILog {
	if logger == nil {
		InitDefault(false, nil, true, exitor.Default())
	}
	return logger
}

// InitDefault initializes default logger with specific values
func InitDefault(debugMode bool, secrets []string, isColored bool, e exitor.IExitor) {
	if e == nil {
		e = exitor.Default()
	}
	logger = New(debugMode, secrets, isColored, e)
}

// Debug prints debug message if debug mode is enabled
func Debug(args ...interface{}) {
	Default().Debug(args...)
}

// Info prints info message
func Info(args ...interface{}) {
	Default().Info(args...)
}

// Warn prints warn message
func Warn(args ...interface{}) {
	Default().Warn(args...)
}

// Error prints error message and exits with code 1
func Error(args ...interface{}) {
	Default().Error(args...)
}
