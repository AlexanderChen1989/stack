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
