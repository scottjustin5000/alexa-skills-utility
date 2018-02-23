package response
import (
  "testing"
  "reflect"
  "github.com/stretchr/testify/assert"
)


func getType(myvar interface{}) string {
    if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
        return t.Elem().Name()
    } else {
        return t.Name()
    }
}

func TestResponseBuilderCreate(t *testing.T) {
  rb := New()
  
  _, ok := rb.(ResponseBuilder)
  if !ok {
    t.Error("Assertion error")
  }
  assert.Equal(t, getType(rb), "responseBuilder","should be an instance of response builder")

}

func TestSimpleSayCreate(t *testing.T) {
  rb := New()
  attributes := map[string]string{}
  response := rb.Say("Hi Alexa").Build(attributes)
  assert.Equal(t, response.Response.OutputSpeech.Ssml, "<speak>Hi Alexa</speak>", "Ssml should be set and match")

}

func TestSimpleAsk(t *testing.T) {
  rb := New()
  attributes := map[string]string{}
  response := rb.Ask("Hi Alexa", "Are you there?").Build(attributes)
  assert.Equal(t, response.Response.OutputSpeech.Ssml, "<speak>Hi Alexa</speak>", "Ssml should be set and match")
  assert.Equal(t, response.Response.Reprompt.Ssml, "<speak>Are you there?</speak>", "reprompt ssml should match")
}

func TestWhisper(t *testing.T) {
  rb := New()
  attributes := map[string]string{}
  response := rb.Whisper("Can you keep a secret?").Build(attributes)
  assert.Equal(t, response.Response.OutputSpeech.Ssml, "<speak><amazon:effect type=\"whispered\">Can you keep a secret?</amazon:effect></speak>", "Ssml should include whisper effect")
}

func TestCreateCard(t *testing.T) {
   rb := New()
  attributes := map[string]string{}
  response := rb.CreateCard("title", "some great content", "https://somewhere/large_image.png", "https://somewhere/small_image.png").Build(attributes)
  assert.Equal(t, response.Response.Card.Type, "Standard", "Card type should be standard")
  assert.Equal(t, response.Response.Card.Title, "title","Card title should match")
  assert.Equal(t, response.Response.Card.Content, "some great content","Card Content should match")
  assert.Equal(t, response.Response.Card.Images.SmallImageUrl, "https://somewhere/small_image.png", "small image should match")
  assert.Equal(t, response.Response.Card.Images.LargeImageUrl, "https://somewhere/large_image.png", "large image should match")
}

func TestSayWithCard(t *testing.T) {
  rb := New()
  c := Card{
    Type: "Standard",
    Title: "s'up card",
    Content: "Card Content",
    Images: Image{
      SmallImageUrl: "https://somewhere/small_image.png",
      LargeImageUrl: "https://somewhere/large_image.png",
    },
  }
  attributes := map[string]string{}
  response := rb.SayWithCard("Check out this card", c).Build(attributes)
  assert.Equal(t, response.Response.OutputSpeech.Ssml, "<speak>Check out this card</speak>", "Ssml should match")
  assert.Equal(t, response.Response.Card.Type, "Standard", "Card type should be standard")
  assert.Equal(t, response.Response.Card.Title, "s'up card","Card title should match")
  assert.Equal(t, response.Response.Card.Content, "Card Content","Card Content should match")
  assert.Equal(t, response.Response.Card.Images.SmallImageUrl, "https://somewhere/small_image.png", "small image should match")
  assert.Equal(t, response.Response.Card.Images.LargeImageUrl, "https://somewhere/large_image.png", "large image should match")
}

func TestChainSayAndCard(t *testing.T) {
  rb := New()
  attributes := map[string]string{}
  response := rb.Say("Hello I'm going to include a card").CreateCard("title", "some great content", "https://somewhere/large_image.png", "https://somewhere/small_image.png").Build(attributes)

  assert.Equal(t, response.Response.OutputSpeech.Ssml, "<speak>Hello I'm going to include a card</speak>", "Ssml should match")
  assert.Equal(t, response.Response.Card.Type, "Standard", "Card type should be standard")
  assert.Equal(t, response.Response.Card.Title, "title","Card title should match")
  assert.Equal(t, response.Response.Card.Content, "some great content","Card Content should match")
  assert.Equal(t, response.Response.Card.Images.SmallImageUrl, "https://somewhere/small_image.png", "small image should match")
  assert.Equal(t, response.Response.Card.Images.LargeImageUrl, "https://somewhere/large_image.png", "large image should match")
}

func TestAskWithCard(t *testing.T) {

  rb := New()
  attributes := map[string]string{}
  response := rb.Ask("hi","Do you know who this is?").CreateCard("person", "of interest", "https://somewhere/large_image.png", "https://somewhere/small_image.png").Build(attributes)

  assert.Equal(t, response.Response.OutputSpeech.Ssml, "<speak>hi</speak>", "Ssml should match")
  assert.Equal(t, response.Response.Reprompt.Ssml, "<speak>Do you know who this is?</speak>", "reprompt ssml should match")
  assert.Equal(t, response.Response.Card.Type, "Standard", "Card type should be standard")
  assert.Equal(t, response.Response.Card.Title, "person","Card title should match")
  assert.Equal(t, response.Response.Card.Content, "of interest","Card Content should match")
  assert.Equal(t, response.Response.Card.Images.SmallImageUrl, "https://somewhere/small_image.png", "small image should match")
  assert.Equal(t, response.Response.Card.Images.LargeImageUrl, "https://somewhere/large_image.png", "large image should match")
}

func TestAudioPlayerPlay(t *testing.T) {
  rb := New()
  attributes := map[string]string{}
  response := rb.AudioPlayerPlay("REPLACE_ALL", "https://somewhere/large_image.mp3", "token", "prevToken", 3000).Build(attributes)
  audioItem := response.Response.Directives[0]["audioItem"].(map[string]interface {})
  audioStream := audioItem["audioStream"].(AudioStream)
  
  assert.Equal(t, response.Response.Directives[0]["type"],"AudioPlayer.Play")
  assert.Equal(t, response.Response.Directives[0]["playBehavior"], "REPLACE_ALL")
  assert.Equal(t, audioStream.Url, "https://somewhere/large_image.mp3")
}

func TestAudioPlayerStop(t *testing.T) {
  rb := New()
  attributes := map[string]string{}
  response := rb.AudioPlayerStop().Build(attributes)
  assert.Equal(t, response.Response.Directives[0]["type"],"AudioPlayer.Stop")

}

func TestAudioPlayerClear(t *testing.T) {
  rb := New()
  attributes := map[string]string{}
  response := rb.AudioPlayerClear("CLEAR_ALL").Build(attributes)
  assert.Equal(t, response.Response.Directives[0]["type"],"AudioPlayer.ClearQueue")
  assert.Equal(t, response.Response.Directives[0]["clearBehavior"],"CLEAR_ALL")
}

func TestPlayVideo(t *testing.T) {
  rb := New()
  attributes := map[string]string{}
  response := rb.PlayVideo("http://fake/move.mov", "movie title", "movie subtitle").Build(attributes)
  videoItem := response.Response.Directives[0]["videoItem"].(map[string]interface {})
  meta := videoItem["metadata"].(map[string]interface{})
  
  assert.Equal(t, response.Response.Directives[0]["type"],"VideoApp.Launch")
  assert.Equal(t, videoItem["source"], "http://fake/move.mov")
  assert.Equal(t, meta["title"], "movie title")
  assert.Equal(t, meta["subtitle"], "movie subtitle")

}

func TestRenderTemplate(t *testing.T) {

}

func TestLinkAccountCard(t *testing.T) {
  rb := New()
  attributes := map[string]string{}
  response := rb.LinkAccountCard().Build(attributes)
  assert.Equal(t, response.Response.Directives[0]["type"],"LinkAccount")
}

func TestAskForPermissionsConsentCard(t *testing.T) {
  rb := New()
  attributes := map[string]string{}
  response := rb.AskForPermissionsConsentCard().Build(attributes)
  assert.Equal(t, response.Response.Directives[0]["type"],"AskForPermissionsConsent")
}
