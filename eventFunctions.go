package jolokiago

import (
	"net/http"

	"github.com/codingchipmunk/jolokiago/events"
	"github.com/codingchipmunk/jolokiago/messages"
	"github.com/codingchipmunk/jolokiago/messages/registerEvent"
)

func (jc *Client) NewSSEListener(client *http.Client) (listener *events.Listener, err error) {
	resp, err := jc.getEventClientID()
	if err != nil{
		return
	}

	return events.NewListener(jc.url, resp.ID, client), err
}

func (jc *Client) getEventClientID() (retVal registerEvent.ResponseValue, err error) {
	req := messages.BaseRequest{Type: "notification", Command: "register"}
	ret, err := jc.MakePOSTRequest(req)
	if err != nil {
		return
	}

	return registerEvent.DecodeResponseValue(ret.Value)
}
