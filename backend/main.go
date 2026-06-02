package main

// "package main" tells Go: this is the starting point of the program

import (
	"encoding/json" // lets us convert Go data into JSON format
	"fmt"           // lets us print things
	"net/http"      // gives us the web server tools
)

// This is a "struct" — it's how Go describes what a Product looks like.
// Think of it as a blueprint. Every product has these 3 fields.
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// This function runs when someone hits GET /products
// w = where you write your response back to the caller
// r = the incoming request (has info like the URL, method, etc.)
func getProducts(w http.ResponseWriter, r *http.Request) {
	// Fake data for now — later this will come from PostgreSQL
	products := []Product{
		{ID: 1, Name: "T-Shirt", Price: 19.99},
		{ID: 2, Name: "Hoodie", Price: 49.99},
		{ID: 3, Name: "Cap", Price: 14.99},
	}

	// Tell the caller "the response will be JSON"
	w.Header().Set("Content-Type", "application/json")

	// Send back 200 OK status code
	w.WriteHeader(http.StatusOK)

	// Convert the products list to JSON and write it to the response
	json.NewEncoder(w).Encode(products)
}

// main() is where the program starts — Go always runs this first
func main() {
	// "when someone hits /products, run the getProducts function"
	http.HandleFunc("/products", getProducts)

	// Start the server on port 8080 and print a message so we know it's running
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
