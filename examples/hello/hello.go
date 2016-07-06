package main

import (
	"fmt"
	"net/http"

	"github.com/AlexanderChen1989/stack"
	"github.com/AlexanderChen1989/stack/frames/log/gplog"
	"github.com/AlexanderChen1989/stack/frames/requestid"
	"github.com/AlexanderChen1989/stack/frames/router"
	"github.com/gorilla/mux"
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

	b.PushFunc(gplog.New())
	b.PushFunc(gplog.Trace)
	b.PushFunc(requestid.New(""))
	b.PushFunc(router.New(setupRouter()))

	http.ListenAndServe(":8080", b.Build())
}
