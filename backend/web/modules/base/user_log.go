package base

import (
	"backend/g"
	"backend/models/base"
)

type UserLog struct {
	*g.X
}

func NewUserLog() *UserLog {
	r := &UserLog{X: g.NewX(&base.UserLog{})}
	r.Tool = []*g.Tool{}
	r.Option = [][]any{}
	return r
}
