package messages

// GETRequest should be implemented by request-structs to support being send via get
type GETRequest interface {
	// GETAppendix returns the string which should be concatenated with the base url of Jolokia to make the request
	GETAppendix() ([]byte, error)
}
