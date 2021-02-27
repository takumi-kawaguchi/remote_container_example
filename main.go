package main

import (
	"fmt"
	"io"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world.")
}

func main() {
	portNumber := "1324"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}