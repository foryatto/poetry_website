package controller

import (
	"context"
	v1 "poetry/api/v1"
	"poetry/internal/service"
)

var Dynasty = cDynasty{}

type cDynasty struct{}

// Index dynasty list
func (*cDynasty) Index(ctx context.Context, req *v1.DynastyIndexReq) (res *v1.DynastyIndexRes, err error) {
	out, err := service.Dynasty().GetList(ctx)
	if err != nil {
		return res, err
	}
	res = &v1.DynastyIndexRes{
		List:  out.List,
		Total: out.Total,
	}
	return
}
