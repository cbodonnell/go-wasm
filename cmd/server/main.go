package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Run server
	// fmt.Println("Server started.")
	const port = 8080
	fmt.Printf("Started server on port %v.\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), http.FileServer(http.Dir("./assets"))))
}
