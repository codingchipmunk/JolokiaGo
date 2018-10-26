package requests

import "bytes"

type GETRequest interface {
	GetAppendix() ([]byte,error)
	AppendRequest(buffer *bytes.Buffer) (error)
}
