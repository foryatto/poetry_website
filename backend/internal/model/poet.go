package model

type PoetItem struct {
	Id        string `json:"id"        `
	Name      string `json:"name"      `
	Profile   string `json:"profile"   `
	DynastyId string `json:"dynastyId" `
}

type Poet struct {
	PoetItem
	Dynasty *DynastyItem `json:"dynasty"`
}

type PoetGetInput struct {
	Page      int    `json:"page" description:"分页码"`
	PageSize  int    `json:"pageSize" description:"分页数量"`
	DynastyId string `json:"dynastyId"`
}
type PoetGetOutput struct {
	List  []Poet `json:"List"`
	Total int    `json:"total"`
}
