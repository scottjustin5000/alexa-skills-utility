package request

import (
	"time"
)

type LaunchRequest struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Context Context `json:context`
	Request struct {
		Timestamp time.Time `json:"timestamp, string"`
		ReqType   string    `json:"type"`
		RequestId string    `json:"requestId"`
		Locale    string    `json:"locale"`
	} `json:request`
}
