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

	"github.com/gorilla/mux"
	"github.com/AlexanderChen1989/stack"
	"github.com/AlexanderChen1989/stack/frames/requestid"
	"github.com/AlexanderChen1989/stack/frames/router"
)

func setupRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!\n")
	})
	return r
}

func main() {
	b := stack.NewBuilder()

	b.PushFunc(requestid.New())
	b.PushFunc(router.New(setupRouter()))

	http.ListenAndServe(":8080", b.Build())
}
```

## Benchmark
* 10 Frame

```sh
➜  go test -run=XXX  -bench=.  -benchmem -v -benchtime=3s
BenchmarkFrameFunc-4   	50000000	        76.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkFrame-4       	50000000	       102 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/AlexanderChen1989/stack	9.162s
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
