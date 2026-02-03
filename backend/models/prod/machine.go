package prod

import "backend/g"

type Machine struct {
	g.Model `xorm:"extends"`

	Name   string `xorm:"name" json:"name"`
	Remark string `xorm:"remark" json:"remark"`
}

func (Machine) TableName() string {
	return "machine"
}

func (Machine) New() g.ModelX {
	return &Machine{}
}

func (m *Machine) Save(sess *g.Sess) error {
	return m.SaveBean(sess, m)
}

func (m *Machine) Delete(sess *g.Sess) error {
	return m.DeleteBean(sess, m)
}
