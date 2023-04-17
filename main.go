package main

import (
	"ascii-art-web/handlers"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/ascii-art", handlers.ProcessorHandler)
	mux.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("templates"))))
	fmt.Println("Server launched ...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("this host is already run")
	}
}
