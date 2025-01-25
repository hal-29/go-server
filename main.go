package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Server running on http://localhost:8080\n")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
