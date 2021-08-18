package controllers

type APIResponse struct {
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
}

func NewAPIError(errno int, msg string) APIResponse {
	return APIResponse{
		Data: nil,
		Meta: map[string]interface{}{
			"errno": errno,
			"error": msg,
		},
	}
}

func NewAPIResponse(data interface{}) APIResponse {
	return APIResponse{
		Data: data,
		Meta: map[string]interface{}{
			"errno": 0,
		},
	}
}
