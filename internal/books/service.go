package books

import (
	"errors"
	"fmt"

	"github.com/aklbnelly/libraryproject/internal/authors"
)

func GetBooks() ([]Book, error) {
	return GetAllBooks()
}

func GetBookService(bookId int) (Book, error) {
	return GetBookById(bookId)
}

func AddBookService(newBook NewBook) (Book, error) {
	var authorId int

	// Если передан `authorId`, проверяем его существование
	if newBook.AuthorId != nil {
		_, err := authors.GetAuthorService(*newBook.AuthorId)
		if err != nil {
			if errors.Is(err, authors.ErrAuthorNotFound) {
				return Book{}, errors.New("author with is not found")
			}
			return Book{}, errors.New("failed to check author")
		}
		authorId = *newBook.AuthorId
	} else if newBook.AuthorName != nil {
		newAuthor := authors.Author{
			FullName:       *newBook.AuthorName,
			Specialization: "no info",
		}

		err := authors.AddAuthorService(&newAuthor)
		if err != nil {
			return Book{}, errors.New("failed to add author")
		}

		authorId = newAuthor.Id
	} else {
		return Book{}, errors.New("either authorId or authorName must be provided")
	}

	book := Book{
		Title:    newBook.Title,
		Genre:    newBook.Genre,
		IsbnCode: newBook.IsbnCode,
		AuthorId: authorId,
	}

	createdBook, err := AddBook(book)
	if err != nil {
		fmt.Print(err) //убрать
		return Book{}, errors.New("failed to add book")

	}

	return createdBook, nil
}
