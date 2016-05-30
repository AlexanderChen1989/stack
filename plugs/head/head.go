/*Package head is a Plug to convert `HEAD` requests to `GET` requests.

## Examples

  b := plug.NewBuilder()
  b.Plug(head.New())
*/
package head

import (
	"net/http"
	"strings"

	"github.com/AlexanderChen1989/plug"
)

type headPlug struct {
	next plug.Plugger
}

// New create head Plugger
func New() plug.Plugger {
	return &headPlug{}
}

func (p *headPlug) Plug(next plug.Plugger) plug.Plugger {
	p.next = next
	return p
}

func (p *headPlug) HandleConn(conn plug.Conn) {
	if strings.ToUpper(conn.Request.Method) == http.MethodHead {
		conn.Request.Method = http.MethodGet
	}
	p.next.HandleConn(conn)
}
