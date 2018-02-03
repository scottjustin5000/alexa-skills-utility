package requests
type LaunchRequest struct {
    Version  string `json:"version"`
    Session Session `json:"session"`
    Context Context `json:context`
    Request struct {
        Timestamp time.Time `json:"timestamp, string"`
        ReqType string `json:"type"`
        RequestId string `json:"requestId"`
        DialogState string `json:"dialogState"`
        Locale string `json:"string"`
    } `json:request`
}