package server

import (
	"net/http"

	"github.com/aklbnelly/libraryproject/internal/authors"
	"github.com/aklbnelly/libraryproject/internal/books"
	"github.com/aklbnelly/libraryproject/internal/readers"
	"github.com/aklbnelly/libraryproject/utils"
)

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := r.URL.Query().Get("id")
			if id != "" {
				books.GetBookHandler(w, r)
			} else {
				books.GetBooksHandler(w, r)
			}

		} else if r.Method == http.MethodPost {
			books.AddBookHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := r.URL.Query().Get("id")
			if id != "" {
				authors.GetAuthorHandler(w, r)
			} else {
				authors.GetAuthorsHandler(w, r)
			}
		} else if r.Method == http.MethodPost {
			authors.AddAuthorHandler(w, r)
		} else if r.Method == http.MethodPatch {
			authors.ChangeAuthorHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/readers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := r.URL.Query().Get("id")
			if id != "" {
				readers.GetReaderHandler(w, r)
			} else {
				readers.GetReadersHandler(w, r)
			}
		} else if r.Method == http.MethodPatch {
			readers.ChangeReaderHandler(w, r)
		} else if r.Method == http.MethodPost {
			readers.AddReaderHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	utils.Logger.Info("server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		utils.Logger.Fatalf("Error starting server: %v", err)
	}

}
