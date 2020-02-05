package mus

import (
	"github.com/gin-gonic/gin"
	"github.com/goecology/muses"
	"github.com/goecology/muses/pkg/logger"
	musgin "github.com/goecology/muses/pkg/server/gin"
	"github.com/goecology/muses/pkg/server/stat"
)

var (
	Logger *logger.Client
	Gin    *gin.Engine
)

// Init 初始化muses相关容器
func Init(cfgFile string) {
	if err := muses.Container(
		cfgFile,
		stat.Register,
		musgin.Register,
	); err != nil {
		panic(err)
	}
	Gin = musgin.Caller()
	Logger = logger.Caller("system")
}
