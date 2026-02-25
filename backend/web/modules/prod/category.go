package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
)

type Category struct {
	*frm.X
}

func NewCategory() *Category {
	r := &Category{X: frm.NewX(&prod.Category{})}
	r.DB = g.CoreDB
	r.Dump = true
	return r
}
