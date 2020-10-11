package sxerror

import (
	"regexp"
	"fmt"
	"runtime"
)

// An SxError represents an error that might occur in savannaah gRPC servers
type SxError struct {
	code string // code to describe the error
	fileName string // filename where the error has occurred
	lineNumber int // line number where the error has occurred
	title string // title of the error
	description string // full description of the error
}

func New(code, description string) error {
	matched, err := regexp.MatchString("SX[DIWSB][A-Z]{3}[0-9]{6}$",code)
	if err != nil {
		panic("something went wrong while creating regex")
	}
	if !matched {
		panic("unrecognised error code. expected pattern: SX[DIWSB][A-Z]{3}[0-9]{6}$")
	}
	title, ok := errorTitle[code[len(code)-6:]]
	if !ok {
		panic(fmt.Sprintf("unknown error digit. please make an entry for %s",code[len(code)-6:]) )
	}
	_, fileName, lineNumber, _ := runtime.Caller(1)

	return &SxError{code,fileName, lineNumber, title, description}
}

func (e *SxError) Error() string {
	if e != nil {
		return fmt.Sprintf(`{"code":"%s", "fileName":"%s", "lineNumber":%d, "title":"%s", "description":"%s"}`, e.code, e.fileName, e.lineNumber, e.title, e.description)
	}
	return ""
}

func (e *SxError) Code() string {
	if e != nil {
		return e.code
	}
	return ""
}

func (e *SxError) FileName() string {
	if e != nil {
		return e.fileName
	}
	return ""
}

func (e *SxError) LineNumber() int {
	if e != nil {
		return e.lineNumber
	}
	return 0
}

func (e *SxError) Title() string {
	if e != nil {
		return e.title
	}
	return ""
}

func (e *SxError) Description() string {
	if e != nil {
		return e.description
	}
	return ""
}
