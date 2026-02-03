package prod

import "backend/g"

type Material struct {
	g.Model `xorm:"extends"`

	Name   string `xorm:"name" json:"name"`
	Remark string `xorm:"remark" json:"remark"`
}

func (Material) TableName() string {
	return "material"
}

func (Material) New() g.ModelX {
	return &Material{}
}

func (m *Material) Save(sess *g.Sess) error {
	return m.SaveBean(sess, m)
}

func (m *Material) Delete(sess *g.Sess) error {
	return m.DeleteBean(sess, m)
}
