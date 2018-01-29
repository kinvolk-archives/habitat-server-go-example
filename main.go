package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func getConfigPath() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "./exampleconfs/service.conf"
}

func hello(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		panic(err)
	}
	msg := string(f)
	io.WriteString(w, msg)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":5555", nil)
}
