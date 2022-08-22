package controller

import (
	"context"
	v1 "poetry/api/v1"
	"poetry/internal/model"
	"poetry/internal/service"
)

var Poem = cPoem{}

type cPoem struct{}

func (*cPoem) Today(ctx context.Context, req *v1.PoemTodayReq) (res *v1.PoemTodayRes, err error) {
	out, err := service.Poem().Today(ctx)
	res = &v1.PoemTodayRes{
		Poem: out,
	}
	return
}

func (*cPoem) GetBrief(ctx context.Context, req *v1.PoemBriefReq) (res *v1.PoemBriefRes, err error) {
	out, err := service.Poem().GetBrief(ctx, &model.PoemBriefInput{
		PoetId:   req.PoetId,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return res, err
	}
	res = &v1.PoemBriefRes{
		List:  out.List,
		Total: out.Total,
	}
	return
}

func (*cPoem) GetDetail(ctx context.Context, req *v1.PoemDetailReq) (res *v1.PoemDetailRes, err error) {
	out, err := service.Poem().GetDetail(ctx, &model.PoemDetailInput{
		PoemId: req.PoemId,
	})
	if err != nil {
		return res, err
	}
	res = &v1.PoemDetailRes{
		Poem: out.Poem,
	}
	return
}

func (*cPoem) Search(ctx context.Context, req *v1.PoemSearchReq) (res *v1.PoemSearchRes, err error) {
	out, err := service.Poem().Search(ctx, &model.PoemSearchInput{
		KeyWord:  req.KeyWord,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return res, err
	}
	res = &v1.PoemSearchRes{
		List:  out.List,
		Total: out.Total,
	}
	return
}
