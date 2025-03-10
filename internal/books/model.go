package books

type Book struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	IsbnCode int    `json:"isbnCode"`
	AuthorId int    `json:"authorId"`
}

type newBook struct {
	Title      string `json:"title"`
	Genre      string `json:"genre"`
	IsbnCode   int    `json:"isbnCode"`
	AuthorName string `json:"authorName"`
}
