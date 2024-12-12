// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	handler "webServer/Handler"
)

func main() {
	fmt.Println("Server starting.............")
	// this is just a comment to check the pipeline
	// Serve static files from the ./static directory
	server := http.FileServer(http.Dir("./static"))
	http.Handle("/", server)

	http.HandleFunc("/form", handler.FormHandler)
	http.HandleFunc("/hello", handler.HelloHandler)

	fmt.Println("Starting server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
