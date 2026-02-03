package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
)

type Color struct {
	*frm.X
}

func NewColor() *Color {
	r := &Color{X: frm.NewX(&prod.Color{})}
	r.DB = g.CoreDB
	r.Dump = true
	return r
}
