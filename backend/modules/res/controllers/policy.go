package controllers

import (
	"backend/g"
	"backend/modules/res/models"
)

type Policy struct {
	*g.XB[*models.Policy]
}

func NewPolicy() *Policy {
	r := &Policy{XB: g.NewXB(&models.Policy{})}
	r.DB = g.CoreDB
	return r
}
