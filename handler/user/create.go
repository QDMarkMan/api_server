package user

import (
	"fmt"

	. "github.com/demos/api_server/handler"
	"github.com/demos/api_server/model"
	"github.com/demos/api_server/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty")
	}
	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty")
	}
	return nil
}

// Create user
func Create(c *gin.Context) {
	// request
	var r CreateRequest
	var err error
	contentType := c.Request.Header.Get("Content-Type")
	fmt.Println(contentType)
	// c.Bind 可以解析参数
	switch contentType {
	case "application/json":
		err = c.BindJSON(&r)
	case "application/x-www-form-urlencoded":
		err = c.BindWith(&r, binding.Form)
	}
	if err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	// 参数校验
	if err = r.checkParam(); err != nil {
		SendResponse(c, err, nil)
		return
	}
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}
	// test param query 和通常一样获取string字符串
	// admin2 := c.Query("username")
	// log.Infof("URL username: %s", admin2)

	// desc := c.Query("desc")
	// log.Infof("URL key param desc: %s", desc)

	// contentType := c.GetHeader("content-type")
	// log.Infof("URL content type is: %s", contentType)

	// log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	// if r.Username == "" {
	// 	err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("this is add message")
	// 	SendResponse(c, err, nil)
	// 	return
	// }
	// if errno.IsErrUserNotFound(err) {
	// 	log.Debug("err type is ErrUserNotFound")
	// }
	// if r.Password == "" {
	// 	err = fmt.Errorf("password is empty")
	// 	SendResponse(c, err, nil)
	// 	return
	// }

	if err = u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err = u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err = u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}
	// success content
	SendResponse(c, nil, rsp)
}
