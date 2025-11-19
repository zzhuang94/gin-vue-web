package models

import (
	"backend/g"

	"xorm.io/xorm"
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

func (r *RoleAction) Save(sess *xorm.Session) error {
	return r.SaveBean(sess, r)
}

func (r *RoleAction) Delete(sess *xorm.Session) error {
	return r.DeleteBean(sess, r)
}
