package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// * Routes
	func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Welcome to my website!")
		})

		// Route for "/about" page
		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "About the website")
		})

		// Route for "/contact" page
		http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Contact us at: contact@example.com")
		})
	}()

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define the port
	port := ":8080" // for example, on localhost:8080
	fmt.Printf("Server starting on port %s\n", port)

	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	

	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
