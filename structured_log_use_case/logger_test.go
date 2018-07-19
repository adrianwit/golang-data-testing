package viant

import (
	"testing"
	"github.com/viant/toolbox/storage"
	"github.com/stretchr/testify/assert"
	"github.com/viant/assertly"
	"time"
	"github.com/viant/toolbox/bridge"
	"net/http"
)

func TestLogger_Log(t *testing.T) {
	type usecase struct {
		description string
		URL         string
		messages    []*Message
		expected    interface{}
	}
	var useCases = []*usecase(nil)
	var storageService = storage.NewMemoryService()

	var now = time.Now()

	useCases = append(useCases, &usecase{
		description: "event type 1 log validation",
		messages: []*Message{
			NewMessage(&now, 1, "127.0.0.1",
				bridge.NewHTTPRequest("get", "http://viantint.com/test/event1.1", "", http.Header{
					"Accept":   []string{"text/html,application/xhtml+xm…plication/xml;q=0.9,*/*;q=0.8"},
					"UseAgent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:61.0) Gecko/20100101 Firefox/61.0"},
				})),
			NewMessage(&now, 1, "127.0.0.1",
				bridge.NewHTTPRequest("get", "http://viantint.com/test/event1.2", "", http.Header{
					"Accept":   []string{"text/html,application/xhtml+xm…plication/xml;q=0.9,*/*;q=0.8"},
					"UseAgent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:61.0)"},
				})),
		},
		expected: []interface{}{
			assertly.TimeFormat("", "yyyy-MM-dd hh:mm:ss"),
			NewMessage(&now, 1, "127.0.0.1",
				bridge.NewHTTPRequest("get", "http://viantint.com/test/event1.1", "", http.Header{
					"Accept":   []string{"text/html,application/xhtml+xm…plication/xml;q=0.9,*/*;q=0.8"},
					"UseAgent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:61.0) Gecko/20100101 Firefox/61.0"},
				})),
			NewMessage(&now, 1, "127.0.0.1",
				bridge.NewHTTPRequest("get", "http://viantint.com/test/event1.2", "", http.Header{
					"Accept":   []string{"text/html,application/xhtml+xm…plication/xml;q=0.9,*/*;q=0.8"},
					"UseAgent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:61.0)"},
				})),
		},
	})


	var testTime  = time.Now()
	useCases = append(useCases, &usecase{
		description: "event type 1 log validation",
		messages: []*Message{
			NewMessage(&testTime, 1, "127.0.0.1",
				bridge.NewHTTPRequest("get", "http://viantint.com/test/event1.1", "", nil)),
			NewMessage(&testTime, 1, "127.0.0.1",
				bridge.NewHTTPRequest("get", "http://viantint.com/test/event1.2", "", nil)),
		},
		expected: `[
  {
    "Timestamp": "<ds:within_sec[\"now\", 6, \"yyyy-MM-ddTHH:mm:sszz:zz\"]>",
    "EventTypeId": 1,
    "IP": "127.0.0.1",
    "Request": {
      "Method": "get",
      "URL": "http://viantint.com/test/event1.1"
    }
  },
  {
    "Timestamp": "<ds:within_sec[\"now\", 3, \"yyyy-MM-ddTHH:mm:sszz:zz\"]>",
    "EventTypeId": 1,
    "IP": "127.0.0.1",
    "Request": {
      "Method": "get",
      "URL": "http://viantint.com/test/event1.2"
    }
  }
]
`,
	})


	for _, useCase := range useCases {
		logger := New(storageService, useCase.URL, 1000)
		defer logger.Close()
		go logger.Sync()
		for _, message := range useCase.messages {
			logger.Log(message)
		}
		time.Sleep(1100 * time.Millisecond)
		logged, err := storage.DownloadText(storageService, useCase.URL)
		if ! assert.Nil(t, err, useCase.description) {
			continue
		}
		assertly.AssertValues(t, useCase.expected, logged)
	}
}
