package sxerror

import "regexp"

type sxError struct {
	code string
	message string
}

/*
	Error code pattern: SX[IWSB][0-9]{6}$
    =============================================
	First two character is reserved for SAVANNAAH
	Third character is for error level
	I = info
    W = warning
    S = severe - internal errors
    B = business - error made by users
	First three digits is used to identify the micro service
    Last three digits is used to identify specific kind of error
 */

func New(code string, message string) error {
	matched, err := regexp.MatchString("SX[IWSB][0-9]{6}$",code)
	if err != nil {
		panic("something went wrong while creating regex")
	}
	if !matched {
		panic("unrecognised error code. expected pattern: SX[IWSB][0-9]{6}$")
	}
	return &sxError{code: code, message: message}
}

func (e *sxError) Error() string {
	if e != nil {
		return e.code + ": " + e.message
	}
	return ""
}

func (e *sxError) Code() string {
	if e != nil {
		return e.code
	}
	return ""
}

func (e *sxError) Message() string {
	if e != nil {
		return e.message
	}
	return ""
}