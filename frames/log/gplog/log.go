// Package gplog is a adapter for github.com/go-playground/log
package gplog

import (
	"context"
	"net/http"

	"github.com/AlexanderChen1989/stack"
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
)

type config struct {
	Levels     []log.Level
	LogHandler log.Handler
}

// EmptyHandler is handler for logger frame, which will discard all log message
type EmptyHandler struct{}

// Run implements log.Handler
func (empty EmptyHandler) Run() chan<- *log.Entry {
	ch := make(chan *log.Entry)
	go func() {
		for {
			<-ch
		}
	}()
	return ch
}

// newWitConfig create log Frame with config,
// if config.Backand is nil, console handler will be used
func newWitConfig(cfg config) stack.FrameFunc {
	if cfg.LogHandler == nil {
		cfg.LogHandler = console.New()
	}
	log.RegisterHandler(cfg.LogHandler, cfg.Levels...)

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		ctx := context.WithValue(r.Context(), &logKey, log.Logger)
		next(w, r.WithContext(ctx))
	}
}

// New create new log Frame
func New(fns ...func(config) config) stack.FrameFunc {
	cfg := config{
		Levels: log.AllLevels,
	}
	for _, fn := range fns {
		cfg = fn(cfg)
	}
	if cfg.LogHandler == nil {
		cfg.LogHandler = console.New()
	}
	return newWitConfig(cfg)
}

var allLevels = []log.Level{
	log.DebugLevel,
	log.TraceLevel,
	log.InfoLevel,
	log.NoticeLevel,
	log.WarnLevel,
	log.ErrorLevel,
	log.PanicLevel,
	log.AlertLevel,
	log.FatalLevel,
}

// Info log info level
func Info() func(config) config {
	return Level(log.InfoLevel)
}

// Debug log debug level
func Debug() func(config) config {
	return Level(log.DebugLevel)
}

// Warn log warn level
func Warn() func(config) config {
	return Level(log.WarnLevel)
}

// Error log error level
func Error() func(config) config {
	return Level(log.ErrorLevel)
}

// Level config log level
func Level(l log.Level) func(config) config {
	var ls []log.Level
	for i, level := range allLevels {
		if level == l {
			ls = allLevels[i:]
			break
		}
	}

	return func(cfg config) config {
		cfg.Levels = ls
		return cfg
	}
}

// LogHandler config log handler
func LogHandler(h log.Handler) func(config) config {
	return func(cfg config) config {
		cfg.LogHandler = h
		return cfg
	}
}

var logKey int

// Logger return inject logger instance, you have to add log Frame first
func Logger(r *http.Request) log.LeveledLogger {
	l, ok := r.Context().Value(&logKey).(log.LeveledLogger)
	if !ok {
		return nil
	}
	return l
}
