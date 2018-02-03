package request

import (
	"time"
)

type PlaybackControllerRequest struct {
	Version string  `json:"version"`
	Context Context `json:context`
	Request struct {
		Timestamp time.Time `json:"timestamp, string"`
		ReqType   string    `json:"type"`
		RequestId string    `json:"requestId"`
		Locale    string    `json:"locale"`
	} `json:request`
}
