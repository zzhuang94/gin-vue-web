package g

import (
	"encoding/json"
	"time"
)

type ModelX interface {
	New() ModelX
	TableName() string
	Save(sess *Sess) error
	Delete(sess *Sess) error
}
type Model struct {
	Id      int       `xorm:"id not null pk autoincr" json:"id,string"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

func (m *Model) DeleteBean(sess *Sess, bean ModelX) error {
	_, err := sess.ID(m.Id).Delete(bean)
	if err != nil {
		return err
	}
	bs, _ := json.Marshal(bean)
	m.saveLog(sess, bean, -1, string(bs), "{}")
	return nil
}

func (m *Model) SaveBean(sess *Sess, bean ModelX) error {
	if m.Id == 0 {
		return m.createBean(sess, bean)
	}
	return m.updateBean(sess, bean)
}

func (m *Model) createBean(sess *Sess, bean ModelX) error {
	_, err := sess.Insert(bean)
	if err != nil {
		return err
	}
	bs, _ := json.Marshal(bean)
	m.saveLog(sess, bean, 1, "{}", string(bs))
	return nil
}

func (m *Model) updateBean(sess *Sess, bean ModelX) error {
	old := bean.New()
	sess.ID(m.Id).Get(old)
	oldBs, _ := json.Marshal(old)
	newBs, _ := json.Marshal(bean)
	_, err := sess.ID(m.Id).AllCols().Update(bean)
	if err != nil {
		return err
	}
	m.saveLog(sess, bean, 0, string(oldBs), string(newBs))
	return nil
}

func (m *Model) saveLog(sess *Sess, bean ModelX, op int, dataOld, dataNew string) {
	log := &Log{
		Uuid:      sess.Ctx.GetString("op_uuid"),
		Op:        op,
		DataTable: bean.TableName(),
		DataId:    m.Id,
		DataOld:   dataOld,
		DataNew:   dataNew,
	}
	BaseDB.Insert(log)
}
