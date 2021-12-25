package app

import (
	"backend/app/system/admin"
	"github.com/gogf/gf/frame/g"
)

func Run() {
	admin.Init()
	g.Server().Run()
}
