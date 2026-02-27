package prod

import (
	"backend/g"
	"fmt"
)

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
	if m.Id == 0 {
		exists, err := sess.Where("name = ?", m.Name).Exist(&Material{})
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("该材质已存在")
		}
	}
	return m.SaveBean(sess, m)
}

func (m *Material) Delete(sess *g.Sess) error {
	return m.DeleteBean(sess, m)
}
