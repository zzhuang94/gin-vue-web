package prod

import (
	"backend/g"
	"fmt"
)

type Label struct {
	g.Model `xorm:"extends"`

	Name   string `xorm:"name" json:"name"`
	Remark string `xorm:"remark" json:"remark"`
}

func (Label) TableName() string {
	return "label"
}

func (Label) New() g.ModelX {
	return &Label{}
}

func (l *Label) Save(sess *g.Sess) error {
	if l.Id == 0 {
		exists, err := sess.Where("name = ?", l.Name).Exist(&Label{})
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("该标签已存在")
		}
	}
	return l.SaveBean(sess, l)
}

func (l *Label) Delete(sess *g.Sess) error {
	return l.DeleteBean(sess, l)
}
