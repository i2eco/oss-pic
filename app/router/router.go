package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goecology/oss-pic/app/pkg/mus"
)

func InitRouter() *gin.Engine {
	r := mus.Gin
	r.GET("/", func(context *gin.Context) {
		context.String(200, "%s", "image server")
	})
	r.NoRoute(Info)
	return r
}
