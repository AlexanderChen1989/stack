package plug

import (
	"net/http"

	"golang.org/x/net/context"
)

type Conn struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Context        context.Context
}
