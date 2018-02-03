package requests

//exposed 
https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#response-body-syntax
//https://developer.amazon.com/docs/custom-skills/request-types-reference.html#launchrequest
https://developer.amazon.com/docs/custom-skills/request-types-reference.html#intentrequest
IntentRequest

https://developer.amazon.com/docs/custom-skills/request-types-reference.html#launchrequest
LaunchRequest

https://developer.amazon.com/docs/custom-skills/request-types-reference.html#sessionendedrequest
SessionEndedRequest
https://developer.amazon.com/docs/custom-skills/playback-controller-interface-reference.html#requests

https://developer.amazon.com/docs/custom-skills/audioplayer-interface-reference.html#requests





/*

{
  "type": "IntentRequest",
  "requestId": "string",
  "timestamp": "string",
  "dialogState": "string",
  "locale": "string",
  "intent": {
    "name": "string",
    "confirmationStatus": "string",
    "slots": {
      "SlotName": {
        "name": "string",
        "value": "string",
        "confirmationStatus": "string",
        "resolutions": {
          "resolutionsPerAuthority": [
            {
              "authority": "string",
              "status": {
                "code": "string"
              },
              "values": [
                {
                  "value": {
                    "name": "string",
                    "id": "string"
                  }
                }
              ]
            }
          ]
        }
      }
    }
  }
}

*/
type ResolutionStatus  struct {
  Code string `json:"code"`
}

type ValueId struct {
  Name string `json:"name"`
  Id string `json:"id"`
}

type ResolutionValue struct {
  Value map[string]ValueId `json:"values"`
}


type ResolutionsAuthority struct {
  Authority string `json:"authority"`
  Status ResolutionStatus `json:"status"`
  Values []ResolutionValue `json:"values"`
}

type Resolution struct {
  ResolutionsPerAuthority []ResolutionAuthority `json:"resolutionsPerAuthority"`
}

type Slot struct {
  Name string `json:"name"`
  Value string `json:"value"`
  ConfirmationStatus string `json:"confirmationStatus"`
  Resolutions Resolution `json:"resolutions"`
}

type Intent struct {
  Name string `json:"name"`
  Slots map[string]Slot `json:"slots"`
  Type string `json:"type"`
  RequestId string `json:"requestId"`
  Timestamp string `json:"timestamp"`
  Locale string `json:locale`
 }


type Launch struct {
  Name string `json:"name"`
  Type string `json:"type"`
  RequestId string `json:"requestId"`
  Timestamp string `json:"timestamp"`
  Locale string `json:locale`
}

type Permission struct {
  ConsentToken string `json:"consentToken"`
}

type User struct {
  UserId string `json:"userId"`,
  AccessToken string `json:"accessToken"`,
  Permissions Permission `json:"permissions"`,
}

type Application struct {
  ApplicationId string `json:"applicationId"`
}

type Attribute struct {
  Key string `json:"key"`
} 

type Session struct {
  New bool `json:"new"`
  SessionId string `json:"sessionId"`
  Application Application `json:"application"`
  Attributes Attribute `json:"attributes"`
  User User `json:"user"`
}

type SupportedInterface struct {
  Display  interface{}
  AudioPlayer interface{} 
  VideoApp interface{}
}

type SupportedDevice struct {
  DeviceId string `json:"deviceId"`
  SupportedInterfaces SupportedInterface `json:"supportedInterfaces"`
}

type AudioPlayer struct {
  PlayerActivity string `json:"playerActivity"`
  Token string `json: "token"`
  OffsetInMilliseconds int64 `json:"offsetInMilliseconds"`
}

type System struct {
  Device SupportedDevice `json:"deviceId"`
  Application Application `json:"application`
  ApiEndpoint string `json:"apiEndpoint"`
  ApiAccessToken string `json:"apiAccessToken"`
  User User `json:"user"`
}

type Context struct {
  System System `json: "system"`
  AudioPlayer AudioPlayer `json:audioPlayer`
}
