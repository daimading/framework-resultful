package server

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type ErrorDetail struct {
	file     string
	funcName string
	line     int

	StatusCode int
	StatusMsg  string
}

func (e *ErrorDetail) Error() string {
	return fmt.Sprintf("[%s:%d] ID: %d Value: %s", e.file, e.line, e.StatusCode, e.StatusMsg)
}

func NewError(errorCode int) *ErrorDetail {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc)
	frame, _ := frames.Next()

	funcName := frame.Func.Name()
	funcName = funcName[strings.LastIndexByte(funcName, filepath.Separator)+1:]
	fileName := frame.File[strings.LastIndexByte(frame.File, filepath.Separator)+1:]
	errorDetail := ErrorDetail{StatusCode: errorCode, StatusMsg: GetErrorReason(errorCode), file: fileName, funcName: funcName, line: frame.Line}
	return &errorDetail
}
