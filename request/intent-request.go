package requests

type IntentRequest struct {
    Version  string `json:"version"`
    Session Session `json:"session"`
    Context Context `json:context`
    Request struct {
        Timestamp time.Time `json:"timestamp, string"`
        ReqType string `json:"type"`
        RequestId string `json:"requestId"`
        DialogState string `json:"dialogState"`
        Locale string `json:"string"`
        Intent struct {
            Type string `json:"type"`
            Name string `json:"name"`
            ConfirmationStatus string `json:"confirmationStatus`
            Slots map[string]Slot `json: "slots"`
        } `json:"intent"`
    } `json:request`
}
