package events

import (
	"fmt"
	"net/http"

	"github.com/codingchipmunk/jolokiago/backend"
	"github.com/codingchipmunk/jolokiago/java"
	"github.com/codingchipmunk/jolokiago/messages"
	"github.com/codingchipmunk/jolokiago/messages/registerEvent"
	"github.com/r3labs/sse"
)

const sseClientURL = "%s/notification/open/%s/sse"

type Listener struct {
	id         string
	baseURL    string
	httpClient *http.Client
	sseClient  *sse.Client
	rawChannel chan *sse.Event
	started    bool
}

func New(baseURL string, id string, httpClient *http.Client) (lst *Listener) {
	lst.id = id
	lst.baseURL = baseURL
	lst.httpClient = httpClient
	lst.sseClient = sse.NewClient(fmt.Sprintf(sseClientURL, baseURL, id))
	lst.rawChannel = make(chan *sse.Event)
	return
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

func (ls *Listener) Start() error {
	if ls.started {
		return nil
	}
	go ls.sseClient.SubscribeChanRaw(ls.rawChannel)
	return nil
}

func (ls *Listener) IsRunning() bool{
	return ls.started
}

func (ls *Listener) Stop() error {
	if !ls.started {
		return nil
	}
	ls.sseClient.Unsubscribe(ls.rawChannel)
	return nil
}

func (ls *Listener) SubscribeToMBean(bean java.MBean) error {
	_, err := backend.MakePOSTRequest(ls.baseURL, ls.httpClient, registerEvent.Request{ClientID: ls.id, Mode: "sse", BaseRequest: messages.BaseRequest{Type: "notification", Command: "add", MBean: bean}})
	if err != nil {
		return err
	}
	return nil
}
