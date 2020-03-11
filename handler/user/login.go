package user

import (
	"fmt"

	. "github.com/demos/api_server/handler"

	"github.com/demos/api_server/model"
	"github.com/demos/api_server/pkg/auth"
	"github.com/demos/api_server/pkg/errno"
	"github.com/demos/api_server/pkg/token"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.ID, Username: d.Username}, "")
	fmt.Println(t)
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
