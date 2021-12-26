package api

import (
	"backend/app/model"
	"backend/app/shared"
	"backend/app/system/admin/internal/define"
	"backend/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

var Poet = poetApi{}

type poetApi struct{}

func (poet *poetApi) Query(r *ghttp.Request) {
	var param *define.PoetQueryParam
	shared.Parse(r, &param)
	queryResult, err := service.Poet.Query(param)
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, queryResult)
}
