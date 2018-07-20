package viant

import (
	"testing"
	"github.com/viant/dsunit"
	_ "github.com/mattn/go-sqlite3"

	"github.com/viant/dsc"
	"github.com/viant/toolbox"
	"path"
	"github.com/stretchr/testify/assert"
	"github.com/viant/assertly"
	"github.com/adrianwit/golang-data-testing/database_use_case/domain"
)

func searchTestSetup(t *testing.T) bool {
	if ! dsunit.InitFromURL(t, "test/init.yaml") {
		return false
	}
	if !dsunit.PrepareFor(t, "library", "test/data/", "search") {
		return false
	}
	return true
}



func getDbManagerForURL(URL string) (dsc.Manager, error) {
	var parentDirectory = toolbox.CallerDirectory(4)
	conifg, err := dsc.NewConfigFromURL(path.Join(parentDirectory, URL))
	if err != nil {
		return nil, err
	}
	return dsc.NewManagerFactory().Create(conifg)
}



func TestService_Search(t *testing.T) {
	if !searchTestSetup(t) {
		return
	}
	db, err := getDbManagerForURL("test/config.yaml")
	if ! assert.Nil(t, err) {
		return
	}
	service := New(db)

	type usecase struct {
		description string
		request     *SearchRequest
		expected    interface{}
	}
	var useCases = []*usecase(nil)

	useCases = append(useCases, &usecase{
		description: "search by ISBN",
		request: &SearchRequest{
			By:    "isbn",
			Value: "978-3-16-148410-1",
		},
		expected:
		&SearchResponse{
			BaseResponse: &BaseResponse{Status: "ok",},
			Books: []*domain.Book{domain.NewBook(2, "978-3-16-148410-1", "title 2", "",
					domain.NewAuthor(1, "Dudi")),
			},
		}})

	useCases = append(useCases, &usecase{
		description: "search by title",
		request: &SearchRequest{
			By:    "title",
			Value: "title 3",
		},
		expected:
		&SearchResponse{
			BaseResponse: &BaseResponse{Status: "ok",},
			Books: []*domain.Book{domain.NewBook(3, "978-3-16-148410-2", "title 3", "",
				domain.NewAuthor(2, "Vidi")),
			},
		}})

	for _, useCase := range useCases {
		response := service.Search(useCase.request)
		assertly.AssertValues(t, useCase.expected, response, useCase.description)
	}
}



func Test_searchTestSetup(t *testing.T) {
	if !searchTestSetup(t) {
		return
	}

	db, err := getDbManagerForURL("test/config.yaml")
	if ! assert.Nil(t, err) {
		return
	}
	var data= []map[string]interface{}{}
	err = db.ReadAll(&data, BooksDQL, nil, nil)
	if ! assert.Nil(t, err) {
		return
	}
	expected := `
  { "id":1, "isbn":"978-3-16-148410-0", "title":"title 1", "author_id":1},
  { "id":2,  "isbn":"978-3-16-148410-1", "title":"title 2", "author_id":1},
  { "id":3,  "isbn":"978-3-16-148410-2",  "title":"title 3", "author_id":2}
`
	assertly.AssertValues(t, expected, data, "validating setup")

	}

