package controllers

import (
	"backend/g"
	"backend/modules/res/models"
)

type Service struct {
	*g.XB[*models.Service]
}

func NewService() *Service {
	r := &Service{XB: g.NewXB(&models.Service{})}
	r.DB = g.CoreDB
	return r
}
