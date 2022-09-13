package v1

import (
	"poetry/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type PoemTodayReq struct {
	g.Meta `path:"/poem/today" tags:"Poem" method:"get" summary:"获取今日诗词"`
}

type PoemTodayRes struct {
	Poem *model.Poem `json:"poem"`
}

type PoemSearchReq struct {
	g.Meta  `path:"/poem/search" tags:"Poem" method:"post" summary:"搜索诗词"`
	KeyWord string `json:"keyWord"`
	CommonPaginationReq
}

type PoemSearchRes struct {
	List  []model.PoemBriefItem `json:"list"`
	Total int                   `json:"total"`
}

type PoemBriefReq struct {
	g.Meta `path:"/poem/brief" tags:"Poem" method:"post" summary:"获取诗词简介"`
	PoetId string `json:"poetId"`
	CommonPaginationReq
}

type PoemBriefRes struct {
	List  []model.PoemBriefItem `json:"list"`
	Total int                   `json:"total"`
}

type PoemDetailReq struct {
	g.Meta `path:"/poem/detail" tags:"Poem" method:"post" summary:"获取诗词详情"`
	PoemId string `json:"poemId"`
}
type PoemDetailRes struct {
	Poem *model.Poem `json:"poem"`
}
