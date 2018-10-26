package requests

import (
	"github.com/codingchipmunk/JolokiaGo"
)

//	Base represents the fields which will always be required in a request to Jolokia.
//	Type should always be given, while Command or MBean may be left empty. If empty they will not be marshalled into JSON.
type Base struct {
	Type    string `json:"type"`
	Command string `json:"command,omitempty"`
	MBean   *JolokiaGo.MBean `json:"mbean,omitempty"`
}

//	RegisterEvent represents the fields needed for a request to register a client identified by the ClientID field for JMX Notifications on a MBean.
type RegisterEvent struct {
	Base
	Mode     string `json:"mode"`
	ClientID string `json:"client"`
}