package models

import (
	"backend/g"

	"xorm.io/xorm"
)

type User struct {
	g.Model `xorm:"extends"`

	Username string `xorm:"username" json:"username"`
	Email    string `xorm:"email" json:"email"`
	CnName   string `xorm:"cn_name" json:"cn_name"`
	Avatar   []byte `xorm:"avatar BLOB" json:"-"`
	Fold     int    `xorm:"fold" json:"fold,string"`
	PageSize int    `xorm:"page_size" json:"page_size,string"`

	Password string `xorm:"password" json:"password"`
}

func (User) TableName() string {
	return "user"
}

func (User) New() g.ModelX {
	return &User{}
}

func (u *User) Save(sess *xorm.Session) error {
	return u.SaveBean(sess, u)
}

func (u *User) Delete(sess *xorm.Session) error {
	ru := new(RoleUser)
	has, _ := sess.Where("username = ?", u.Username).Get(ru)
	if has {
		ru.Delete(sess)
	}
	return u.DeleteBean(sess, u)
}
