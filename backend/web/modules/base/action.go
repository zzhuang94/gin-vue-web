package base

import (
	"backend/g"
	"backend/models/base"
)

type Action struct {
	*g.XB[*base.Action]
}

func NewAction() *Action {
	a := &Action{XB: g.NewXB(&base.Action{})}
	a.BatchDelete = false
	a.Tool = []*g.Tool{}
	a.Option = [][]any{{"编 辑", "edit", "edit", "modal", []string{"id"}}}
	return a
}
