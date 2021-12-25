package admin

import (
	"backend/app/system/admin/internal/api"
	"backend/app/system/admin/internal/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Init() {
	s := g.Server()
	s.Group("/api/v1/poetry", func(group *ghttp.RouterGroup) {

		group.Middleware(service.Middleware.CORS)

		group.Group("/dynasties", func(group *ghttp.RouterGroup) {
			group.GET("/", api.Dynasty.QueryAll)
		})

		group.Group("/poets", func(group *ghttp.RouterGroup) {
			group.POST("/", api.Poet.Query)
		})

		group.Group("/poems", func(group *ghttp.RouterGroup) {
			group.POST("/today", api.Poem.Today)
			group.POST("/search", api.Poem.Search)
			group.POST("/brief", api.Poem.QueryBrief)
			group.POST("/detail", api.Poem.QueryDetail)
		})
	})
}
