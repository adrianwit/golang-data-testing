package viant

import (
	"testing"
	"github.com/viant/toolbox"
	"path"
	"github.com/stretchr/testify/assert"
	"github.com/viant/assertly"
)

func TestNewConfigFromURL(t *testing.T) {
	type usecase struct {
		description string
		configURL   string
		expected    *Config
		hasError    bool
	}
	var useCases []*usecase

	var parentDirectory = toolbox.CallerDirectory(3)
	useCases = append([]*usecase(nil),
		&usecase{
			description: "loading config from json",
			configURL:   path.Join(parentDirectory, "test", "config.json"),
			expected: &Config{
				Endpoint: &Endpoint{
					Port:      8080,
					TimeoutMs: 2000,
				},
				LogTypes: map[string]*LogType{
					"type1": {
						Locations: []*Location{{
							URL: "file:///data/log/type1",
						}},
						MaxQueueSize:       2048,
						QueueFlashCount:    1024,
						FlushFrequencyInMs: 500,
					},
					"type2": {
						Locations: []*Location{{
							URL: "file:///data/log/type2",
						}},
						MaxQueueSize:       4096,
						QueueFlashCount:    2048,
						FlushFrequencyInMs: 1000,
					},
				},
			},
		})

	useCases = append(useCases,
		&usecase{
			description: "loading config from yaml",
			configURL:   path.Join(parentDirectory, "test", "config.yaml"),
			expected: &Config{
				Endpoint: &Endpoint{
					Port:      8080,
					TimeoutMs: 2000,
				},
				LogTypes: map[string]*LogType{
					"type1": {
						Locations: []*Location{{
							URL: "file:///data/log/type1",
						}},
						MaxQueueSize:       2048,
						QueueFlashCount:    1024,
						FlushFrequencyInMs: 500,
					},
					"type2": {
						Locations: []*Location{{
							URL: "file:///data/log/type2",
						}},
						MaxQueueSize:       4096,
						QueueFlashCount:    2048,
						FlushFrequencyInMs: 1000,
					},
				},
			},
		})

	useCases = append(useCases,
		&usecase{
			description: "error with invalid URL",
			configURL:   path.Join(parentDirectory, "test", "c.yaml"),
			hasError:true,
		})



	for _, useCase := range useCases {
		config, err := NewConfigFromURL(useCase.configURL)
		if useCase.hasError {
			if assert.NotNil(t, err, useCase.description) {
				continue
			}
		}
		if ! assert.Nil(t, err, useCase.description) {
			continue
		}
		assert.Equal(t, useCase.expected, config, useCase.description)
	}
}


func TestNewConfigFromURL_v2(t *testing.T) {
	type usecase struct {
		description string
		configURL   string
		expected    interface{}
		hasError    bool
	}
	var useCases []*usecase

	var parentDirectory = toolbox.CallerDirectory(3)
	useCases = append([]*usecase(nil),
		&usecase{
			description: "loading config from json",
			configURL:   path.Join(parentDirectory, "test", "config.json"),
			expected: `{
	  "Endpoint": {
		"Port": 8080,
		"TimeoutMs": 2000
	  },
	  "LogTypes": {
		"type1": {
		  "Locations":[
			{
			  "URL":"~/type1/"
			}
		  ],
		  "MaxQueueSize": 2048,
		  "QueueFlashCount": 1024,
		  "FlushFrequencyInMs": 500
		},
		"type2": "@exists@" 
	  }
	}`,
		},


		)




	for _, useCase := range useCases {
		config, err := NewConfigFromURL(useCase.configURL)
		if useCase.hasError {
			if assert.NotNil(t, err, useCase.description) {
				continue
			}
		}
		if ! assert.Nil(t, err, useCase.description) {
			continue
		}
		assertly.AssertValues(t, useCase.expected, config, useCase.description)
	}
}