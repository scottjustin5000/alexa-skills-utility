package response

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"type"`
	Ssml string `json:"ssml"`
}

type Image struct {
	SmallImageUrl string `json:"smallImageUrl"`
	LargeImageUrl string `json:"largeImageUrl"`
}

type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Text    string `json:"text"`
	Images  Image  `json:"image"`
}

type Reprompt struct {
	OutputSpeech `json:"outputSpeech"`
}

type Directive struct {
	Type string `json:"type"`
}

type ResponseContent struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	Card             Card         `json:"card"`
	Reprompt         OutputSpeech     `json:"reprompt"`
	Directives       []map[string]interface {}  `json:"directives"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

type ResponseBody struct {
	Version           string            `json:"version"`
	SessionAttributes map[string]string `json:"sessionAttributes"`
	Response          ResponseContent   `json:"response"`
}

type OutputSpeechResponseContent struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
	Card         Card         `json::"card"`
	Reprompt     OutputSpeech `json:"outputSpeech"`
}

type OutputSpeechResponse struct {
	Version           string                      `json:"version"`
	SessionAttributes map[string]string           `json:"sessionAttributes"`
	Response          OutputSpeechResponseContent `json:"response"`
}
