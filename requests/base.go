package requests

import (
	"github.com/codingchipmunk/JolokiaGo"
)

//	Base represents the fields which will always be required in a request to Jolokia.
//	Type should always be given, while Command or MBean may be left empty. If empty they will not be marshalled into JSON.
type Base struct {
	//Type of the request. It is basically always required
	Type string `json:"type"`
	//Command which is needed for some requests. It won't be marshaled if left empty.
	Command string `json:"command,omitempty"`
	//MBean which is needed for some requests. It won't be marshaled if left empty.
	MBean *jolokiago.MBean `json:"mbean,omitempty"`
}

//RegisterEvent represents the fields needed for a request to register a client identified by the ClientID field for JMX Notifications on a MBean.
type RegisterEvent struct {
	Base
	// Mode for JMXEvents. Can be SSE or Pull-based.
	Mode string `json:"mode"`
	// ClientID of the Jolokia Client
	ClientID string `json:"client"`
}
