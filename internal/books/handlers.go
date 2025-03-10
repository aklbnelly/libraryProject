package books

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	AuthorId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid book Id", http.StatusBadRequest)
		return
	}
	author, err := GetBookService(AuthorId)
	if err != nil {
		if err.Error() == "no books found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var newBook NewBook

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	addedBook, err := AddBookService(newBook)
	if err != nil {
		http.Error(w, "Failed to add book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(addedBook)
}
