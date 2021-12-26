package define

import "backend/app/shared"

type PoemQueryResp struct {
	Id      string         `json:"id"`
	Form    string         `json:"form"`
	Tags    string         `json:"tags"`
	Name    string         `json:"name"`
	Content string         `json:"content"`
	Poet    *PoetQueryResp `json:"poet"`
}

type PoemSearchReq struct {
	Keyword string `json:"keyword"`
	shared.StandReqParam
}

type PoemBriefReq struct {
	PoetId string `json:"poetId"`
	shared.StandReqParam
}

type PoemBriefResp struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type PoemDetailReq struct {
	PoemId string `json:"poemId"`
}
