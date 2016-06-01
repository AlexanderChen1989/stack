// Package gplog is a adapter for github.com/go-playground/log
package gplog

import (
	"github.com/AlexanderChen1989/plug"
	"github.com/go-playground/log"
	"github.com/go-playground/log/handlers/console"
)

type logPlug struct {
	log.FieldLeveledLogger
	next plug.Plugger
}

// Config config logger plugger
type Config struct {
	Levels     []log.Level
	LogHandler log.Handler
}

// EmptyHandler is handler for logger plug, which will discard all log message
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

// NewWitConfig create log plugger with config,
// if config.Backand is nil, console handler will be used
func NewWitConfig(cfg Config) plug.Plugger {
	if cfg.LogHandler == nil {
		cfg.LogHandler = console.New()
	}
	log.RegisterHandler(cfg.LogHandler, cfg.Levels...)
	return &logPlug{FieldLeveledLogger: log.Logger}
}

// New create new log plugger
func New(fns ...func(Config) Config) plug.Plugger {
	cfg := Config{
		Levels: log.AllLevels,
	}
	for _, fn := range fns {
		cfg = fn(cfg)
	}
	if cfg.LogHandler == nil {
		cfg.LogHandler = console.New()
	}
	return NewWitConfig(cfg)
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
func Info() func(Config) Config {
	return Level(log.InfoLevel)
}

// Debug log debug level
func Debug() func(Config) Config {
	return Level(log.DebugLevel)
}

// Warn log warn level
func Warn() func(Config) Config {
	return Level(log.WarnLevel)
}

// Error log error level
func Error() func(Config) Config {
	return Level(log.ErrorLevel)
}

// Level config log level
func Level(l log.Level) func(Config) Config {
	var ls []log.Level
	for i, level := range allLevels {
		if level == l {
			ls = allLevels[i:]
			break
		}
	}

	return func(cfg Config) Config {
		cfg.Levels = ls
		return cfg
	}
}

// LogHandler config log handler
func LogHandler(h log.Handler) func(Config) Config {
	return func(cfg Config) Config {
		cfg.LogHandler = h
		return cfg
	}
}

func (l *logPlug) Plug(next plug.Plugger) plug.Plugger {
	l.next = next
	return l
}

var logKey int

func (l *logPlug) HandleConn(conn plug.Conn) {
	conn = plug.WithValue(conn, &logKey, l)

	l.next.HandleConn(conn)
}

// Logger return inject logger instance, you have to add log plug first
func Logger(conn plug.Conn) log.FieldLeveledLogger {
	l, ok := conn.Value(&logKey).(log.FieldLeveledLogger)
	if !ok {
		return nil
	}
	return l
}
