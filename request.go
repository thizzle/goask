package goask

import (
	"strconv"
	"time"
)

type Timestamp time.Time

type SessionAttributes map[string]interface{}

type AlexaSession struct {
	New         bool
	SessionID   string
	Attributes  SessionAttributes
	Application struct {
		ApplicationID string
	}
	User struct {
		UserID, AccessToken string
	}
}

type AlexaRequest struct {
	Type      string
	Timestamp Timestamp
	RequestID string
	Intent    struct {
		Name  string
		Slots map[string]struct {
			Name, Value string
		}
	}
}

type AlexaRequestContainer struct {
	Version string
	Session AlexaSession
	Request AlexaRequest
}

// UnmarshalJSON converts the string representation of a Unix timestamp
// (which is how Alexa formats timestamp values) to our Timestamp type
func (t Timestamp) UnmarshalJSON(b []byte) (err error) {
	tsUnix, _ := strconv.ParseInt(string(b), 10, 0)
	t = Timestamp(time.Unix(tsUnix, 0))
	return
}
