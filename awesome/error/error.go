package error

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	Code int32  `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func (e *Error) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("error: code=%v msg=%v", e.Code, e.Msg)
	}
	return string(b)
}

func New(code int32, msg string) error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}
