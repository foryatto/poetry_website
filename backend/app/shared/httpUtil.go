package shared

import (
	"backend/app/model"
	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Parse(r *ghttp.Request, param interface{}) {
	err := r.Parse(param)
	if err != nil {
		errorCode := model.ERR_INVALID_PARAM
		ResponseError(r, errorCode, model.ERR_MAP[errorCode])
	}
}

func NewError(errorCode int) error {
	errorDetail, ok := model.ERR_MAP[errorCode]
	if !ok {
		errorDetail = "undefined error code"
	}
	err := gerror.NewCode(gcode.New(errorCode, "", nil), errorDetail)
	return err
}

func Response(r *ghttp.Request, data interface{}) {
	_ = r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func ResponseError(r *ghttp.Request, code int, msg string) {
	_ = r.Response.WriteJsonExit(g.Map{
		"code": code,
		"msg":  msg,
	})
}
