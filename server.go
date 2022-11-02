package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "80"

func main() {
	fmt.Printf("Starting server at port %v\n", port)
	http.Handle("/", http.FileServer(http.Dir("./")))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
