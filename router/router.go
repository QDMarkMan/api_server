package router

import (
	"net/http"

	"github.com/demos/api_server/handler/sd"
	"github.com/demos/api_server/handler/user"
	"github.com/demos/api_server/router/middleware"
	"github.com/gin-gonic/gin"
)

//Load loads middlewares routes handlers
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())

	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	// 自定义全局RequestId中间件
	g.Use(middleware.RequestId())
	g.Use(mw...)
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "404 Not Found")
	})
	// test
	g.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to api server",
		})
	})
	g.POST("/login", user.Login)
	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.ChekcHealth)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	u := g.Group("v1/user")
	{
		u.POST("create", user.Create)
		u.GET("list", user.UserList)
	}
	return g
}
