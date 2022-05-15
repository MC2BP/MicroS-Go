package httplib

type Response struct {
	Success   bool        `json:"Success"`
	Data      interface{} `json:"Data"`
	Message   string      `json:"Message,omitempty"`
	ErrorCode int         `json:"ErrorCode,omitempty"`
}
