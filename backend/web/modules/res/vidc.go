package res

import (
	"backend/g"
	"backend/models/res"
)

type Vidc struct {
	*g.XB[*res.Vidc]
}

func NewVidc() *Vidc {
	r := &Vidc{XB: g.NewXB(&res.Vidc{})}
	r.Option = append([][]any{{"IP管理", "list", "/res/vidc-ip/list-ip"}}, r.Option...)
	r.DB = g.CoreDB
	return r
}
