package registerEvent

import "github.com/codingchipmunk/jolokiago/messages"

//Request represents the fields needed for a request to register a client identified by the ClientID field for JMX Notifications on a MBean.
//The value of the response is represented in responses.ResponseValue
type Request struct {
	messages.BaseRequest
	// Mode for JMXEvents. Can be SSE or Pull-based.
	Mode string `json:"mode"`
	// ClientID of the Jolokia Client
	ClientID string `json:"client"`
}

// Returns a JSON representation of the struct to use as body when making POST requests
// Simply calls SimplePOSTImpl
func (re *Request) POSTBody() ([]byte, error) {
	return messages.SimplePOSTImpl(re)
}
