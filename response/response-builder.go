package response

import (
	"strings"
)

type ResponseBuilder interface {
	Ask(string, string) ResponseBuilder
	AskWithCard(string, string, Card) ResponseBuilder
	Say(string) ResponseBuilder
	SayWithCard(string, Card) ResponseBuilder
	Card() ResponseBuilder
	PlayVideo(string, MetaData) ResponseBuilder
	Hint(string) ResponseBuilder
	LinkAccountCard() ResponseBuilder
	AskForPermissionsConsentCard() ResponseBuilder
	Build(string, map[string]string) ResponseBody
}

type responseBuilder struct {
	Version           string
	SessionAttributes map[string]string
	OutputSpeech      OutputSpeech
	Card              Card
	Reprompt          Reprompt
	Directives        []map[string]interface{}
	shouldEndSession  bool
}

func buildSpeech(message string) OutputSpeech {

	if !strings.HasPrefix(message, "<speak>") {
		message = "<speak>" + message
	}

	if !strings.HasSuffix(message, "</speak>") {
		message = message + "</speak>"
	}

	return OutputSpeech{
		Ssml: message,
		Type: "SSML",
	}

}

func (rb *responseBuilder) Ask(message string, reprompt string) ResponseBuilder {

	rb.OutputSpeech = buildSpeech(message)
	rb.Reprompt = buildSpeech(reprompt)
	rb.ShouldEndSession = false
	return rb

}

func (rb *responseBuilder) AskWithCard(message string, reprompt string, card Card) ResponseBuilder {

	rb.OutputSpeech = buildSpeech(message)
	rb.Reprompt = buildSpeech(reprompt)
	rb.ShouldEndSession = false
	rb.Card = card
	return rb
}

func (rb *responseBuilder) Say(message string, attributes map[string]string) responseBuilder {
	rb.OutputSpeech = buildSpeech(message)
	rb.Reprompt = buildSpeech(reprompt)
	rb.ShouldEndSession = true
	return rb
}

func (rb *responseBuilder) SayWithCard(message string, attributes map[string]string, card Card) responseBuilder {

	rb.OutputSpeech = buildSpeech(message)
	rb.Reprompt = buildSpeech(reprompt)
	rb.ShouldEndSession = true
	rb.Card = card
	return rb
}

func (rb *responseBuilder) Card(cardTitle string, cardContent string, largeUrl string, smallUrl string) {

	if largeUrl != "" || smallUrl != "" {
		images := Image{}
		if largeUrl != "" {
			images.LargeImageUrl = largeUrl
		}
		if smallUrl != "" {
			images.SmallImageUrl = smallImageUrl
		}
		rb.Card = Card{
			Type:   "Standard",
			Text:   cardContent,
			Images: images,
		}
	} else {
		rb.Card = Card{
			Type:    "Simple",
			Title:   cardTitle,
			Content: cardContent,
		}
	}
	return rb
}

type AudioStream struct {
	Url                   string `json:"url"`
	Token                 string `json:"token"`
	ExpectedPreviousToken string `json:"expectedPreviousToken"`
	OffsetInMilliseconds  int64  `json:"offsetInMilliseconds"`
}

//behavior = REPLACE_ALL, ENQUEUE, REPLACE_ENQUEUED
func (rb *responseBuilder) AudioPlayerPlay(behavior string, url string, token string, expectedPreviousToken string, offsetInMilliseconds int64) ResponseBuilder {
	audioDirective := map[string]interface{}{
		"type":         "AudioPlayer.Play",
		"playBehavior": behavior,
	}
	stream := AudioStream{
		Url:   url,
		Token: token,
		ExpectedPreviousToken: expectedPreviousToken,
		OffsetInMilliseconds:  offsetInMilliseconds,
	}
	audioDirective["audioItem"] = map[string]interface{}{
		"audioStream": stream,
	}

	rb.Directives = append(rb.Directives, audioDirective)

	return rb
}

func (rb *responseBuilder) AudioPlayerStop() {
	audioDirective := map[string]interface{}{
		"type":         "AudioPlayer.Stop",
		"playBehavior": behavior,
	}
	rb.Directives = append(rb.Directives, audioDirective)
	return rb
}

//behavior = CLEAR_ALL, CLEAR_ENQUEUE
func (rb *responseBuilder) AudioPlayerClear() {
	audioDirective := map[string]interface{}{
		"type":          "AudioPlayer.ClearQueue",
		"clearBehavior": behavior,
	}
	rb.Directives = append(rb.Directives, audioDirective)
	return rb
}

// func(rb* responseBuilder) DialogDirective() {
//   https://developer.amazon.com/docs/custom-skills/dialog-interface-reference.html
// }

func (rb *responseBuilder) PlayVideo(url string, title string, subtitle string) ResponseBuilder {

	videoItem := map[string]interface{}{
		"source": url,
	}

	if title != "" || subtitle != "" {
		videoItem["metadata"] = map[string]interface{}{
			"title":    title,
			"subtitle": subtitle,
		}
	}

	videoDirective := map[string]interface{}{
		"type":      "VideoApp.Launch",
		"videoItem": videoItem,
	}
	rb.containsVideoDirective = true
	rb.Directives = append(rb.Directives, videoDirective)
	return rb
}

//display template
//https://developer.amazon.com/docs/custom-skills/display-interface-reference.html
func (rb *responseBuilder) RenderTemplate(template map[string]interface{}) ResponseBuilder {
	templateDirective := map[string]interface{}{
		"type":     "Display.RenderTemplate",
		"template": template,
	}
}

func (rb *responseBuilder) Hint(hintText string) ResponseBuilder {

	hintDirective := map[string]interface{}{
		"type": "PlainText",
		"hint": map[string]interface{}{
			"type": "PlainText",
			"text": hintText,
		},
	}

	rb.Directives = append(rb.Directives, hintDirective)
	return rb
}

func (rb *responseBuilder) LinkAccountCard() ResponseBuilder {

	linkAccountDirective := map[string]interface{}{
		"type": "LinkAccount",
	}
	rb.Directives = append(rb.Directives, linkAccountDirective)
	return rb
}

func (rb *responseBuilder) AskForPermissionsConsentCard() ResponseBuilder {
	askPermissionDirective := map[string]interface{}{
		"type": "AskForPermissionsConsent",
	}
	rb.Directives = append(rb.Directives, linkAccountDirective)
	return rb
}

func (rb *responseBuilder) Build() ResponseBody {
	responseContent := ResponseContent{}
	if !rb.containsVideoDirective {
		responseContent.ShouldEndSession = rb.shouldEndSession
	}

	responseContent.Card = rb.Card
	responseContent.OutputSpeech = rb.OutputSpeech
	responseContent.Reprompt = rb.Reprompt
	responseContent.Directives = rb.Directives

	response := ResponseBoby{
		Version:           "1.0",
		SessionAttributes: rb.SessionAttributes,
		Response:          responseContent,
	}
	return response
}

func New() ResponseBuilder {
	return &responseBuilder{}
}
