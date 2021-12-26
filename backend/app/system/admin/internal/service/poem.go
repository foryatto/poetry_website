package service

import (
	"backend/app/dao"
	"backend/app/model"
	"backend/app/shared"
	"backend/app/system/admin/internal/define"
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"math/rand"
)

var Poem = poemService{}

type poemService struct{}

var todayPoem *define.PoemQueryResp

func (poem *poemService) Today() (*define.PoemQueryResp, error) {
	if todayPoem == nil {
		todayPoem = new(define.PoemQueryResp)
		count, err := dao.Poem.Ctx(context.TODO()).Count()
		if err != nil {
			g.Log().Line().Warning(err)
			return nil, err
		}
		if count <= 0 {
			return nil, nil
		}
		page := rand.Intn(count)
		var randomPoem *model.Poem
		err = dao.Poem.Ctx(context.TODO()).Page(page, 1).Scan(&randomPoem)
		if err != nil {
			return nil, err
		}
		poetInfo, err := Poet.GetById(randomPoem.PoetId)
		err = gconv.Scan(randomPoem, todayPoem)
		if err != nil {
			g.Log().Line().Warning("struct change failed:", err)
			return nil, err
		}
		todayPoem.Poet = poetInfo
	}
	return todayPoem, nil
}

func (poem *poemService) Search(param *define.PoemSearchReq) (interface{}, error) {
	sql := dao.Poem.Ctx(context.TODO())
	if param.Keyword != "" {
		sql = sql.WhereLike(dao.Poem.Columns.Name, "%"+param.Keyword+"%")
	}
	if param.Page >= 0 || param.PageSize >= 1 {
		sql = sql.Page(param.Page, param.PageSize)
	}
	if param.OrderBy != "" {
		sql = sql.Order(param.OrderBy)
	}
	count, err := sql.Count()
	if err != nil {
		g.Log().Line().Warning("poem count error:", err)
		return nil, err
	}

	var poemsInfo []*model.Poem
	err = sql.Scan(&poemsInfo)
	if err != nil {
		g.Log().Line().Warning("poem query error:", err)
		return nil, err
	}
	if len(poemsInfo) == 0 {
		return nil, nil
	}
	var result = make([]*define.PoemQueryResp, len(poemsInfo))
	for i := 0; i < len(result); i++ {
		result[i] = new(define.PoemQueryResp)
		err := gconv.Scan(poemsInfo[i], result[i])
		if err != nil {
			g.Log().Line().Warning("struct change failed:", err)
			return nil, err
		}
		poetInfo, err := Poet.GetById(poemsInfo[i].PoetId)
		if err != nil {
			g.Log().Line().Warning("get poet error:", err)
		}
		result[i].Poet = poetInfo
	}
	return g.Map{
		"total": count,
		"items": result,
	}, nil
}

func (poem *poemService) QueryBrief(param *define.PoemBriefReq) (interface{}, error) {
	if param.PoetId == "" {
		return nil, shared.NewError(model.ERR_INVALID_PARAM)
	}
	sql := dao.Poem.Ctx(context.TODO())
	sql = sql.Where(dao.Poem.Columns.PoetId, param.PoetId)
	if param.Page >= 0 || param.PageSize >= 1 {
		sql = sql.Page(param.Page, param.PageSize)
	}
	if param.OrderBy != "" {
		sql = sql.Order(param.OrderBy)
	}
	count, err := sql.Count()
	if err != nil {
		g.Log().Line().Warning("poem count error:", err)
		return nil, err
	}
	var result []*define.PoemBriefResp
	sql = sql.Fields(
		dao.Poem.Columns.Id,
		dao.Poem.Columns.Name,
		dao.Poem.Columns.Content,
	)
	err = sql.Scan(&result)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}

	return g.Map{
		"total": count,
		"items": result,
	}, nil
}

func (poem *poemService) QueryDetail(param *define.PoemDetailReq) (*define.PoemQueryResp, error) {
	if param.PoemId == "" {
		return nil, shared.NewError(model.ERR_INVALID_PARAM)
	}
	sql := dao.Poem.Ctx(context.TODO())
	sql = sql.Where(dao.Poem.Columns.Id, param.PoemId)

	var poemInfo *model.Poem
	err := sql.Scan(&poemInfo)
	if err != nil {
		g.Log().Line().Warning(err)
		return nil, err
	}
	result := new(define.PoemQueryResp)
	err = gconv.Scan(poemInfo, result)
	if err != nil {
		return nil, err
	}
	poetInfo, err := Poet.GetById(poemInfo.PoetId)
	if err != nil {
		g.Log().Line().Warning(err)
	}
	result.Poet = poetInfo
	return result, nil
}
