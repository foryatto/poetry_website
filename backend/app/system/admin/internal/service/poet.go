package service

import (
	"backend/app/dao"
	"backend/app/model"
	"backend/app/system/admin/internal/define"
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var Poet = poetService{}

type poetService struct{}

func (poet *poetService) Query(param *define.PoetQueryParam) (interface{}, error) {
	sql := dao.Poet.Ctx(context.TODO())
	if param.Page >= 0 || param.PageSize >= 1 {
		sql = sql.Page(param.Page, param.PageSize)
	}
	if param.OrderBy != "" {
		sql = sql.Order(param.OrderBy)
	}
	if param.DynastyId != "" {
		sql = sql.Where(dao.Poet.Columns.DynastyId, param.DynastyId)
	}

	count, err := sql.Count()
	if err != nil {
		g.Log().Line().Warning("poet count error:", err)
		return nil, err
	}
	var poetsInfo []*model.Poet
	err = sql.Scan(&poetsInfo)
	if err != nil {
		g.Log().Line().Warning("poet query error:", err)
		return nil, err
	}
	if len(poetsInfo) == 0 {
		return nil, nil
	}
	var result = make([]*define.PoetQueryResp, len(poetsInfo))
	for i := 0; i < len(result); i++ {
		result[i] = new(define.PoetQueryResp)
		err := gconv.Scan(poetsInfo[i], result[i])
		if err != nil {
			g.Log().Line().Warning(err)
			return nil, err
		}
		dynastyInfo, err := Dynasty.GetById(poetsInfo[i].DynastyId)
		if err != nil {
			g.Log().Line().Warning("get dynasty error:", err)
		}
		result[i].Dynasty = dynastyInfo
	}

	return g.Map{
		"total": count,
		"items": result,
	}, nil
}

func (poet *poetService) GetById(id string) (*define.PoetQueryResp, error) {
	var poetInfo model.Poet
	err := dao.Poet.Ctx(context.TODO()).Scan(&poetInfo, dao.Poet.Columns.Id, id)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	dynasty, err := Dynasty.GetById(poetInfo.DynastyId)
	if err != nil {
		return nil, err
	}
	result := new(define.PoetQueryResp)
	err = gconv.Scan(poetInfo, result)
	if err != nil {
		g.Log().Line().Warning("struct change failed:", err)
		return nil, err
	}
	result.Dynasty = dynasty
	return result, nil
}
