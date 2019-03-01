package handlers

import (
	"bytes"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testProject/core/server"
	"testing"
)

func TestRegister(t *testing.T) {
	uri := "/user/register"
	data := make(map[string]interface{})
	data["user_name"] = "test"
	data["password"] = "12345"
	content, _ := jsoniter.Marshal(data)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", uri, bytes.NewReader(content))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var resp server.BaseResponse
	jsoniter.Unmarshal(w.Body.Bytes(), &resp)
	t.Log("response:", resp)
	assert.Equal(t, server.STATUS_SERVER_SUCCESS_CODE, resp.Status)
	assert.Equal(t, server.GetErrorReason(server.STATUS_SERVER_SUCCESS_CODE), resp.Reason)
}
