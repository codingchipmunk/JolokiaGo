package messages

import (
	"github.com/codingchipmunk/jolokiago"
)

//	BaseRequest represents the fields which will always be required in a request to Jolokia.
//	Type should always be given, while Command or MBean may be left empty. If empty they will not be marshalled into JSON.
type BaseRequest struct {
	//Type of the request. It is basically always required
	Type string `json:"type"`
	//Command which is needed for some requests. It won't be marshaled if left empty.
	Command string `json:"command,omitempty"`
	//MBean which is needed for some requests. It won't be marshaled if left empty.
	MBean *jolokiago.MBean `json:"mbean,omitempty"`
}

// Returns a JSON representation of the struct to use as body when making POST requests
// Simply calls SimplePOSTImpl
func (b *BaseRequest) POSTBody() ([]byte, error){
	return SimplePOSTImpl(b)
}
