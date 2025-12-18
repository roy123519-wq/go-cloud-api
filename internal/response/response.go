package response

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Data  any    `json:"data"`
	Error *Error `json:"error"`
}

func Success(data any) Response {
	return Response{
		Data:  data,
		Error: nil,
	}
}
func Fail(code, message string) Response {
	return Response{
		Data: nil,
		Error: &Error{
			Code:    code,
			Message: message,
		},
	}
}
