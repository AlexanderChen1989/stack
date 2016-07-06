/*Package head is a Plug to convert `HEAD` requests to `GET` requests.

## Examples

  b := plug.NewBuilder()
  b.Plug(head.PlugFunc)
*/
package head

import (
	"net/http"
	"strings"
)

func PlugFunc(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if strings.ToUpper(r.Method) == http.MethodHead {
		r.Method = http.MethodGet
	}

	next(w, r)
}
