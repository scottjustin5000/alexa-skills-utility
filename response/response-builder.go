package response

import (
	"strings"
)

type ResponseBuilder interface {
	Ask(string, string) ResponseBuilder
	AskWithCard(string, string, Card) ResponseBuilder
	Say(string) ResponseBuilder
	Whisper(string) ResponseBuilder
	SayWithCard(string, Card) ResponseBuilder
	CreateCard(string, string, string, string) ResponseBuilder
	PlayVideo(string, string, string) ResponseBuilder
	DialogDirective(bool) ResponseBuilder
	Hint(string) ResponseBuilder
	LinkAccountCard() ResponseBuilder
	AskForPermissionsConsentCard() ResponseBuilder
	Build(map[string]string) ResponseBody
	AudioPlayerPlay(string, string, string, string, int64) ResponseBuilder
	AudioPlayerStop() ResponseBuilder
	AudioPlayerClear(string) ResponseBuilder
}

type responseBuilder struct {
	Version                string
	SessionAttributes      map[string]string
	OutputSpeech           OutputSpeech
	Card                   Card
	Reprompt               OutputSpeech
	Directives             []map[string]interface{}
	shouldEndSession       bool
	containsVideoDirective bool
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
	rb.shouldEndSession = false
	return rb

}

func (rb *responseBuilder) AskWithCard(message string, reprompt string, card Card) ResponseBuilder {

	rb.OutputSpeech = buildSpeech(message)
	rb.Reprompt = buildSpeech(reprompt)
	rb.shouldEndSession = false
	rb.Card = card
	return rb
}

func (rb *responseBuilder) Say(message string) ResponseBuilder {
	rb.OutputSpeech = buildSpeech(message)
	rb.shouldEndSession = true
	return rb
}

func (rb *responseBuilder) Whisper(message string) ResponseBuilder {
	text := "<amazon:effect type=\"whispered\">" + message + "</amazon:effect>"
	rb.OutputSpeech = buildSpeech(text)
	rb.shouldEndSession = true
	return rb
}

func (rb *responseBuilder) SayWithCard(message string, card Card) ResponseBuilder {

	rb.OutputSpeech = buildSpeech(message)
	rb.shouldEndSession = true
	rb.Card = card
	return rb
}

func (rb *responseBuilder) CreateCard(cardTitle string, cardContent string, largeUrl string, smallUrl string) ResponseBuilder {

	if largeUrl != "" || smallUrl != "" {
		images := Image{}
		if largeUrl != "" {
			images.LargeImageUrl = largeUrl
		}
		if smallUrl != "" {
			images.SmallImageUrl = smallUrl
		}
		rb.Card = Card{
			Type:    "Standard",
			Title:   cardTitle,
			Content: cardContent,
			Images:  images,
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

func (rb *responseBuilder) AudioPlayerStop() ResponseBuilder {
	audioDirective := map[string]interface{}{
		"type": "AudioPlayer.Stop",
	}
	rb.Directives = append(rb.Directives, audioDirective)
	return rb
}

//behavior = CLEAR_ALL, CLEAR_ENQUEUE
func (rb *responseBuilder) AudioPlayerClear(behavior string) ResponseBuilder {
	audioDirective := map[string]interface{}{
		"type":          "AudioPlayer.ClearQueue",
		"clearBehavior": behavior,
	}
	rb.Directives = append(rb.Directives, audioDirective)
	return rb
}

func(rb *responseBuilder) DialogDirective(endSession bool) ResponseBuilder {
  dialogDirective := map[string]interface{}{
		"type":          "Dialog.Delegate"
	}
	rb.Directives = append(rb.Directives, dialogDirective)
	rb.shouldEndSession = endSession
	return rb
}

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
	rb.Directives = append(rb.Directives, templateDirective)
	return rb
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
	rb.Directives = append(rb.Directives, askPermissionDirective)
	return rb
}

func (rb *responseBuilder) Build(attributes map[string]string) ResponseBody {
	responseContent := ResponseContent{}
	if !rb.containsVideoDirective {
		responseContent.ShouldEndSession = rb.shouldEndSession
	}
	responseContent.Card = rb.Card
	responseContent.OutputSpeech = rb.OutputSpeech
	responseContent.Reprompt = rb.Reprompt
	responseContent.Directives = rb.Directives

	response := ResponseBody{
		Version:           "1.0",
		SessionAttributes: rb.SessionAttributes,
		Response:          responseContent,
	}
	return response
}

func New() ResponseBuilder {
	return &responseBuilder{}
}
