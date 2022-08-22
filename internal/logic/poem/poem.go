package poem

import (
	"context"
	"math/rand"
	"poetry/internal/dao"
	"poetry/internal/logic/poet"
	"poetry/internal/model"
	"poetry/internal/service"
)

type sPoem struct{}

func init() {
	service.RegisterPoem(New())
}

func New() *sPoem {
	return &sPoem{}
}

func (*sPoem) Today(ctx context.Context) (out *model.Poem, err error) {
	sql := dao.Poem.Ctx(ctx)
	count, err := sql.Count()
	if err != nil {
		return out, err
	}
	page := rand.Intn(count)
	err = sql.Page(page, 1).Scan(&out)
	if err != nil {
		return out, err
	}

	out.Writer, err = poet.New().GetById(ctx, out.PoetId)
	if err != nil {
		return out, err
	}
	return
}

func (*sPoem) GetBrief(ctx context.Context, in *model.PoemBriefInput) (out *model.PoemBriefOutput, err error) {
	sql := dao.Poem.Ctx(ctx).Where(dao.Poem.Columns().PoetId, in.PoetId)
	count, err := sql.Count()
	if err != nil {
		return out, err
	}
	out = &model.PoemBriefOutput{Total: count}
	err = sql.Page(in.Page, in.PageSize).Scan(&out.List)
	return out, nil
}

func (*sPoem) GetDetail(ctx context.Context, in *model.PoemDetailInput) (out *model.PoemDetailOutput, err error) {
	sql := dao.Poem.Ctx(ctx).Where(dao.Poem.Columns().Id, in.PoemId)

	out = &model.PoemDetailOutput{}
	err = sql.Scan(&out.Poem)
	if err != nil {
		return nil, err
	}
	out.Poem.Writer, err = poet.New().GetById(ctx, out.Poem.PoetId)
	return out, nil
}

func (*sPoem) Search(ctx context.Context, in *model.PoemSearchInput) (out *model.PoemSearchOutput, err error) {
	sql := dao.Poem.Ctx(ctx).WhereLike(dao.Poem.Columns().Content, "%"+in.KeyWord+"%")
	count, err := sql.Count()
	if err != nil {
		return out, err
	}
	out = &model.PoemSearchOutput{Total: count}
	err = sql.Page(in.Page, in.PageSize).Scan(&out.List)
	return out, nil
}
