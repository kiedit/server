package main

import (
	"fmt"
	"kiedit/handler"
	"net/http"
)

func setupRoutes() {
	uploadHandler := new(handler.UploadHandler)

	http.HandleFunc("/upload", uploadHandler.UploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}
