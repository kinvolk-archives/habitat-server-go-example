package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("service.conf")
	if err != nil {
		panic(err)
	}
	msg := string(f)

	io.WriteString(w, msg)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
