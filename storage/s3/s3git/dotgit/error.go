package dotgit

import (
	"fmt"
	"strconv"
	"strings"
)

type Error struct {
	msg     string
	code    ErrorCode
	err     error
	details []string
}

func (e Error) Error() string {
	if e.msg == "" {
		if msg, ok := messages[e.code]; ok {
			e.msg = msg
		}
	}
	res := "s3 (" + e.code.String() + ")"
	if e.msg != "" {
		res += " - " + e.msg
	}
	if len(e.details) > 0 {
		res += fmt.Sprintf("[%q]", strings.Join(e.details, ","))
	}
	return res
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e Error) Is(err error) bool {
	if x, ok := err.(*Error); ok {
		return e.code == x.code
	}
	if x, ok := err.(Error); ok {
		return e.code == x.code
	}
	if x, ok := err.(ErrorCode); ok {
		return e.code == x
	}
	return false
}

func (e Error) WithError(err error) *Error {
	e.err = err
	return &e
}

func (e Error) WithMsg(msg string) *Error {
	e.msg = msg
	return &e
}

func (e Error) AddDetail(msg string, args ...interface{}) *Error {
	e.details = append(e.details, fmt.Sprintf(msg, args...))
	return &e
}

type ErrorCode int64

func (e ErrorCode) Error() string {
	return Error{
		code: e,
	}.Error()
}

func (e ErrorCode) WithError(err error) *Error {
	return &Error{
		code: e,
		err:  err,
	}
}

func (e ErrorCode) ToError() *Error {
	return &Error{
		code: e,
	}
}

func (e ErrorCode) String() string {
	if _, ok := messages[e]; ok {
		return strconv.FormatInt(int64(e), 10)
	} else {
		return "unknown"
	}
}

type ErrMessages map[ErrorCode]string

var messages = ErrMessages{
	ErrPackedRefsBadFormat: "malformed packed-ref",
}

const (
	ErrPackedRefsBadFormat = ErrorCode(2101)
)
