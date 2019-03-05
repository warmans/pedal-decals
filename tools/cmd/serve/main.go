package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("The first argument should be the dir to serve")
	}
	fs := http.FileServer(http.Dir(os.Args[1]))
	http.Handle("/", fs)

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}