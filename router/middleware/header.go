package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func NoCache(g *gin.Context) {
	g.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	g.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	g.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	g.Next()
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(g *gin.Context) {
	if g.Request.Method != "OPTIONS" {
		g.Next()
	} else {
		// 跨域请求
		g.Header("Access-Control-Allow-Origin", "*")
		g.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		g.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		g.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		g.Header("Content-Type", "application/json")
		g.AbortWithStatus(200)
	}
}

// Secure is a middleware function that appends security
// and resource access headers.
func Secure(g *gin.Context) {
	g.Header("Access-Control-Allow-Origin", "*")
	g.Header("X-Frame-Options", "DENY")
	g.Header("X-Content-Type-Options", "nosniff")
	g.Header("X-XSS-Protection", "1; mode=block")
	if g.Request.TLS != nil {
		g.Header("Strict-Transport-Security", "max-age=31536000")
	}
}
