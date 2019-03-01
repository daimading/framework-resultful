package handlers

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"testProject/core/server"
	"testProject/other"
)

type FileUploadRequest struct {
}
type FileUploadResponse struct {
	FileId int    `json:"file_id" yaml:"file_id"`
	Url    string `json:"url" yaml:"url"`
}

func (req *FileUploadRequest) uploadFile(file multipart.File, fileName string, size int64) (fileId int, url string, err error) {
	return
}
func UploadFile(ctx *gin.Context) {
	var (
		err  error
		resp FileUploadResponse
		req  FileUploadRequest
	)
	defer func() {
		server.HandleResponse(ctx, err, &resp)
	}()

	if err = server.HandleRequest(ctx, &req); err != nil {
		err = server.NewError(server.STATUS_PARAM_PARES_ERROR_CODE)
		return
	}
	file, header, err := ctx.Request.FormFile("upload_file")
	if file == nil || header == nil {
		err = server.NewError(other.FILE_UPLOAD_FAIL_STATUS_CODE)
	}
	fileId, url, err := req.uploadFile(file, header.Filename, header.Size)
	resp.FileId = fileId
	resp.Url = url
}
