package models

import (
	"backend/g"
)

type Policy struct {
	g.Model   `xorm:"extends"`
	Name      string `xorm:"name" json:"name"`
	Remark    string `xorm:"remark" json:"remark"`
	ServiceId int    `xorm:"service_id" json:"service_id,string"`
	VidcId    int    `xorm:"vidc_id" json:"vidc_id,string"`
	Status    int    `xorm:"status" json:"status,string"`
	Labels    string `xorm:"labels" json:"labels"`
	Args      string `xorm:"args" json:"args"`
	Vers      string `xorm:"vers" json:"vers"`
}

func (Policy) TableName() string {
	return "policy"
}

func (Policy) New() g.ModelX {
	return &Policy{}
}

func (p *Policy) Save(sess *g.Sess) error {
	return p.SaveBean(sess, p)
}

func (p *Policy) Delete(sess *g.Sess) error {
	return p.DeleteBean(sess, p)
}
