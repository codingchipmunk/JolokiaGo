package events

import (
	"encoding/json"
	"strconv"
	"time"
)

// Root of an JMXEvent exposed by Jolokia
type Root struct {
	Dropped       int            `json:"dropped"`
	Handle        string         `json:"handle"`
	Handback      string         `json:"handback"`
	Notifications []Notification `json:"notifications"`
}

// Notification of the JMX Event
type Notification struct {
	TimeStamp      IntTimeWrapper `json:"timeStamp"`
	SequenceNumber int            `json:"sequenceNumber"`
	// Data of the JMX-event
	UserData json.RawMessage `json:"userData"`
	Source   string          `json:"source"`
	Message  string          `json:"message"`
	Type     string          `json:"type"`
}

// Wrapper for time.Time for easy unmarshaling: time.Times unmarshal function requires a string in time.UnixDate format, but Jolokia returns an int in a string (duh!)
type IntTimeWrapper struct {
	time.Time
}

// Unmarshals an integer into time.Time
// Converts the bytes into string and then uses the last 3 digits as nanos and converts this using the time.Unix() function
// Example:
// The byte representation of the string "12345678" will be interpreted as 12345.678 secs in unix time
func (cT *IntTimeWrapper) UnmarshalJSON(bs []byte) error {
	const decPointOffset = 3                   // how many digits of the string are nanoseconds. Will be used to split the string at len(st) - decPointOffset
	const nanoMult = 10 ^ (9 - decPointOffset) // Needed to convert the nanoseconds correctly

	st := string(bs)                     // String of the []byte. This is needed because Jolokia gives an int in a string meaning it can't be decoded directly into an int
	decPoint := len(st) - decPointOffset // position of the (imagined) decimal point in this string

	sSecs, err := strconv.Atoi(st[:decPoint])
	if err != nil {
		return err
	}
	sNans, err := strconv.Atoi(st[decPoint:])
	if err != nil {
		return err
	}

	secs := int64(sSecs)
	nanos := int64(sNans) * nanoMult

	cT.Time = time.Unix(secs, nanos)
	return nil
}
