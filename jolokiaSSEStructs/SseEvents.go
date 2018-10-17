package jolokiaSSEStructs

import (
	"encoding/json"
	"strconv"
	"time"
)

type SSERoot struct {
	Dropped       int               `json:"dropped"`
	Handle        string            `json:"handle"`
	Handback      string            `json:"handback"`
	Notifications []SSENotification `json:"notifications"`
}

type SSENotification struct {
	TimeStamp      IntTimeWrapper  `json:"timeStamp"`
	SequenceNumber int             `json:"sequenceNumber"`
	UserData       json.RawMessage `json:"userData"`
	Source         string          `json:"source"`
	Message        string          `json:"message"`
	Type           string          `json:"type"`
}

type IntTimeWrapper struct {
	time.Time
}

func (cT *IntTimeWrapper) UnmarshalJSON(bs []byte) error {
	st := string(bs)
	decPoint := len(st) - 3
	in, err := strconv.Atoi(st[:decPoint])
	if err != nil {
		return err
	}
	is, err := strconv.Atoi(st[decPoint:])
	if err != nil {
		return err
	}
	secs := int64(in)
	nanos := int64(is) * 1000000

	cT.Time = time.Unix(secs, nanos)
	return nil
}