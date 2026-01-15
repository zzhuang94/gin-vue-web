package res

import (
	"backend/g"
	"backend/models/res"
	"backend/web/frm"
)

type Policy struct {
	*frm.XB[*res.Policy]
}

func NewPolicy() *Policy {
	r := &Policy{XB: frm.NewXB(&res.Policy{})}
	r.DB = g.CoreDB
	return r
}
