package handlers

import (
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testProject/core/server"
	"testing"
)

func TestPing(t *testing.T) {
	uri := "/user/ping"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", uri, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response server.BaseResponse
	err := jsoniter.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	t.Log(string(w.Body.Bytes()))
	assert.Equal(t, server.STATUS_SERVER_SUCCESS_CODE, response.Status)
	assert.Equal(t, server.GetErrorReason(server.STATUS_SERVER_SUCCESS_CODE), response.Reason)
}
