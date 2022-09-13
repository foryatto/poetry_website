package model

type PoemItem struct {
	Id      string `json:"id"`
	Form    string `json:"form"`
	Tags    string `json:"tags"`
	Name    string `json:"name"`
	Content string `json:"content"`
	PoetId  string `json:"poetId"  `
}

type Poem struct {
	PoemItem
	Writer *Poet `json:"poet"`
}

type PoemBriefInput struct {
	PoetId   string `json:"poetId"`
	Page     int    `json:"page" description:"分页码"`
	PageSize int    `json:"pageSize" description:"分页数量"`
}

type PoemBriefItem struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type PoemBriefOutput struct {
	List  []PoemBriefItem `json:"list"`
	Total int             `json:"total"`
}

type PoemDetailInput struct {
	PoemId string `json:"poemId"`
}

type PoemDetailOutput struct {
	Poem *Poem `json:"poem"`
}

type PoemSearchInput struct {
	KeyWord  string `json:"keyword"`
	Page     int    `json:"page" description:"分页码"`
	PageSize int    `json:"pageSize" description:"分页数量"`
}

type PoemSearchOutput struct {
	List  []PoemBriefItem `json:"list"`
	Total int             `json:"total"`
}
