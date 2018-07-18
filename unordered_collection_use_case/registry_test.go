package viant

import (
	"testing"
	"github.com/viant/assertly"
)

func TestRegistry_AsSlice(t *testing.T) {

	type usecase struct {
		description string
		messages    []*Message
		expected    interface{}
	}
	var useCases = []*usecase(nil)

	useCases = append(useCases , &usecase{
		description:"overriding message",
		messages:[]*Message{
			NewMessage(1, "dummy 1"),
			NewMessage(2, "name 2"),
			NewMessage(1, "name 1"),
		},
		expected:`[
	  {
		"@indexBy@":"id"
	  },
	  {
		"ID":1,
		"Name":"name 1"
	  },
	  {
		"ID":2,
		"Name":"name 2"
	  }
	]
`,
	})

	for _, useCase := range useCases {
		var registry Registry = map[int]*Message{}
		for _, message := range useCase.messages {
			registry.Register(message)
		}
		assertly.AssertValues(t, useCase.expected, registry.AsSlice())
	}
}
