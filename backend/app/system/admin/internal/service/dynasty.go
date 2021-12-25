package service

import (
	"backend/app/dao"
	"backend/app/system/admin/internal/define"
	"context"
	"github.com/gogf/gf/frame/g"
)

var Dynasty = dynastyService{}

type dynastyService struct{}

func (ds *dynastyService) QueryAll() (*[]define.QueryDynastyResp, error) {
	var queryRes []define.QueryDynastyResp
	err := dao.Dynasty.Ctx(context.TODO()).Order(dao.Dynasty.Columns.Time).Scan(&queryRes)
	if err != nil {
		g.Log().Line().Warning("query all dynasties err: ", err)
		return nil, err
	}
	return &queryRes, nil
}

func (ds *dynastyService) GetById(dynastyId string) (*define.QueryDynastyResp, error) {
	var queryRes define.QueryDynastyResp
	err := dao.Dynasty.Ctx(context.TODO()).Scan(&queryRes, g.Map{
		dao.Dynasty.Columns.Id: dynastyId,
	})
	if err != nil {
		g.Log().Line().Warning("query dynasty err: ", err)
		return nil, err
	}
	return &queryRes, nil
}
