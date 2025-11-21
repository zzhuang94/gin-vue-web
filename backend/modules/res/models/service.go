package models

import (
	"backend/g"
)

type Service struct {
	g.Model  `xorm:"extends"`
	Name     string `xorm:"name" json:"name"`
	Business string `xorm:"business" json:"business"`
	Remark   string `xorm:"remark" json:"remark"`
	Owner    string `xorm:"owner" json:"owner"`
	Status   int    `xorm:"status" json:"status,string"`
}

func (Service) TableName() string {
	return "service"
}

func (Service) New() g.ModelX {
	return &Service{}
}

func (s *Service) Save(sess *g.Sess) error {
	return s.SaveBean(sess, s)
}

func (s *Service) Delete(sess *g.Sess) error {
	return s.DeleteBean(sess, s)
}
