package response

type ResponseProject struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    interface{}
}

func NewResponseProject(code int, message, status string, data interface{}) *ResponseProject {
	return &ResponseProject{
		Code:    code,
		Message: message,
		Status:  status,
		Data:    data,
	}
}
