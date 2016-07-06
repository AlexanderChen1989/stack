#Stack

## What's Stack?
* Stack is a simple web framework for golang
* A Frame is a middleware which can be stacked together
* Router, Handler, stack of Frames... everything is Frame

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

	b.Stack(requestid.New())

	router := mux.NewRouter()

	router.DispatchFunc(
		"/hello",
		func(conn plug.Conn) {
			fmt.Fprintln(conn, "Hello, world!")
		},
	)

	b.Stack(router)

	http.ListenAndServe(":8080", b.BuildHTTPHandler())
}
```

## Benchmark
* 10 Frame

```sh
➜  go test -run=XXX  -bench=.  -benchmem -v -benchtime=3s
BenchmarkStackFunc-4   	50000000	        75.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkStack-4       	50000000	       101 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/AlexanderChen1989/frame	9.077s
```

## Stack Architecture
```
                 request
                    +
                    |
               +---------+
               |  Frame  |
               +---------+
                    |
               +---------+
               |  Frame  |
               +---------+
                    |
            +----------------+
         +--+  Router(Frame) +--+
         |  +----------------+  |
         |                      |
         |                      |
         |                      v
         v                 +---------+
     +---------+           |  Frame  |
     |  Frame  |           +---------+
     +---------+           +---------+
          |                |  Frame  |
 +-----------------+       +---------+
 |  Handler(Frame) |            |
 +-----------------+   +-----------------+
                       |  Handler(Frame) |
                       +-----------------+
```
