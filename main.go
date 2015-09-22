package main

import (
	"github.com/zhuharev/s/shortener"
	"net/http"
)

func main() {

	srv, e := shortener.New("cnf")
	if e != nil {
		panic(e)
	}

	http.Handle("/s/", srv)
	e = http.ListenAndServe(":8089", nil)
	if e != nil {
		panic(e)
	}
}
