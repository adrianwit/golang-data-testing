package viant

import (
	"testing"
	"github.com/viant/assertly"
)

func TestService_Predict(t *testing.T) {
	type usecase struct {
		description string
		request     *PredicationRequest
		expected    interface{}
	}
	var useCases = []*usecase(nil)

	useCases = append(useCases,
		&usecase{
			description:"random range prediction basic use case",
			request:&PredicationRequest{},
			expected:`{
  "@switchCaseBy@":["ExperimentID"],
  "1":{
    "Value1":"/[0.11..0.44]/",
    "Value2":"/[0.43..0.84]/",
    "ValueN":"/[0.01..0.3]/"
  },
  "2":{
    "Value1":"/[0.22..0.5]/",
    "Value2": "/[0.44..0.70]/",
    "ValueN":"/[0.2..0.4]/"
  },
  "shared": {
    "Static1":"value1",
    "Static2":"value2",
    "StaticN":"valueN"
  }
}`,
		})



	var service = New()
	for _, useCase := range useCases {
		response := service.Predict(useCase.request)
		assertly.AssertValues(t, useCase.expected, response, useCase.description)
	}
}
