package res

import (
	"backend/g"
	"backend/models/res"
)

type Service struct {
	*g.XB[*res.Service]
}

func NewService() *Service {
	r := &Service{XB: g.NewXB(&res.Service{})}
	r.DB = g.CoreDB
	return r
}
