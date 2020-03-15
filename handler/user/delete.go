package user

import (
	"strconv"

	. "github.com/demos/api_server/handler"
	"github.com/demos/api_server/model"
	"github.com/demos/api_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	// TODO: strconv 需要了解的模块
	userId, _ := strconv.Atoi(c.Param("id"))

	if err := model.Delete(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
