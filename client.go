package jolokiago

import (
	"bytes"
	"encoding/json"
	"github.com/codingchipmunk/jolokiago/messages"
	"github.com/codingchipmunk/jolokiago/responses"
	"io"
	"io/ioutil"
	"net/http"
)

const contentType = "application/json"

// Client holds fields needed to communicate with a Jolokia agent
type Client struct {
	url    string
	client *http.Client
	sseID  string
}

// MakePOSTRequest makes an POST request to the Jolokia agent using the http.Client given to the client struct
func (jc *Client) MakePOSTRequest(request messages.POSTRequest) (resp responses.Root, err error) {
	// Marshal the request
	body, err := request.POSTBody()
	if err != nil {
		return
	}

	// Use the http client to make the request
	httpResp, err := jc.client.Post(jc.url, contentType, bytes.NewReader(body))
	if err != nil {
		return
	}
	// Immediately defer Body.Close() (idiomatic)
	defer httpResp.Body.Close()

	return unmarshalResponse(httpResp.Body)
}

// MakeGETRequest makes an GET request to the Jolokia agent using the http.Client given to the client struct
func (jc *Client) MakeGETRequest(request messages.GETRequest) (resp responses.Root, err error) {
	// Create a new Buffer for the url and the get-params
	urlBuff := bytes.Buffer{}
	urlBuff.WriteString(jc.url)
	bts, err := request.GetAppendix()
	if err != nil {
		return
	}
	urlBuff.Write(bts)
	// Use the http client to make the request
	httpResp, err := jc.client.Get(urlBuff.String())
	if err != nil {
		return
	}
	// Immediately defer Body.Close() (idiomatic)
	defer httpResp.Body.Close()

	return unmarshalResponse(httpResp.Body)

}

// unmarshalResponse unmarshals the response from a response body (or any struct implementing the io.ReadCloser interface)
func unmarshalResponse(responseBody io.ReadCloser) (resp responses.Root, err error) {
	// Read the response body
	httpBody, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return
	}

	// Unmarshal the response body into the response.Root struct
	err = json.Unmarshal(httpBody, &resp)

	return
}
