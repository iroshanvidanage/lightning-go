package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Reserved bool   `json:"reserved"`
}

var books = []Book{}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// Helper func to list the available books in the library
	// Only the id and title will be returned
	type BookSummary struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	}

	var booksum []BookSummary
	// Create the summary of the books available
	for _, b := range books {
		booksum = append(booksum, BookSummary{
			Id:    b.Id,
			Title: b.Title,
		})
	}

	json.NewEncoder(w).Encode(booksum)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	// Helper function to get the book details
	// The request should refer to the id of the book
	// The id is a string comes via the request
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Book ID type", http.StatusBadRequest)
		return
	}

	// If no error, look for the book id
	for _, book := range books {
		if book.Id == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	// If book id not found
	http.NotFound(w, r)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	// Helper function to add a new book
	var newBook Book
	// Valid body type '{"title": "tile of the book", "author": "author of the book"}'
	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		http.Error(w, "Invalid request body type", http.StatusBadRequest)
		return
	}

	// Fill missing fields
	newBook.Id = len(books) + 1
	newBook.Reserved = false

	// Add to the library
	books = append(books, newBook)

	// Return the response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func reserveBook(w http.ResponseWriter, r *http.Request) {
	// Helper function to reserve a book
	idStr := strings.TrimPrefix(r.URL.Path, "/reserve/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Book ID type", http.StatusBadRequest)
		return
	}

	// If no error, look for the book id
	for i, book := range books {
		if book.Id == id {
			if !book.Reserved {
				books[i].Reserved = true
				w.Header().Set("X-Reserved", "Reserved")
				io.WriteString(w, "Reserved: ")
				json.NewEncoder(w).Encode(book.Title)
				return
			} else {
				// Set a custom header
				w.Header().Set("X-Reserved", "NotAvailable")
				// w.WriteHeader(http.StatusOK)
				io.WriteString(w, "This book is already reserved\n")
				return
			}
		}
	}

	// If book id not found
	http.NotFound(w, r)
}

func returnBook(w http.ResponseWriter, r *http.Request) {
	// Helper function to remove the reservation
	idStr := strings.TrimPrefix(r.URL.Path, "/return/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Book ID type", http.StatusBadRequest)
		return
	}

	// If no error, look for the book id
	for i, book := range books {
		if book.Id == id {
			books[i].Reserved = false
			w.Header().Set("X-Reserved", "Available")
			io.WriteString(w, "Returned: ")
			json.NewEncoder(w).Encode(book.Title)
			return
		}
	}

	// If book id not found
	http.NotFound(w, r)
}
