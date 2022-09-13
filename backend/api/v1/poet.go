package v1

import (
	"poetry/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type PoetGetReq struct {
	g.Meta `path:"/poet" tags:"Poet" method:"post" summary:"获取诗人列表"`
	CommonPaginationReq
	DynastyId string `json:"dynastyId"`
}

type PoetGetRes struct {
	List  []model.Poet `json:"list"`
	Total int          `json:"total"`
}
