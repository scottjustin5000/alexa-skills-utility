package response

type AudioStream struct {
	Url                   string `json:"url"`
	Token                 string `json:"token"`
	ExpectedPreviousToken string `json:"expectedPreviousToken"`
	OffsetInMilliseconds  int64  `json:"offsetInMilliseconds"`
}

type AudioItem struct {
	Stream AudioStream `json:"stream"`
}

type PlayDirective struct {
	Type         string `json:"type"`
	PlayBehavior string `json:playBehavior`
	Request      struct {
		Timestamp            time.Time `json:"timestamp, string"`
		ReqType              string    `json:"type"`
		RequestId            string    `json:"requestId"`
		Locale               string    `json:"locale"`
		Token                string    `json:"token"`
		OffsetInMilliseconds int64     `json:"offsetInMilliseconds"`
	} `json:request`
}
