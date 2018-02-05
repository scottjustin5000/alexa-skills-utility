package request

import (
  "testing"
  "encoding/json"

  "github.com/stretchr/testify/assert"
) 


func TestLaunchRequest(t *testing.T) {
  requestString := `{"version": "1.0", "session": {"new": true, "sessionId": "amzn1.echo-api.session.session_id", "application": {"applicationId": "amzn1.ask.skill.application_id"}, "attributes": {"key": {"value": "123"} }, "user": {"userId": "amzn1.ask.account.1234", "accessToken": "Atza|AAAAAAAA", "permissions": {"consentToken": "ZZZZZZZ"} } }, "context": {"System": {"device": {"deviceId": "string", "supportedInterfaces": {"AudioPlayer": {} } }, "application": {"applicationId": "amzn1.ask.skill.skill_id"}, "user": {"userId": "amzn1.ask.account.user_id", "accessToken": "Atza|AAAAAAAA", "permissions": {"consentToken": "ZZZZZZZ"} }, "apiEndpoint": "https://api.amazonalexa.com", "apiAccessToken": "AxThk"}, "AudioPlayer": {"playerActivity": "PLAYING", "token": "audioplayer-token", "offsetInMilliseconds": 1 } }, "request": {"type": "LaunchRequest", "requestId": "req123", "timestamp": "2018-02-05T01:01:20Z", "locale": "en-US"} }`
  res := &LaunchRequest{}

  err := json.Unmarshal([]byte(requestString), res)

  if err != nil {
    t.Errorf("something went wrong", err)
  }

  assert.Equal(t, res.Version, "1.0", "version should be equal")
  assert.Equal(t, res.Session.New, true, "version should be equal")
  assert.Equal(t, res.Session.SessionId, "amzn1.echo-api.session.session_id", "sessionId should match")
  assert.Equal(t, res.Session.Application.ApplicationId, "amzn1.ask.skill.application_id", "app id should match")
  n := len(res.Session.Attributes)
  v := res.Session.Attributes["key"]
  assert.Equal(t, n, 1, "attributes length should be 1")
  assert.Equal(t, v["value"], "123", "attributes key value should = 123")
  assert.Equal(t, res.Session.User.UserId, "amzn1.ask.account.1234", "userIds should match")
  assert.Equal(t, res.Session.User.AccessToken, "Atza|AAAAAAAA", "access tokens should match")
  assert.Equal(t, res.Session.User.Permissions.ConsentToken, "ZZZZZZZ", "consent tokens should match")

  assert.Equal(t, res.Context.System.Application.ApplicationId, "amzn1.ask.skill.skill_id", "context.applicationIds should match")
  assert.Equal(t, res.Context.System.User.UserId, "amzn1.ask.account.user_id", "context.userIds should match")
  interfaces := len(res.Context.System.Device.SupportedInterfaces)
  assert.Equal(t, int(interfaces), int(1), "device supported interfaces length should be 1")
  assert.Equal(t, int(res.Context.AudioPlayer.OffsetInMilliseconds), 1, "offsetInMilliseconds should match")
  assert.Equal(t, res.Context.AudioPlayer.PlayerActivity, "PLAYING", "playerActivity should match")
  assert.Equal(t, "LaunchRequest", res.Request.ReqType, "Request type should be LaunchRequest")
  assert.Equal(t, "req123", res.Request.RequestId, "Request Id should be req123")
  assert.Equal(t, "en-US", res.Request.Locale, "Locale should be en-US")

}

func TestIntentRequest(t *testing.T) {
  requestString := `{"version": "1.0", "session": {"new": true, "sessionId": "amzn1.echo-api.session.session_id", "application": {"applicationId": "amzn1.ask.skill.application_id"}, "attributes": {"key": {"value": "123"} }, "user": {"userId": "amzn1.ask.account.1234", "accessToken":"Atza|AAAAAAAA", "permissions":{"consentToken": "ZZZZZZZ"}  } }, "context": {"System": {"application": {"applicationId": "amzn1.ask.skill.skill_id"}, "user": {"userId": "amzn1.ask.account.user_id"}, "device": {"supportedInterfaces": {"AudioPlayer": {} } } }, "AudioPlayer": {"offsetInMilliseconds": 1, "playerActivity": "IDLE"} }, "request": {"type": "IntentRequest", "requestId": "amzn1.echo-api.request.123", "timestamp": "2015-05-13T12:34:56Z", "dialogState": "COMPLETED", "locale": "en-US", "intent": {"name": "GetZodiacHoroscopeIntent", "confirmationStatus": "NONE","slots": {"ZodiacSign": {"name": "ZodiacSign", "value": "virgo", "confirmationStatus": "NONE"} } } } }`
  res := &IntentRequest{}

  err := json.Unmarshal([]byte(requestString), res)

  if err != nil {
    t.Errorf("something went wrong", err)
  }

  assert.Equal(t, res.Version, "1.0", "version should be equal")
  assert.Equal(t, res.Session.New, true, "version should be equal")
  assert.Equal(t, res.Session.SessionId, "amzn1.echo-api.session.session_id", "sessionId should match")
  assert.Equal(t, res.Session.Application.ApplicationId, "amzn1.ask.skill.application_id", "app id should match")
  n := len(res.Session.Attributes)
  v := res.Session.Attributes["key"]
  assert.Equal(t, n, 1, "attributes length should be 1")
  assert.Equal(t, v["value"], "123", "attributes key value should = 123")
  assert.Equal(t, res.Session.User.UserId, "amzn1.ask.account.1234", "userIds should match")
  assert.Equal(t, res.Session.User.AccessToken, "Atza|AAAAAAAA", "access tokens should match")
  assert.Equal(t, res.Session.User.Permissions.ConsentToken, "ZZZZZZZ", "consent tokens should match")

  assert.Equal(t, res.Context.System.Application.ApplicationId, "amzn1.ask.skill.skill_id", "context.applicationIds should match")
  assert.Equal(t, res.Context.System.User.UserId, "amzn1.ask.account.user_id", "context.userIds should match")
  interfaces := len(res.Context.System.Device.SupportedInterfaces)
  assert.Equal(t, int(interfaces), int(1), "device supported interfaces length should be 1")

  assert.Equal(t, int(res.Context.AudioPlayer.OffsetInMilliseconds), 1, "offsetInMilliseconds should match")
  assert.Equal(t, res.Context.AudioPlayer.PlayerActivity, "IDLE", "playerActivity should match")
  assert.Equal(t, "IntentRequest", res.Request.ReqType, "Request type should be IntentRequest")
  assert.Equal(t, "amzn1.echo-api.request.123", res.Request.RequestId, "Request Id should be amzn1.echo-api.request.123")
  assert.Equal(t, "COMPLETED", res.Request.DialogState, "Dialog state should be complete")
  assert.Equal(t, "en-US", res.Request.Locale, "Locale should be en-US")
  assert.Equal(t, "GetZodiacHoroscopeIntent", res.Request.Intent.Name, "Name should be GetZodiacHoroscopeIntent")
  assert.Equal(t, "NONE", res.Request.Intent.ConfirmationStatus, "ConfirmationStatus should be NONE")
  assert.Equal(t, "virgo", res.Request.Intent.Slots["ZodiacSign"].Value, "Slot value should be ZodiacSign")

}

func TestPlayBackControllerRequest(t *testing.T) {

}

func TestAudioPlayerRequest(t *testing.T) {

}
