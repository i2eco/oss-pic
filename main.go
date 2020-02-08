package main

import (
	"github.com/goecology/muses"
	"github.com/goecology/muses/pkg/cmd"
	musgin "github.com/goecology/muses/pkg/server/gin"
	"github.com/goecology/muses/pkg/server/stat"
	"github.com/goecology/oss-pic/app/pkg/conf"
	"github.com/goecology/oss-pic/app/pkg/mus"
	"github.com/goecology/oss-pic/app/router"
)

func main() {
	app := muses.Container(
		cmd.Register,
		stat.Register,
		musgin.Register,
	)
	app.SetGinRouter(router.InitRouter)
	app.PreRun(mus.Init, conf.Init)
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
