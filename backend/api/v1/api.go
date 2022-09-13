package v1

type CommonPaginationReq struct {
	Page     int `json:"page" d:"1"  v:"min:1#分页号码错误"     dc:"分页号码，默认1"`
	PageSize int `json:"pageSize" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}

type CommonPaginationRes struct {
	Total int `dc:"总数"`
}
