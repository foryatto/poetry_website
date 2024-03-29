// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"poetry/internal/model"
)

type IPoet interface {
	Get(ctx context.Context, in *model.PoetGetInput) (out *model.PoetGetOutput, err error)
	GetById(ctx context.Context, in string) (out *model.Poet, err error)
}

var localPoet IPoet

func Poet() IPoet {
	if localPoet == nil {
		panic("implement not found for interface IPoet, forgot register?")
	}
	return localPoet
}

func RegisterPoet(i IPoet) {
	localPoet = i
}
