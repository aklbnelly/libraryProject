package authors

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func GetAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	authors, err := GetAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)

}

// getting author by id
func GetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	AuthorId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid author Id", http.StatusBadRequest)
		return
	}
	author, err := GetAuthorService(AuthorId)
	if err != nil {
		if err.Error() == "author is not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func AddAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var author Author

	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if author.FullName == "" || author.Specialization == "" {
		http.Error(w, "invalid full_name/specialization in request body", http.StatusBadRequest)
		return
	}

	err = AddAuthorService(&author)
	if err != nil {

		http.Error(w, "add author error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

// change author info
func ChangeAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	readerID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid author id", http.StatusBadRequest)
		return
	}
	var updateBody updateAuthorBody

	err = json.NewDecoder(r.Body).Decode(&updateBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if updateBody.FullName == "" || updateBody.Specialization == "" {
		http.Error(w, "Invalid full_name/specialization in request body", http.StatusBadRequest)
		return
	}

	err = UpdateAuthorService(readerID, updateBody.FullName, updateBody.Specialization)
	if err != nil {
		if errors.Is(err, ErrAuthorNotFound) { //проверка на соответствие ошибке!
			http.Error(w, "author is not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update reader info", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Author updated successfully"))
}
