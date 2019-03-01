package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testProject/core/server"
)

type PingRequest struct {
	Ping string `json:"ping"`
}

func Ping(ctx *gin.Context) {
	var (
		err  error
		req  PingRequest
		data = make(map[string]interface{}, 1)
	)
	defer func() {
		server.HandleResponse(ctx, err, &data)
	}()
	err = server.HandleRequest(ctx, &req)
	if err != nil {
		fmt.Println("request server error")
	}

	data["method"] = ctx.Request.Method
	data["pong"] = req.Ping
	data["hello"] = "hello word !!!!"
}
