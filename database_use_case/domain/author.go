package domain

type Author struct {
	ID   int
	Name string
}

func NewAuthor(id int, name string) *Author {
	return &Author{
		ID:   id,
		Name: name,
	}
}
