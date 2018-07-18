package viant

type Message struct {
	ID int
	Name string
}

func NewMessage(id int, name string) *Message{
	return  &Message{
		ID:id,
		Name:name,
	}
}
