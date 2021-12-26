package api

import (
	"backend/app/model"
	"backend/app/shared"
	"backend/app/system/admin/internal/define"
	"backend/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

var Poem = poemApi{}

type poemApi struct{}

func (poet *poemApi) Today(r *ghttp.Request) {
	queryResult, err := service.Poem.Today()
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, queryResult)
}

func (poet *poemApi) QueryBrief(r *ghttp.Request) {
	var param *define.PoemBriefReq
	shared.Parse(r, &param)
	queryResult, err := service.Poem.QueryBrief(param)
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, queryResult)
}

func (poet *poemApi) QueryDetail(r *ghttp.Request) {
	var param *define.PoemDetailReq
	shared.Parse(r, &param)
	queryResult, err := service.Poem.QueryDetail(param)
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, queryResult)
}

func (poet *poemApi) Search(r *ghttp.Request) {
	var param *define.PoemSearchReq
	shared.Parse(r, &param)
	searchResult, err := service.Poem.Search(param)
	if err != nil {
		errCode := model.ERR_DATABASE_READ
		shared.ResponseError(r, errCode, model.ERR_MAP[errCode])
	}
	shared.Response(r, searchResult)
}
