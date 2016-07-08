package static

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexanderChen1989/stack"
	"github.com/stretchr/testify/assert"
)

func TestStatic(t *testing.T) {
	headers := map[string]string{
		"X-Test": "Hello, world!",
	}

	fn := Static(Config{
		From:    "./test_static",
		At:      "public",
		Headers: headers,
	})

	b := stack.NewBuilder()
	b.PushFunc(fn)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/public/style.css", nil)
	b.Build().ServeHTTP(w, r)

	for k, v := range headers {
		assert.Equal(t, v, w.Header().Get(k))
	}

	data, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)
	data2, err := ioutil.ReadFile("./test_static/style.css")
	assert.Nil(t, err)
	assert.Equal(t, data, data2)
}
