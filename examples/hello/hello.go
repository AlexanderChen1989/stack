package main

import (
	"fmt"
	"net/http"

	"github.com/AlexanderChen1989/plug"
	"github.com/AlexanderChen1989/plug/plugs/requestid"
)

func main() {
	b := plug.NewBuilder()

	b.Plug(requestid.New())
	b.Plug(requestid.NewWithHeader("My-Request-ID"))

	b.Plug(plug.HandleFunc(func(conn plug.Conn) {
		fmt.Fprintln(conn, "Hello, world!")
	}))

	http.ListenAndServe(":8080", b.BuildHTTPHandler())
}
