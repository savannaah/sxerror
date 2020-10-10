package sxerror

import "regexp"

// An SxError represents an error that might occur in savannaah gRPC servers
type SxError struct {
	code string // code to describe the error
	fileName string // filename where the error has occurred
	lineNumber string // line number where the error has occurred
	title string // title of the error
	description string // full description of the error
}

func New(code string, fileName, lineNumber, title, description string) error {
	matched, err := regexp.MatchString("SX[IWSB][A-Z]{3}[0-9]{6}$",code)
	if err != nil {
		panic("something went wrong while creating regex")
	}
	if !matched {
		panic("unrecognised error code. expected pattern: SX[IWSB][0-9]{6}$")
	}
	return &SxError{code,fileName, lineNumber, title, description}
}

func (e *SxError) Error() string {
	if e != nil {
		return e.code + " - " + e.title
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

func (e *SxError) LineNumber() string {
	if e != nil {
		return e.lineNumber
	}
	return ""
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
