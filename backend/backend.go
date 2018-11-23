package backend

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/codingchipmunk/jolokiago/messages"
)

// MakePOSTRequest makes an POST request to the Jolokia agent using the given http.Client
func MakePOSTRequest(url string, client *http.Client, request messages.POSTRequest) (resp messages.ResponseRoot, err error) {
	// Marshal the request
	body, err := request.POSTBody()
	if err != nil {
		return
	}
	// Use the http client to make the request
	httpResp, err := client.Post(url, request.ContentType(), bytes.NewReader(body))
	if err != nil {
		return
	}
	// Immediately defer Body.Close() (idiomatic)
	defer httpResp.Body.Close()

	resp, err = unmarshalResponse(httpResp.Body)
	if err != nil {
		return
	}

	if !resp.Successful() {
		return resp, resp.ResponseError
	}
	return
}

// MakeGETRequest makes an GET request to the Jolokia agent using the given http.Client
func MakeGETRequest(url string, client *http.Client, request messages.GETRequest) (resp messages.ResponseRoot, err error) {
	// Create a new Buffer for the url and the get-params
	urlBuff := bytes.Buffer{}
	urlBuff.WriteString(url)
	bts, err := request.GETAppendix()
	if err != nil {
		return
	}
	urlBuff.Write(bts)
	// Use the http client to make the request
	httpResp, err := client.Get(urlBuff.String())
	if err != nil {
		return
	}
	// Immediately defer Body.Close() (idiomatic)
	defer httpResp.Body.Close()
	resp, err = unmarshalResponse(httpResp.Body)
	if err != nil {
		return
	}

	if !resp.Successful() {
		return resp, resp.ResponseError
	}
	return
}

// unmarshalResponse unmarshals the response from a response body (or any struct implementing the io.ReadCloser interface)
func unmarshalResponse(responseBody io.ReadCloser) (resp messages.ResponseRoot, err error) {
	// Read the response body
	httpBody, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return
	}

	// Unmarshal the response body into the response.ResponseRoot struct
	err = json.Unmarshal(httpBody, &resp)
	if err != nil {
		return
	}

	return
}

func CheckResponseError(resp messages.ResponseRoot) error {
	if resp.Status != http.StatusOK {
		return resp.ResponseError
	}
	return nil
}
