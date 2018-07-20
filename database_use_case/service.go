package viant

import (
	"github.com/viant/dsc"
	"github.com/adrianwit/golang-data-testing/database_use_case/domain"
)


const BooksDQL = "SELECT  id, isbn, title, summary, author_id FROM books"
const AuthorsDQL = "SELECT  id, name FROM authors"

type Service interface {
	Search(request *SearchRequest) *SearchResponse
}

type SearchRequest struct {
	By string
	Value interface{}
}


type BaseResponse struct {
	Status string
	Error string
}

type SearchResponse struct {
	*BaseResponse
	Books []*domain.Book
}


type service struct {
	manager dsc.Manager
}


func (s *service) Search(request *SearchRequest) *SearchResponse {
	DQL := BooksDQL
	var response = &SearchResponse{
		BaseResponse:&BaseResponse{Status:"ok"},
		Books:[]*domain.Book(nil),
	}
	var params = []interface{}(nil)
	if request.By != "" {
		params = append(params, request.Value)
		DQL += " WHERE " + request.By + " = ?"
	}
	var authorCache = make(map[int]*domain.Author)
	err := s.manager.ReadAllWithHandler(DQL, params, func(scanner dsc.Scanner) (toContinue bool, err error) {
		var book = &domain.Book{
			Author:&domain.Author{},
		}
		var summary  = &book.Summary
		if err := scanner.Scan(&book.ID, &book.ISBN, &book.Title, &summary, &book.Author.ID);err != nil {
			return false, err
		}
		if summary != nil {
			book.Summary = *summary
		}

		if author, ok := authorCache[book.Author.ID];ok {
			book.Author = author
		} else {
			if _, err := s.manager.ReadSingle(book.Author, AuthorsDQL+" WHERE id = ?", []interface{}{book.Author.ID}, nil); err != nil {
				return false, err
			}
			authorCache[book.Author.ID] = book.Author
		}
		response.Books = append(response.Books, book)
		return true, nil
	})
	if err != nil {
		response.Error = err.Error()
		response.Status = "error"
	}
	return response
}

func New(manager dsc.Manager) Service {
	return &service{manager: manager}
}