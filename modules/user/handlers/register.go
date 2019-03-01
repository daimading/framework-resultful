package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"testProject/core/database"
	"testProject/core/server"
	"testProject/modules/user"
	"testProject/utils"
)

type AdminRequest struct {
	UserName string `json:"user_name"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
type AdminResponse struct {
	UserId int `json:"user_id"`
}

func (req *AdminRequest) register() (userId int, err error) {
	if req.UserName == "" || req.Password == "" {
		return userId, server.NewError(user.STATUS_IMPORT_USERNAME_PASSWORD_CODE)
	}
	db := database.GetDB()
	forumAdmin := user.ForumAdmin{UserName: req.UserName}
	err = db.Model(&forumAdmin).First(&forumAdmin).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return userId, server.NewError(server.STATUS_SERVER_ERROR_CODE)
	} else {
		return userId, server.NewError(user.STATUS_USERNAME_EXITS_CODE)
	}
	salt := utils.GenSalt()
	password := utils.GenHashedPassword(req.Password, salt)
	forumAdmin.Password = password
	forumAdmin.NickName = req.Nickname
	err = db.Save(&forumAdmin).Error
	if err != nil {
		return userId, server.NewError(server.STATUS_SERVER_ERROR_CODE)
	}
	userId = forumAdmin.ID
	return
}
func Register(ctx *gin.Context) {
	var (
		err  error
		req  AdminRequest
		resp AdminResponse
	)
	defer func() {
		server.HandleResponse(ctx, err, &resp)
	}()
	err = server.HandleRequest(ctx, &req)
	if err != nil {
		return
		logrus.Error("request error", err.Error())
	}
	resp.UserId, err = req.register()
	if err != nil {
		return
	}
}
