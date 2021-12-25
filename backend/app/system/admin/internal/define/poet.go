package define

import "backend/app/shared"

type QueryPoetParam struct {
	shared.StandReqParam
	DynastyId string `json:"dynastyId"`
}

type QueryPoetResp struct {
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Profile string            `json:"profile"`
	Dynasty *QueryDynastyResp `json:"dynasty"`
}
