package static

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	conf := &Config{
		At:           "public",
		From:         "static",
		Only:         []string{"images", "js", "favicon.co"},
		OnlyMatching: []string{"css"},
	}
	conf.setup()
	fmt.Println(conf)

	// test allow
	paths := []struct {
		path    string
		trimed  string
		matched bool
	}{
		{"/public/images/hello.png", "/images/hello.png", true},
		{"/public/images-hello.png", "/images/hello.png", false},
		{"/public/favicon.co", "/favicon.co", true},
		{"/public/css-style.css", "/css-style.css", true},
		{"/public/css/style.css", "/css/style.css", true},
	}
	for _, p := range paths {
		r, _ := http.NewRequest("GET", p.path, nil)
		trimed, matched := allow(conf, r)
		assert.Equal(t, p.matched, matched)
		if p.matched {
			assert.Equal(t, p.trimed, trimed)
		}
	}
}
