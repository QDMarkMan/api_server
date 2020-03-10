package handler

import (
	"net/http"

	"github.com/demos/api_server/pkg/errno"
	"github.com/gin-gonic/gin"
)

// Response stuct
type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
// SendResponse send respose for http
func SendResponse(c *gin.Context, err error ,data interface{})  {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code: code,
		Message: message,
		Data: data,
	})
}