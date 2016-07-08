package static

import (
	"net/http"
	"path"
	"strings"

	"github.com/AlexanderChen1989/stack"
)

type Config struct {
	At           string
	From         string
	Gzip         bool
	Brotli       bool
	Only         []string
	OnlyMatching []string
	Prefix       string
	CacheVSN     string
	CacheEtg     string
	Headers      map[string]string
}

func (conf *Config) setup() {
	if conf.At == "" || conf.From == "" {
		panic("Config.At or Config.From cant be empty")
	}
	conf.At = path.Join("/", conf.At)
	conf.From = path.Join("/", conf.From)

	for i, d := range conf.Only {
		if strings.ContainsRune(d, '.') {
			conf.Only[i] = path.Join("/", d)
			continue
		}
		conf.Only[i] = path.Join("/", d) + "/"
	}

	for i, d := range conf.OnlyMatching {
		conf.OnlyMatching[i] = path.Join("/", d)
	}

	if conf.CacheVSN == "" {
		conf.CacheVSN = "public, max-age=31536000"
	}

	if conf.CacheEtg == "" {
		conf.CacheEtg = "public"
	}

	if conf.Headers == nil {
		conf.Headers = map[string]string{}
	}
}

var Default = Static(Config{
	At:   "public",
	From: "static",
})

func allow(conf *Config, r *http.Request) (string, bool) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return "", false
	}

	if !strings.HasPrefix(r.URL.Path, conf.At) {
		return "", false
	}

	trimed := r.URL.Path[len(conf.At):]

	for _, onlyMatch := range conf.OnlyMatching {
		if strings.HasPrefix(trimed, onlyMatch) {
			return trimed, true
		}
	}

	for _, only := range conf.Only {
		if strings.HasPrefix(trimed, only) {
			return trimed, true
		}
	}

	return "", false
}

func Static(conf Config) stack.FrameFunc {

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			next(w, r)
			return
		}

		// serve files

	}
}
