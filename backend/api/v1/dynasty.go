package v1

import (
	"poetry/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DynastyIndexReq struct {
	g.Meta `path:"/dynasty" tags:"Dynasty" method:"get" summary:"获取所有朝代列表"`
}

type DynastyIndexRes struct {
	List  []model.DynastyItem `json:"list"`
	Total int                 `json:"total"`
}
