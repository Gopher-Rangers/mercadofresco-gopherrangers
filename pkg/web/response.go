package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(statusCode int, data interface{}) (int, Response) {
	return statusCode, Response{statusCode, data, ""}
}

func DecodeError(statusCode int, err string) (int, Response) {
	return statusCode, Response{statusCode, nil, err}
}
