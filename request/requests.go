package request

type ResolutionStatus struct {
	Code string `json:"code"`
}

type ValueId struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type ResolutionValue struct {
	Value map[string]ValueId `json:"values"`
}

type ResolutionAuthority struct {
	Authority string            `json:"authority"`
	Status    ResolutionStatus  `json:"status"`
	Values    []ResolutionValue `json:"values"`
}

type Resolution struct {
	ResolutionsPerAuthority []ResolutionAuthority `json:"resolutionsPerAuthority"`
}

type Slot struct {
	Name               string     `json:"name"`
	Value              string     `json:"value"`
	ConfirmationStatus string     `json:"confirmationStatus"`
	Resolutions        Resolution `json:"resolutions"`
}

type Intent struct {
	Name      string          `json:"name"`
	Slots     map[string]Slot `json:"slots"`
	Type      string          `json:"type"`
	RequestId string          `json:"requestId"`
	Timestamp string          `json:"timestamp"`
	Locale    string          `json:locale`
}

type Launch struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	RequestId string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Locale    string `json:locale`
}

type Permission struct {
	ConsentToken string `json:"consentToken"`
}

type User struct {
	UserId      string     `json:"userId"`
	AccessToken string     `json:"accessToken"`
	Permissions Permission `json:"permissions"`
}

type Application struct {
	ApplicationId string `json:"applicationId"`
}

type Session struct {
	New         bool                         `json:"new"`
	SessionId   string                       `json:"sessionId"`
	Application Application                  `json:"application"`
	Attributes  map[string]map[string]string `json:"attributes"`
	User        User                         `json:"user"`
}

type SupportedDevice struct {
	DeviceId            string                       `json:"deviceId"`
	SupportedInterfaces map[string]map[string]string `json:"supportedInterfaces"`
}

type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity"`
	Token                string `json: "token"`
	OffsetInMilliseconds int64  `json:"offsetInMilliseconds"`
}

type System struct {
	Device         SupportedDevice `json:"device"`
	Application    Application     `json:"application`
	ApiEndpoint    string          `json:"apiEndpoint"`
	ApiAccessToken string          `json:"apiAccessToken"`
	User           User            `json:"user"`
}

type Context struct {
	System      System      `json: "system"`
	AudioPlayer AudioPlayer `json:audioPlayer`
}
