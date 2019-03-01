package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseResponse struct {
	Status int         `json:"status"`
	Reason string      `json:"reason"`
	Data   interface{} `json:"data,omitempty"`
}

func HandleError(err error) *BaseResponse {
	baseResult := BaseResponse{Status: STATUS_SERVER_SUCCESS_CODE, Reason: ""}
	if err != nil {
		switch err.(type) {
		case *ErrorDetail:
			err2 := err.(*ErrorDetail)
			baseResult = BaseResponse{Status: err2.StatusCode, Reason: err2.StatusMsg}
		default:
			baseResult = BaseResponse{Status: STATUS_SERVER_ERROR_CODE, Reason: ErrorMap[STATUS_SERVER_ERROR_CODE]}
		}
	}
	return &baseResult
}
func HandleResponse(ctx *gin.Context, err error, response interface{}) {
	baseResult := HandleError(err)
	baseResult.Data = response
	ctx.JSON(http.StatusOK, baseResult)
}
