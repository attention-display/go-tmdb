package proto

import "encoding/json"

type CommonResponse struct {
	StatusCode    *int32 `json:"status_code,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
}

func UnmarshalCommonResponse(data []byte) (CommonResponse, error) {
	var r CommonResponse
	err := json.Unmarshal(data, &r)
	return r, err
}
