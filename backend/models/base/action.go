package base

import (
	"backend/g"
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

func (a *Action) Save(sess *g.Sess) error {
	return a.SaveBean(sess, a)
}

func (a *Action) Delete(sess *g.Sess) error {
	ra := new(RoleAction)
	has, _ := sess.Where("action_id = ?", a.Id).Get(ra)
	if has {
		ra.Delete(sess)
	}
	nt := new(Navtree)
	has, _ = sess.Where("action_id = ?", a.Id).Get(nt)
	if has {
		nt.Delete(sess)
	}
	return a.DeleteBean(sess, a)
}
