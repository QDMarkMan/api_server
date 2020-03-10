package user

import (
	. "github.com/demos/api_server/handler"
	"github.com/demos/api_server/pkg/errno"
	"github.com/demos/api_server/service"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	var r ListRequest

	if err := c.Bind(r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)

	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
