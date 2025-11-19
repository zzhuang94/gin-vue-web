package controllers

import (
	"backend/g"
	"backend/modules/base/models"
)

type Action struct {
	*g.XB[*models.Action]
}

func NewAction() *Action {
	a := &Action{XB: g.NewXB(&models.Action{})}
	a.BatchDelete = false
	a.Tool = []*g.Tool{}
	a.Option = [][]any{{"编 辑", "edit", "edit", "modal", []string{"id"}}}
	return a
}
