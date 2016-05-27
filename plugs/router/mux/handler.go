package mux

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"golang.org/x/net/context"
)

type Handler interface {
	ServeHTTP(context.Context, http.ResponseWriter, *http.Request)
}

type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	f(ctx, w, r)
}

// NotFound replies to the request with an HTTP 404 not found error.
func NotFound(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 page not found", http.StatusNotFound)
}

// NotFoundHandler returns a simple request handler
// that replies to each request with a ``404 page not found'' reply.
func NotFoundHandler() Handler { return HandlerFunc(NotFound) }

// Redirect to a fixed URL
type redirectHandler struct {
	url  string
	code int
}

func (rh *redirectHandler) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	Redirect(w, r, rh.url, rh.code)
}

// RedirectHandler returns a request handler that redirects
// each request it receives to the given url using the given
// status code.
//
// The provided code should be in the 3xx range and is usually
// StatusMovedPermanently, StatusFound or StatusSeeOther.
func RedirectHandler(url string, code int) Handler {
	return &redirectHandler{url, code}
}

// The provided code should be in the 3xx range and is usually
// StatusMovedPermanently, StatusFound or StatusSeeOther.
func Redirect(w http.ResponseWriter, r *http.Request, urlStr string, code int) {
	if u, err := url.Parse(urlStr); err == nil {
		// If url was relative, make absolute by
		// combining with request path.
		// The browser would probably do this for us,
		// but doing it ourselves is more reliable.

		// NOTE(rsc): RFC 2616 says that the Location
		// line must be an absolute URI, like
		// "http://www.google.com/redirect/",
		// not a path like "/redirect/".
		// Unfortunately, we don't know what to
		// put in the host name section to get the
		// client to connect to us again, so we can't
		// know the right absolute URI to send back.
		// Because of this problem, no one pays attention
		// to the RFC; they all send back just a new path.
		// So do we.
		if u.Scheme == "" && u.Host == "" {
			oldpath := r.URL.Path
			if oldpath == "" { // should not happen, but avoid a crash if it does
				oldpath = "/"
			}

			// no leading http://server
			if urlStr == "" || urlStr[0] != '/' {
				// make relative path absolute
				olddir, _ := path.Split(oldpath)
				urlStr = olddir + urlStr
			}

			var query string
			if i := strings.Index(urlStr, "?"); i != -1 {
				urlStr, query = urlStr[:i], urlStr[i:]
			}

			// clean up but preserve trailing slash
			trailing := strings.HasSuffix(urlStr, "/")
			urlStr = path.Clean(urlStr)
			if trailing && !strings.HasSuffix(urlStr, "/") {
				urlStr += "/"
			}
			urlStr += query
		}
	}

	w.Header().Set("Location", urlStr)
	w.WriteHeader(code)

	// RFC2616 recommends that a short note "SHOULD" be included in the
	// response because older user agents may not understand 301/307.
	// Shouldn't send the response for POST or HEAD; that leaves GET.
	if r.Method == "GET" {
		note := "<a href=\"" + htmlEscape(urlStr) + "\">" + http.StatusText(code) + "</a>.\n"
		fmt.Fprintln(w, note)
	}
}

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	// "&#34;" is shorter than "&quot;".
	`"`, "&#34;",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	"'", "&#39;",
)

func htmlEscape(s string) string {
	return htmlReplacer.Replace(s)
}
