package controllers

import (
	"backend/g"
	"backend/modules/res/models"
)

type Vidc struct {
	*g.XB[*models.Vidc]
}

func NewVidc() *Vidc {
	r := &Vidc{XB: g.NewXB(&models.Vidc{})}
	r.Option = append([][]any{{"IP管理", "list", "/res/vidc-ip/list-ip"}}, r.Option...)
	r.DB = g.CoreDB
	return r
}
