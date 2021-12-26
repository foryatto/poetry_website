package define

import "backend/app/shared"

type PoetQueryParam struct {
	shared.StandReqParam
	DynastyId string `json:"dynastyId"`
}

type PoetQueryResp struct {
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Profile string            `json:"profile"`
	Dynasty *DynastyQueryResp `json:"dynasty"`
}
