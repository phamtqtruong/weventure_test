package rest

// Response ...
type Response struct {
	Status int         `json:"status"`
	Code   int         `json:"code,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Error  *Error      `json:"error,omitempty"`
}

// Error ...
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}
