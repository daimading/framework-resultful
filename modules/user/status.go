package user

import "testProject/core/server"

const (
	STATUS_IMPORT_USERNAME_PASSWORD_CODE = iota - 19999
	STATUS_USERNAME_EXITS_CODE
)

func init() {
	server.RegisterErrorReason(STATUS_IMPORT_USERNAME_PASSWORD_CODE, "请输入用户名或密码")
	server.RegisterErrorReason(STATUS_USERNAME_EXITS_CODE, "用户名已经存在")
}
