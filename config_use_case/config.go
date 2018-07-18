package viant

import "github.com/viant/toolbox/url"

type Location struct {
	URL string
	Credentials string
}

type LogType struct {
	Locations []*Location
	MaxQueueSize       int
	QueueFlashCount    int
	FlushFrequencyInMs int

}

type Endpoint struct {
	Port      int
	TimeoutMs int
}

type Config struct {
	Endpoint *Endpoint
	LogTypes map[string]*LogType
}


func NewConfigFromURL(URL string) (*Config, error) {
	resource := url.NewResource(URL)
	config := &Config{}
	return config,resource.Decode(config)
}