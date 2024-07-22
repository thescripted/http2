package main

import (
	"github.com/thescripted/http/http2"
)

func main() {
	// send "hello world" via HTTP/2 to localhost:8080
	conn, err := http2.Connect("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	resp, err := conn.Request("GET", "localhost:8080", "hello world")
	if err != nil {
		panic(err)
	}

	println(resp)

	// start an HTTP/2 server on localhost:8080
	srv, err := http2.NewServer()
	if err != nil {
		panic(err)
	}

	if err := srv.ListenAndServe("http://localhost:8080"); err != nil {
		panic(err)
	}
}
