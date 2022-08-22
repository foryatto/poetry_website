package poet

import (
	"context"
	"poetry/internal/dao"
	"poetry/internal/model"
	"poetry/internal/service"
)

type sPoet struct{}

func init() {
	service.RegisterPoet(New())
}

func New() *sPoet {
	return &sPoet{}
}

func (*sPoet) Get(ctx context.Context, in *model.PoetGetInput) (out *model.PoetGetOutput, err error) {
	sql := dao.Poet.Ctx(ctx).Where(dao.Poet.Columns().DynastyId, in.DynastyId)
	count, err := sql.Count()
	if err != nil {
		return nil, err
	}
	out = &model.PoetGetOutput{Total: count}
	err = sql.Page(in.Page, in.PageSize).Scan(&out.List)
	if err != nil {
		return nil, err
	}
	length := len(out.List)
	for i := 0; i < length; i++ {
		dao.Dynasty.Ctx(ctx).Where(dao.Dynasty.Columns().Id, out.List[i].DynastyId).Scan(&out.List[i].Dynasty)
	}

	return
}

func (*sPoet) GetById(ctx context.Context, in string) (out *model.Poet, err error) {
	sql := dao.Poet.Ctx(ctx).Where(dao.Poet.Columns().Id, in)
	err = sql.Scan(&out)
	if err != nil {
		return out, err
	}
	err = dao.Dynasty.Ctx(ctx).Where(dao.Dynasty.Columns().Id, out.DynastyId).Scan(&out.Dynasty)

	return
}
