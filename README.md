# Plug

## What's Plug?
* Plug is a simple web framework for golang
* A Plugger is a middleware which can be stacked together
* Router, Handler... everything is Plugger

## Example

```golang
package main

import (
	"fmt"
	"net/http"

	"github.com/AlexanderChen1989/plug"
	"github.com/AlexanderChen1989/plug/plugs/requestid"
	"github.com/AlexanderChen1989/plug/plugs/router/mux"
)

func main() {
	b := plug.NewBuilder()

	b.Plug(requestid.New())

	router := mux.NewRouter()

	router.DispatchFunc(
		"/hello",
		func(conn plug.Conn) {
			fmt.Fprintln(conn, "Hello, world!")
		},
	)

	b.Plug(router)

	http.ListenAndServe(":8080", b.BuildHTTPHandler())
}
```

## Plug Architecture
```
                 request
                    +
                    |
               +---------+
               | Plugger |
               +---------+
                    |
               +---------+
               | Plugger |
               +---------+
                    |
            +----------------+
         +--+ Router(Plugger)+--+
         |  +----------------+  |
         |                      |
         |                      |
         |                      v
         v                 +---------+
     +---------+           | Plugger |
     | Plugger |           +---------+
     +---------+           +---------+
          |                | Plugger |
 +-----------------+       +---------+
 | Handler(Plugger)|            |
 +-----------------+   +-----------------+
                       | Handler(Plugger)|
                       +-----------------+
```
