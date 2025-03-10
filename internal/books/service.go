package books

import "github.com/aklbnelly/libraryproject/internal/authors"

func GetBooks() ([]Book, error) {

	return GetAllBooks()

}

func GetBookService(bookId int) (Book, error) {
	return GetBookById(bookId)
}

func AddBookService(book newBook) (Book, error) {
	author, err := authors.GetAuthorById(book.AuthorId)

}
