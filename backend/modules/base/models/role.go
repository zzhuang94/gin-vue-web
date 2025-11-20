package models

import (
	"backend/g"

	"xorm.io/xorm"
)

type Role struct {
	g.Model `xorm:"extends"`
	Name    string `xorm:"name" json:"name"`
	Remark  string `xorm:"remark" json:"remark"`
	Icon    string `xorm:"icon" json:"icon"`
	Status  int    `xorm:"status" json:"status,string"`
}

func (Role) TableName() string {
	return "role"
}

func (Role) New() g.ModelX {
	return &Role{}
}

func (r *Role) Save(sess *xorm.Session) error {
	return r.SaveBean(sess, r)
}

func (r *Role) Delete(sess *xorm.Session) error {
	ra := new(RoleAction)
	has, _ := sess.Where("role_id = ?", r.ID).Get(ra)
	if has {
		ra.Delete(sess)
	}
	ru := new(RoleUser)
	has, _ = sess.Where("role_id = ?", r.ID).Get(ru)
	if has {
		ru.Delete(sess)
	}
	return r.DeleteBean(sess, r)
}

func (r *Role) GetActionIds() []string {
	ids := []string{}
	sql := "SELECT action_id FROM role_action WHERE role_id = ?"
	g.BaseDB.SQL(sql, r.ID).Find(&ids)
	return ids
}
