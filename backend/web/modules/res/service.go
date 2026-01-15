package res

import (
	"backend/g"
	"backend/models/res"
	"backend/web/frm"
)

type Service struct {
	*frm.XB[*res.Service]
}

func NewService() *Service {
	r := &Service{XB: frm.NewXB(&res.Service{})}
	r.DB = g.CoreDB
	return r
}
