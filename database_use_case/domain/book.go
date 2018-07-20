package domain

type Book struct {
	ID      int
	ISBN    string
	Title   string
	Summary string
	Author  *Author
}

func NewBook(id int, ISBN string, title string, summary string, author *Author) *Book {
	return &Book{
		ID:      id,
		ISBN:    ISBN,
		Title:   title,
		Summary: summary,
		Author:  author,
	}
}
