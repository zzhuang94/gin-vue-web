package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
)

type Material struct {
	*frm.X
}

func NewMaterial() *Material {
	r := &Material{X: frm.NewX(&prod.Material{})}
	r.DB = g.CoreDB
	r.Dump = true
	return r
}
