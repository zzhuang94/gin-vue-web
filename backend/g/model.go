package g

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type Sess struct {
	*xorm.Session
	Ctx *gin.Context
}
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
	tableName := bean.TableName()
	if _, ok := Ops[tableName]; !ok {
		return
	}
	oldStr := formatDataStr(dataOld)
	newStr := formatDataStr(dataNew)
	if oldStr == newStr {
		return
	}
	log := &Log{
		Uuid:      sess.Ctx.GetString("op_uuid"),
		Op:        op,
		DataTable: tableName,
		DataId:    m.Id,
		DataOld:   oldStr,
		DataNew:   newStr,
	}
	BaseDB.Insert(log)

	sess.Ctx.Set("log_exists", true)
}

func formatDataStr(str string) string {
	if str == "{}" {
		return str
	}
	m := make(map[string]string)
	json.Unmarshal([]byte(str), &m)
	delete(m, "created")
	delete(m, "updated")
	bs, _ := json.Marshal(m)
	return string(bs)
}
