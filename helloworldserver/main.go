package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server started")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":5000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.Header.Get("message"))
}
