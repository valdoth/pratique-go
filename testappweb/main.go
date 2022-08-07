package main

import (
	"fmt"
	"net/http"
)

const port = ":8000"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("http://localhost:8000 - Server started on port: ", port)
	http.ListenAndServe(port, nil)
}
