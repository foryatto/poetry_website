package api

import (
	"backend/app/model"
	"backend/app/shared"
	"backend/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

var Dynasty = dynastyApi{}

type dynastyApi struct{}

func (d *dynastyApi) QueryAll(r *ghttp.Request) {
	queryResult, err := service.Dynasty.QueryAll()
	if err != nil {
		errorCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errorCode, model.ERR_MAP[errorCode])
	}
	shared.Response(r, queryResult)
}
