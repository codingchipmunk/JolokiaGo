package jolokiago

import (
	"net/http"

	"github.com/codingchipmunk/jolokiago/backend"
	"github.com/codingchipmunk/jolokiago/messages"
)

const contentType = "application/json"

// Client holds fields needed to communicate with a Jolokia agent
type Client struct {
	url    string
	client *http.Client
}

func New(url string, client *http.Client) Client{
	return Client{url,client}
}

// MakePOSTRequest makes an POST request to the Jolokia agent using the http.Client given to the client struct
func (jc *Client) MakePOSTRequest(request messages.POSTRequest) (resp messages.ResponseRoot, err error) {
	resp, err = backend.MakePOSTRequest(jc.url, jc.client, request)
	if err != nil {
		return
	}
	err = backend.CheckResponseError(resp)
	return
}

// MakeGETRequest makes an GET request to the Jolokia agent using the http.Client given to the client struct
func (jc *Client) MakeGETRequest(request messages.GETRequest) (resp messages.ResponseRoot, err error) {
	resp, err = backend.MakeGETRequest(jc.url, jc.client, request)
	if err != nil {
		return
	}
	err = backend.CheckResponseError(resp)
	return
}
