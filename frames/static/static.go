package static

import (
	"net/http"

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

var Default = Static(Config{
	At:   "public",
	From: "static",
})

func allow(conf *Config, r *http.Request) bool {
	// check path based on only and prefix
	return true
}

func filePath(conf *Config, r *http.Request) (string, error) {
	return "", nil
}

func Static(conf Config) stack.FrameFunc {
	if conf.At == "" || conf.From == "" {
		panic("Config.At or Config.From cant be empty")
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

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			next(w, r)
			return
		}

		// serve files

	}
}
