package base

import (
	"backend/models/base"
	"backend/web/frm"
)

type Action struct {
	*frm.XB[*base.Action]
}

func NewAction() *Action {
	a := &Action{XB: frm.NewXB(&base.Action{})}
	a.BatchDelete = false
	a.Tool = []*frm.Tool{}
	a.Option = [][]any{{"编 辑", "edit", "edit", "modal", []string{"id"}}}
	return a
}
