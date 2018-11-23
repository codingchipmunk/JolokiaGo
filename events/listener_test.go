package events

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/codingchipmunk/jolokiago/backend"
	"github.com/codingchipmunk/jolokiago/java"
	"github.com/codingchipmunk/jolokiago/messages"
	"github.com/codingchipmunk/jolokiago/messages/registerEvent"
)

func TestListener_StartListening(t *testing.T) {
	const base = "http://localhost:8778/jolokia/"
	idReq, err := getSSEID(base)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	l := NewListener(base, idReq.ID, http.DefaultClient)
	fmt.Println("New listener acquired!")
	err = l.SubscribeToMBean(java.MBean{Type: "tracer", Context: "pop-os/camel-1", Name: "Tracer", Domain: "your.domain.name"})
	if err != nil{
		t.Error(err)
	}
	fmt.Println("Subscribed!")
	l.StartListening()
	fmt.Println("Listening")
	for l.started {
	}
}

func getSSEID(url string) (registerEvent.ResponseValue, error) {
	req := messages.BaseRequest{Type: "notification", Command: "register"}
	ret, err := backend.MakePOSTRequest(url, http.DefaultClient, req)
	if err != nil {
		return registerEvent.ResponseValue{}, err
	}

	return registerEvent.DecodeResponseValue(ret.Value)
}
