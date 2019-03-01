package other

import "testProject/core/server"

const (
	FILE_UPLOAD_FAIL_STATUS_CODE = iota + 4000
)

func init() {
	server.RegisterErrorReason(FILE_UPLOAD_FAIL_STATUS_CODE, "请选择需要上传的文件")
}
