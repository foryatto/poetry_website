package service

import "github.com/gogf/gf/net/ghttp"

var Middleware = middleware{}

type middleware struct {
}

// CORS 允许跨域请求中间件
func (m *middleware) CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	//corsOptions.AllowDomain = []string{"localhost", "127.0.0.1"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
