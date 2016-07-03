# Plug

## What's Plug?
* Plug is a simple web framework for golang
* A Plugger is a middleware which can be stacked together
* Router, Handler, stack of Pluggers... everything is Plugger

## Example

```go
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

## Benchmark
* 10 Plug

```sh
➜  plug git:(go1.7) go test -run=XXX  -bench=.  -benchmem -v -benchtime=3s
BenchmarkPlugFunc-4   	50000000	        75.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkPlug-4       	50000000	       101 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/AlexanderChen1989/plug	9.077s
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
