package viant

import (
	"github.com/viant/toolbox/storage"
	"time"
	"bytes"
	"io"
	"github.com/viant/toolbox"
	"sync/atomic"
	"strconv"
	"log"
)

const queueSize = 4086
var syncTimeout = 100 * time.Millisecond


type Logger interface {
	Log(message *Message)
	Sync()
	Close()
}


type logger struct {
	encoderFactory toolbox.EncoderFactory
	messages chan *Message
	dest     storage.Service
	baseURL  string
	flushCounter uint64
	queueSize int
	closed chan bool
	flushFrequencyInMs int
}


func (l *logger) Sync() {
	var lastFlashTime = time.Now()
	var flashFrequency = time.Millisecond * time.Duration(l.flushFrequencyInMs)
	var buf = new(bytes.Buffer)
	encoder := l.encoderFactory.Create(buf)
	var count = 0
	for ; ; {
		select {
		case <-l.closed:
			return
		case <-time.After(syncTimeout):
			elapsed := time.Now().Sub(lastFlashTime)
			if elapsed < flashFrequency || count == 0 {
				continue
			}
			lastFlashTime = time.Now()
			l.flush(buf)
			count = 0
			buf = new(bytes.Buffer)
			encoder = l.encoderFactory.Create(buf)

		case message := <-l.messages:
			count++
			encoder.Encode(message)
		}
	}
}
func (l *logger) flush(reader io.Reader)  {
	flushCounter := atomic.LoadUint64(&l.flushCounter)
	URL := l.baseURL
	if flushCounter > 0 {
		URL += strconv.Itoa(int(flushCounter))
	}
	if err := l.dest.Upload(URL, reader);err != nil {
		log.Fatal(err)
	}
	atomic.AddUint64(&l.flushCounter, 1)
}

func (l *logger) Close() {
	l.closed <- true
}

func (l *logger) Log(message *Message) {
	l.messages <- message
}


func New(dest storage.Service, baseURL string, flushFrequencyInMs int) Logger {
	return &logger{
		queueSize: queueSize,
		dest:     dest,
		baseURL:  baseURL,
		encoderFactory : toolbox.NewJSONEncoderFactory(),
		flushFrequencyInMs: flushFrequencyInMs,
		closed: make(chan bool, 1),
		messages: make(chan *Message, queueSize),
	}
}
