package models

import (
	"backend/g"

	"xorm.io/xorm"
)

type Action struct {
	g.Model `xorm:"extends"`
	Path    string `xorm:"path" json:"path"`
	Remark  string `xorm:"remark" json:"remark"`
	Green   int    `xorm:"green" json:"green,string"`
	Status  int    `xorm:"status" json:"status,string"`
}

func (Action) TableName() string {
	return "action"
}

func (Action) New() g.ModelX {
	return &Action{}
}

func (a *Action) Save(sess *xorm.Session) error {
	return a.SaveBean(sess, a)
}

func (a *Action) Delete(sess *xorm.Session) error {
	return a.DeleteBean(sess, a)
}
