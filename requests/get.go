package requests

import "bytes"

// GETRequest should be implemented by request-structs to support being send via get
type GETRequest interface {
	// GetAppendix returns the string which should be concatenated with the base url of Jolokia to make the request
	GetAppendix() ([]byte, error)
	// AppendRequest appends the get-request to the buffer
	AppendRequest(buffer *bytes.Buffer) error
}
