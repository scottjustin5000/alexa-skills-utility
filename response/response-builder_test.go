package response
import (
  "testing"
  "fmt"
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
  fmt.Println(getType(rb))
  assert.Equal(t, getType(rb), "responseBuilder","should be an instance of response builder")

}

func TestSimpleSayCreate(t *testing.T) {
  rb :=New()
  attributes := map[string]string{}
  response := rb.Say("Hi Alexa").Build(attributes)
  assert.Equal(t, response.Response.OutputSpeech.Ssml, "<speak>Hi Alexa</speak>", "Ssml should be set and match")

}