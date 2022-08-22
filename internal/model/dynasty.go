package model

type DynastyItem struct {
	Id   string `json:"id"   `
	Name string `json:"name" `
}

// DynastyGetListOutput 查询朝代列表结果
type DynastyGetListOutput struct {
	List  []DynastyItem `json:"list" description:"列表"`
	Total int           `json:"total" description:"数据总数"`
}


