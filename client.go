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

// MakePOSTRequest makes an POST request to the Jolokia agent using the http.Client given to the client struct
func (jc *Client) MakePOSTRequest(request messages.POSTRequest) (resp messages.ResponseRoot, err error) {
	return backend.MakePOSTRequest(jc.url, jc.client, request)
}

// MakeGETRequest makes an GET request to the Jolokia agent using the http.Client given to the client struct
func (jc *Client) MakeGETRequest(request messages.GETRequest) (resp messages.ResponseRoot, err error) {
	return backend.MakeGETRequest(jc.url,jc.client,request)

}


