package base

import (
	"backend/g"
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

func (u *UserLog) Save(sess *g.Sess) error {
	return u.SaveBean(sess, u)
}

func (u *UserLog) Delete(sess *g.Sess) error {
	return u.DeleteBean(sess, u)
}
