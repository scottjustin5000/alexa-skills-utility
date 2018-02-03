package requests

type AudioPlayerRequest struct {
    Version  string `json:"version"`
    Context Context `json:context`
    Request struct {
        Timestamp time.Time `json:"timestamp, string"`
        ReqType string `json:"type"`
        RequestId string `json:"requestId"`
        Locale string `json:"locale"`
        Token string `json:"token"`
        OffsetInMilliseconds int64 `json:"offsetInMilliseconds"`
        
    } `json:request`
}
