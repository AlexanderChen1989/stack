package static

import (
	"fmt"
	"net/http"
	"os"
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

var Default = Static(&Config{
	At:   "public",
	From: "static",
})

func allow(conf *Config, r *http.Request) (string, bool) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return "", false
	}

	cleaned := path.Clean(r.URL.Path)

	if !strings.HasPrefix(cleaned, conf.At) {
		return "", false
	}

	trimed := cleaned[len(conf.At):]

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

func fileEncode(conf *Config, file string) (string, os.FileInfo, error) {
	if conf.Brotli {
		path := file + ".br"
		info, err := os.Stat(path)
		if err == nil {
			return path, info, nil
		}
	}

	if conf.Gzip {
		path := file + ".gzip"
		info, err := os.Stat(path)
		if err == nil {
			return path, info, nil
		}
	}

	info, err := os.Stat(file)
	return file, info, err
}

func Static(conf *Config) stack.FrameFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		trimed, allowed := allow(conf, r)

		if !allowed {
			next(w, r)
			return
		}

		file, info, err := fileEncode(conf, path.Join(conf.From, trimed))
		if err != nil {
			next(w, r)
			return
		}
		// add cache

		fmt.Println(file, info, err)
	}
}
