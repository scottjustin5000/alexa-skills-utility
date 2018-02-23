package request

import (
	"time"
)

type SessionEndedRequest struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Context Context `json:context`
	Request struct {
		Timestamp time.Time `json:"timestamp, string"`
		ReqType   string    `json:"type"`
		RequestId string    `json:"requestId"`
		Reason    string    `json:"reason"`
		Locale    string    `json:"string"`
		Error     struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		} `json:"error"`
	} `json:"request"`
}
