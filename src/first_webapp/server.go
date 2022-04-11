package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	fmt.Println("Hello golang from docker!")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
