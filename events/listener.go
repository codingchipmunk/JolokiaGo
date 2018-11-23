package events

import (
	"net/http"
	"strings"

	"github.com/codingchipmunk/jolokiago/backend"
	"github.com/codingchipmunk/jolokiago/java"
	"github.com/codingchipmunk/jolokiago/messages"
	"github.com/codingchipmunk/jolokiago/messages/registerEvent"
	"github.com/r3labs/sse"
)

const notificationPath = "/notification/open/"
const ssePath = "/sse"

type Listener struct {
	id         string
	baseURL    string
	httpClient *http.Client
	sseClient  *sse.Client
	rawChannel chan *sse.Event
	started    bool
}

func NewListener(baseURL string, id string, httpClient *http.Client) (*Listener) {
	var lst Listener
	lst.id = id
	lst.baseURL = baseURL
	lst.httpClient = httpClient
	lst.sseClient = sse.NewClient(strings.TrimRight(baseURL, "/") + notificationPath + id + ssePath)
	lst.started = false
	return &lst
}

func (ls *Listener) ClientID() string {
	return ls.id
}

func (ls *Listener) JolokiaBaseURL() string {
	return ls.baseURL
}

func (ls *Listener) FullURL() string {
	return ls.sseClient.URL
}

func (ls *Listener) RawEvents() <-chan *sse.Event {
	return ls.rawChannel
}

func (ls *Listener) StartListening() (<-chan *sse.Event, bool) {
	if ls.started {
		return ls.rawChannel, false
	}
	ls.rawChannel = make(chan *sse.Event)
	go ls.sseClient.SubscribeChanRaw(ls.rawChannel)
	ls.started = true
	return ls.rawChannel, true
}

func (ls *Listener) IsRunning() bool {
	return ls.started
}

func (ls *Listener) Stop() bool {
	if !ls.started {
		return false
	}
	close(ls.rawChannel)
	go ls.sseClient.Unsubscribe(ls.rawChannel)
	ls.started = false
	return true
}

func (ls *Listener) SubscribeToMBean(bean java.MBean) error {
	_, err := backend.MakePOSTRequest(ls.baseURL, ls.httpClient, registerEvent.Request{ClientID: ls.id, Mode: "sse", BaseRequest: messages.BaseRequest{Type: "notification", Command: "add", MBean: bean}})
	if err != nil {
		return err
	}
	return nil
}
