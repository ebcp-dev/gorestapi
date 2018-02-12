package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct (Model example)
type Book struct {
	// Properties start with capitals
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	// Reference Author struct
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as slice of Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Send books as json
	json.NewEncoder(w).Encode(books)
}

// Get book
func getBook(w http.ResponseWriter, r *http.Request) {
	// Get parameters from request
	params := mux.Vars(r)

	// Loop through books slice and find book by id
	// _ is for temporary, unused variable
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// Create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// Create random id and convert to string
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)

			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Initialize mux router
	router := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "11111", Title: "Book One", Author: &Author{Firstname: "Johnny", Lastname: "One"}})
	books = append(books, Book{ID: "2", Isbn: "22222", Title: "Book Two", Author: &Author{Firstname: "Johnny", Lastname: "Two"}})
	books = append(books, Book{ID: "3", Isbn: "33333", Title: "Book Three", Author: &Author{Firstname: "Johnny", Lastname: "Three"}})
	books = append(books, Book{ID: "4", Isbn: "44444", Title: "Book Four", Author: &Author{Firstname: "Johnny", Lastname: "Four"}})
	books = append(books, Book{ID: "5", Isbn: "55555", Title: "Book Five", Author: &Author{Firstname: "Johnny", Lastname: "Five"}})

	// Create route handlers to establish endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Log failures
	log.Fatal(http.ListenAndServe(":8000", router))
}
