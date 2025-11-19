package g

import (
	"time"

	"xorm.io/xorm"
)

type Model struct {
	ID      int       `xorm:"id not null pk autoincr" json:"id,string"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

func (m *Model) SaveBean(sess *xorm.Session, bean any) error {
	if m.ID == 0 {
		_, err := sess.Insert(bean)
		return err
	}
	_, err := sess.ID(m.ID).AllCols().Update(bean)
	return err
}

func (m *Model) DeleteBean(sess *xorm.Session, bean any) error {
	_, err := sess.ID(m.ID).Delete(bean)
	return err
}

func (m *Model) Save(session *xorm.Session) error {
	if m.ID == 0 {
		_, err := session.Insert(m)
		return err
	}
	_, err := session.ID(m.ID).Update(m)
	return err
}

func (m *Model) Delete(session *xorm.Session) error {
	_, err := session.ID(m.ID).Delete(m)
	return err
}
