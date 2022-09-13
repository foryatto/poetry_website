package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"poetry/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 静态文件服务
			// 设置默认提供服务的静态文件目录
			s.SetServerRoot("resource/public")
			// 添加URI与目录路径的映射关系
			s.AddStaticPath("/", "resource/public")

			s.Group("/api", func(group *ghttp.RouterGroup) {
				// 跨域
				group.Middleware(func(r *ghttp.Request) {
					corsOptions := r.Response.DefaultCORSOptions()
					corsOptions.AllowDomain = []string{"localhost:8081"}
					r.Response.CORS(corsOptions)
					r.Middleware.Next()

				})
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Dynasty,
					controller.Poem,
					controller.Poet,
				)
			})
			s.Run()
			return nil
		},
	}
)
