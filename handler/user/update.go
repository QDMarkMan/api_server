package user

import (
	"strconv"

	. "github.com/demos/api_server/handler"
	"github.com/demos/api_server/model"
	"github.com/demos/api_server/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lexkong/log"
)

func Update(c *gin.Context) {
	log.Info("Call update user")
	userId, _ := strconv.Atoi(c.Param("id"))
	var u model.UserModel
	var err error
	contentType := c.Request.Header.Get("Content-Type")
	// c.Bind 可以解析参数
	switch contentType {
	case "application/json":
		err = c.BindJSON(&u)
	case "application/x-www-form-urlencoded":
		err = c.BindWith(&u, binding.Form)
	}

	if err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.ID = uint64(userId)

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
