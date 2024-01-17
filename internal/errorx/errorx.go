package errorx

import (
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Error struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Detail    string                `json:"detail,omitempty"`
	Info      map[string]any        `json:"info,omitempty"`
	Timestamp timestamppb.Timestamp `json:"timestamp"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("{code: %d, message: %s}", err.Code, err.Message)
}

func NewAPIError(code int, message string) error {
	return &Error{
		Code:      code,
		Message:   message,
		Timestamp: *timestamppb.Now(),
	}
}

func NewAPIErrorWithDetail(code int, message, detail string) error {
	return &Error{
		Code:      code,
		Message:   message,
		Detail:    detail,
		Timestamp: *timestamppb.Now(),
	}
}

func NewAPIErrorWithInfo(code int, message string, info map[string]any) error {
	return &Error{
		Code:      code,
		Message:   message,
		Info:      info,
		Timestamp: *timestamppb.Now(),
	}
}
