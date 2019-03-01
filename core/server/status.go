package server

import (
	"fmt"
	"strconv"
)

const (
	STATUS_SERVER_SUCCESS_CODE = iota
	STATUS_SERVER_ERROR_CODE
	STATUS_PARAM_PARES_ERROR_CODE
)

var ErrorMap = map[int]string{
	STATUS_SERVER_ERROR_CODE:      "服务器异常",
	STATUS_PARAM_PARES_ERROR_CODE: "参数解析异常",
}

func GetErrorReason(code int) string {
	return ErrorMap[code]
}

func RegisterErrorReason(code int, reason string) {
	if _, ok := ErrorMap[code]; ok {
		fmt.Println("code" + strconv.Itoa(code) + "is already")
	}
	ErrorMap[code] = reason
}
func init() {
	RegisterErrorReason(STATUS_SERVER_SUCCESS_CODE, "")
}
