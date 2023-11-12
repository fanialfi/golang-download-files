package main

import (
	"golang-download-files/handler"
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", handler.HandleIndex)
	http.HandleFunc("/list-files", handler.HandleListFiles)
	http.HandleFunc("/download", handler.HandleDownload)

	log.Printf("server running on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
