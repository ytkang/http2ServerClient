package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"os"
)

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}

func main() {
	h2s := &http2.Server{}
	mux := NewMultiplexer()
	mux.HandleFunc("/test", TestRequestHandler)

	server := &http.Server{
		Addr:    "0.0.0.0:9998",
		Handler: h2c.NewHandler(mux.Handler, h2s),
	}

	fmt.Printf("Listening [0.0.0.0:9998]...\n")
	checkErr(server.ListenAndServe(), "while listening")
}