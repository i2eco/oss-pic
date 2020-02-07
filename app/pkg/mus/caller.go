package mus

import (
	"github.com/gin-gonic/gin"
	"github.com/goecology/muses/pkg/logger"
	musgin "github.com/goecology/muses/pkg/server/gin"
)

var (
	Logger *logger.Client
	Gin    *gin.Engine
)

// Init 初始化muses相关容器
func Init() error {
	Gin = musgin.Caller()
	Logger = logger.Caller("system")
	return nil
}
