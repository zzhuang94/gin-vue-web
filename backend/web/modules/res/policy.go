package res

import (
	"backend/g"
	"backend/models/res"
)

type Policy struct {
	*g.XB[*res.Policy]
}

func NewPolicy() *Policy {
	r := &Policy{XB: g.NewXB(&res.Policy{})}
	r.DB = g.CoreDB
	return r
}
