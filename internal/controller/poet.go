package controller

import (
	"context"
	v1 "poetry/api/v1"
	"poetry/internal/model"
	"poetry/internal/service"
)

var Poet = cPoet{}

type cPoet struct{}

func (c *cPoet) Get(ctx context.Context, req *v1.PoetGetReq) (res *v1.PoetGetRes, err error) {
	out, err := service.Poet().Get(ctx, &model.PoetGetInput{
		Page:      req.Page,
		PageSize:  req.PageSize,
		DynastyId: req.DynastyId,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.PoetGetRes{
		List:  out.List,
		Total: out.Total,
	}
	return
}
