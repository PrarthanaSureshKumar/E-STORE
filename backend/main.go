package main

import (
	"database/sql" // Go's built-in tool for talking to any database
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // the PostgreSQL driver (the _ means "import for side effects" — it registers itself with database/sql)
)

// db is a global variable that holds our database connection
// the whole app shares one connection pool
var db *sql.DB

// Product now matches our real database columns
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"image_url"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	// Read the ?category= query parameter from the URL
	// e.g. /products?category=Clothing → category = "Clothing"
	// e.g. /products               → category = "" (empty, return all)
	category := r.URL.Query().Get("category")

	var rows *sql.Rows
	var err error

	if category == "" {
		// No filter — return all products
		rows, err = db.Query("SELECT id, name, description, category, price, stock, image_url FROM products ORDER BY category, name")
	} else {
		// Filter by category — $1 is a placeholder, Go fills it in safely (prevents SQL injection)
		rows, err = db.Query("SELECT id, name, description, category, price, stock, image_url FROM products WHERE category = $1 ORDER BY name", category)
	}

	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		log.Println("DB error:", err)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Category, &p.Price, &p.Stock, &p.ImageURL)
		if err != nil {
			http.Error(w, "Failed to read product", http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func main() {
	// Connection string — tells Go where the database is and how to connect
	// format: "host=... port=... dbname=... sslmode=disable"
	connStr := "host=localhost port=5432 dbname=estore sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open DB connection:", err)
	}

	// Ping checks the connection is actually working
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot reach database:", err)
	}

	fmt.Println("Connected to database successfully!")

	http.HandleFunc("/products", getProducts)

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
