package jolokiaClient

import (
	"encoding/json"
)

//THE RESPONSE ROOT
type ResponseRoot struct {
	Status    int             `json:"status"`
	Timestamp int             `json:"timestamp"`
	Request   BaseRequest     `json:"request"`
	Value     json.RawMessage `json:"value"`
	History   json.RawMessage `json:"history"`
	ResponseError
}

type ResponseError struct {
	Type       string `json:"error_type"`
	Message    string `json:"error"`
	Stacktrace string `json:"stacktrace"`
}
