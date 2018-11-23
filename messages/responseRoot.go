package messages

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//	ResponseRoot represents the root of the JSON Response. Value and History are not unmarshaled since their type will vary from request to request.
//	To unmarshal Value and History fields define an own struct and unmarshal the fields accordingly.
type ResponseRoot struct {
	Status    int             `json:"status"`
	Timestamp int             `json:"timestamp"`
	Request   BaseRequest     `json:"request"`
	Value     json.RawMessage `json:"value"`
	History   json.RawMessage `json:"history"`
	ResponseError
}

func (respRoot ResponseRoot) Successful() bool {
	return respRoot.Status == http.StatusOK
}


//	ResponseError contains fields related to internal erros in Jolokia.
type ResponseError struct {
	Type       string `json:"error_type"`
	Message    string `json:"error"`
	Stacktrace string `json:"stacktrace"`
}

func (respErr ResponseError) Error() string {
	return fmt.Sprintf("Error in Jolokia: Type: %s ; Message: %s ; Stacktrace: %s", respErr.Type, respErr.Message, respErr.Stacktrace)
}
