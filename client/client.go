package client

import (
	"bytes"
	"encoding/json"
	"github.com/codingchipmunk/JolokiaGo/requests"
	"github.com/codingchipmunk/JolokiaGo/responses"
	"io"
	"io/ioutil"
	"net/http"
)

const contentType = "application/json"

type Client struct {
	url    string
	client *http.Client
	sseID  string
}

func (jc *Client) MakePOSTRequest(request requests.POSTRequest) (resp responses.Root, err error){
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

func (jc *Client) MakeGETRequest(request requests.GETRequest) (resp responses.Root, err error){
	// Create a new Buffer for the url and the get-params
	urlBuff := bytes.Buffer{}
	urlBuff.WriteString(jc.url)
	err = request.AppendRequest(&urlBuff)
	if(err != nil){
		return
	}

	// Use the http client to make the request
	httpResp, err := jc.client.Get(urlBuff.String())
	if err != nil {
		return
	}
	// Immediately defer Body.Close() (idiomatic)
	defer httpResp.Body.Close()

	return unmarshalResponse(httpResp.Body)

}

func unmarshalResponse(responseBody io.ReadCloser) (resp responses.Root, err error) {
	// Read the response body
	httpBody, err := ioutil.ReadAll(responseBody)
	if (err != nil) {
		return
	}

	// Unmarshal the response body into the response.Root struct
	err = json.Unmarshal(httpBody, &resp)

	return
}
