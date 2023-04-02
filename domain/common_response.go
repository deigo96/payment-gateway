package domain

type EmptyObj struct{}

// Response is used for static shape json return
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"status"`
	Data    interface{} `json:"data"`
}

// BuildResponse method is to inject data value to dynamic success response
func BuildResponse(code int, message string, data interface{}) Response {
	res := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return res
}
