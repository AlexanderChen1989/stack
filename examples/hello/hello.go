package main

import (
	"net/http"

	"github.com/AlexanderChen1989/plug"
	"github.com/AlexanderChen1989/plug/plugs/requestid"
)

func main() {
	b := plug.NewBuilder()

	b.Plug(requestid.New())

	http.ListenAndServe(":8080", b.BuildHTTPHandler())
}
