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
	Levels  []log.Level
	Backand log.Handler
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
	if cfg.Backand == nil {
		cfg.Backand = console.New()
	}
	log.RegisterHandler(cfg.Backand, cfg.Levels...)
	return &logPlug{FieldLeveledLogger: log.Logger}
}

// New create new log plugger
func New() plug.Plugger {
	return NewWitConfig(Config{
		Levels: log.AllLevels,
	})
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
	l, _ := conn.Value(&logKey).(log.FieldLeveledLogger)

	return l
}
