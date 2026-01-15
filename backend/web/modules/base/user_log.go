package base

import (
	"backend/models/base"
	"backend/web/frm"
)

type UserLog struct {
	*frm.X
}

func NewUserLog() *UserLog {
	r := &UserLog{X: frm.NewX(&base.UserLog{})}
	r.Tool = []*frm.Tool{}
	r.Option = [][]any{}
	return r
}
