package messages

import "encoding/json"

// POSTRequest should be implemented by request-structs to support being send via post.
// Since Jolokia expects a JSON as POST it's easy to simply use json.Marshal to marshal the struct
type POSTRequest interface {
	// Returns the body of the POST request which will be send to Jolokia
	POSTBody() ([]byte, error)
	GetContentType() string
}

// Simply calls and returns the values of json.Marshal() with the given interface
// Used internally to not pollute import statements
func SimplePOSTImpl(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}
