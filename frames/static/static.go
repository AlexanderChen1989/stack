package static

import (
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/AlexanderChen1989/stack"
)

var Default = Static(Config{
	At:   "public",
	From: "static",
})

type Config struct {
	At      string
	From    string
	Headers map[string]string
}

func (conf *Config) setup() {
	if conf.At == "" || conf.From == "" {
		panic("Config.At or Config.From cant be empty")
	}
	conf.At = path.Join("/", conf.At)
	from, err := filepath.Abs(conf.From)
	if err != nil {
		panic(err)
	}
	conf.From = from
}

func allow(conf *Config, r *http.Request) bool {
	return (r.Method == http.MethodGet || r.Method == http.MethodHead) && strings.HasPrefix(r.URL.Path, conf.At)
}

func Static(conf Config) stack.FrameFunc {
	conf.setup()

	h := http.StripPrefix(
		conf.At,
		http.FileServer(
			http.Dir(conf.From),
		),
	)

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if allow(&conf, r) {
			for k, v := range conf.Headers {
				w.Header().Add(k, v)
			}
			h.ServeHTTP(w, r)
			return
		}

		next(w, r)
	}
}
