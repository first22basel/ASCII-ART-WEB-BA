package main

import (
	"fmt"
	"net/http"

	BA "BA/internal/Functions"
)

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../internal/frontend"))))

	// Catch all routes with a custom handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			// Set 404 status and display a message
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "../internal/frontend/404.html")
			return
		}
		
		BA.FormHandler(w, r)
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
