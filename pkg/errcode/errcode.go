package errcode

import (
	"fmt"
	"net/http"
)

// Error ...
type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

// NewError ...
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

// Error ...
func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.Code(), e.Msg())
}

// Code ...
func (e *Error) Code() int {
	return e.code
}

// Msg ...
func (e *Error) Msg() string {
	return e.msg
}

// Msgf ...
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

// Details ...
func (e *Error) Details() []string {
	return e.details
}

// WithDetails ...
func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}

	return e
}

// StatusCode ...
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		fallthrough
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
