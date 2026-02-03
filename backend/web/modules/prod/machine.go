package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
)

type Machine struct {
	*frm.X
}

func NewMachine() *Machine {
	r := &Machine{X: frm.NewX(&prod.Machine{})}
	r.DB = g.CoreDB
	r.Dump = true
	return r
}
