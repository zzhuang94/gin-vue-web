package base

import (
	"backend/g"
)

type RoleAction struct {
	g.Model  `xorm:"extends"`
	RoleId   int `xorm:"role_id" json:"role_id,string"`
	ActionId int `xorm:"action_id" json:"action_id,string"`
}

func (RoleAction) TableName() string {
	return "role_action"
}

func (RoleAction) New() g.ModelX {
	return &RoleAction{}
}

func (r *RoleAction) Save(sess *g.Sess) error {
	return r.SaveBean(sess, r)
}

func (r *RoleAction) Delete(sess *g.Sess) error {
	return r.DeleteBean(sess, r)
}
