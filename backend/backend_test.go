package backend

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"
)

type rountTripperMock struct {
	share    *shared
	response http.Response
}

type shared struct {
	requests   []httpRequestData
	processing sync.WaitGroup
}

type httpRequestData struct {
	url         string
	contentType string
	body        string
	method      string
}

func (mck rountTripperMock) RoundTrip(req *http.Request) (*http.Response, error) {
	mck.share.processing.Add(1)
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	mck.share.requests = append(mck.share.requests, httpRequestData{url: req.URL.String(), method: req.Method, contentType: req.Header.Get("content-type"), body: string(body)})
	mck.share.processing.Done()
	return &mck.response, nil
}

type mockPOST struct {
	body        string
	contentType string
}

func (m mockPOST) POSTBody() ([]byte, error) {
	return []byte(m.body), nil
}

func (m mockPOST) ContentType() string {
	return m.contentType
}

func TestMakePOSTRequest(t *testing.T) {
	const body = "testing body"
	const cType = "testing/contenttype"
	const url = "test.url"
	reqPoint := shared{}
	response := http.Response{Body: ioutil.NopCloser(strings.NewReader("{}"))}
	mockClient := http.Client{Transport: rountTripperMock{response: response, share: &reqPoint}}

	_, err := MakePOSTRequest(url, &mockClient, mockPOST{body, cType})
	if err != nil {
		t.Logf("Unexpected Error: %s", err)
		t.Fail()
	}

	reqPoint.processing.Wait()
	if len(reqPoint.requests) != 1 {
		t.Logf("Unexpected number of requests: %d", len(reqPoint.requests))
		t.FailNow()
	}
	request := reqPoint.requests[0]

	if request.contentType != cType {
		t.Logf("Unexpected content type: %s", request.contentType)
		t.Fail()
	}
	if request.method != "POST" {
		t.Logf("Unexpected method: %s", request.method)
		t.Fail()
	}
	if request.body != body {
		t.Logf("Unexpected body: %s", request.body)
		t.Fail()
	}
	if request.url != url {
		t.Logf("Unexpected url: %s", request.url)
		t.Fail()
	}
}
