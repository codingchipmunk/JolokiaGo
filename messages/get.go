package messages

// GETRequest should be implemented by request-structs to support being send via get
type GETRequest interface {
	// GetAppendix returns the string which should be concatenated with the base url of Jolokia to make the request
	GetAppendix() ([]byte, error)
}