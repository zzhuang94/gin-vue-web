package models

import (
	"backend/g"

	"xorm.io/xorm"
)

type UserLog struct {
	g.Model  `xorm:"extends"`
	Username string `xorm:"username" json:"username"`
	Path     string `xorm:"path" json:"path"`
}

func (UserLog) TableName() string {
	return "user_log"
}

func (UserLog) New() g.ModelX {
	return &UserLog{}
}

func (u *UserLog) Save(sess *xorm.Session) error {
	return u.SaveBean(sess, u)
}

func (u *UserLog) Delete(sess *xorm.Session) error {
	return u.DeleteBean(sess, u)
}
