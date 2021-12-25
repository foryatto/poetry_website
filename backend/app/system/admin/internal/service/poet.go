package service

import (
	"backend/app/dao"
	"backend/app/model"
	"backend/app/system/admin/internal/define"
	"context"
	"github.com/gogf/gf/frame/g"
)

var Poet = poetService{}

type poetService struct{}

func (poet *poetService) Query(param *define.QueryPoetParam) (interface{}, error) {
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
	var poetsInfo []model.Poet
	err = sql.Scan(&poetsInfo)
	if err != nil {
		g.Log().Line().Warning("poet query error:", err)
		return nil, err
	}
	if len(poetsInfo) == 0 {
		return nil, nil
	}
	var result = make([]define.QueryPoetResp, len(poetsInfo))
	for i := 0; i < len(result); i++ {
		result[i] = define.QueryPoetResp{
			Id:      poetsInfo[i].Id,
			Name:    poetsInfo[i].Name,
			Profile: poetsInfo[i].Profile,
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
