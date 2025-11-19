package models

import (
	"backend/g"

	"xorm.io/xorm"
)

type RoleUser struct {
	g.Model  `xorm:"extends"`
	RoleId   int    `xorm:"role_id" json:"role_id,string"`
	Username string `xorm:"username" json:"username"`
	Remark   string `xorm:"remark" json:"remark"`
	Status   int    `xorm:"status" json:"status,string"`
}

func (RoleUser) TableName() string {
	return "role_user"
}

func (RoleUser) New() g.ModelX {
	return &RoleUser{}
}

func (r *RoleUser) Save(sess *xorm.Session) error {
	return r.SaveBean(sess, r)
}

func (r *RoleUser) Delete(sess *xorm.Session) error {
	return r.DeleteBean(sess, r)
}
