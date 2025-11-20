package controllers

import (
	"backend/g"
	"backend/modules/base/models"
)

type UserLog struct {
	*g.X
}

func NewUserLog() *UserLog {
	r := &UserLog{X: g.NewX(&models.UserLog{})}
	r.Tool = []*g.Tool{}
	r.Option = [][]any{}
	return r
}
