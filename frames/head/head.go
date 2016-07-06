/*Package head is a Frame to convert `HEAD` requests to `GET` requests.

## Examples

  b := plug.NewBuilder()
  b.Plug(head.FrameFunc)
*/
package head

import (
	"net/http"
	"strings"
)

func FrameFunc(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if strings.ToUpper(r.Method) == http.MethodHead {
		r.Method = http.MethodGet
	}

	next(w, r)
}
