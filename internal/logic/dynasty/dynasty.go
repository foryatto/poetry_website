package dynasty

import (
	"context"
	"poetry/internal/dao"
	"poetry/internal/model"
	"poetry/internal/service"
)

type sDynasty struct{}

func init() {
	service.RegisterDynasty(New())
}

func New() *sDynasty {
	return &sDynasty{}
}

func (s *sDynasty) GetList(ctx context.Context) (out *model.DynastyGetListOutput, err error) {
	sql := dao.Dynasty.Ctx(ctx)

	out = &model.DynastyGetListOutput{}
	out.Total, err = sql.Count()
	if err != nil {
		return out, err
	}
	err = sql.Scan(&out.List)
	return out, nil
}
