package prod

import "backend/g"

type Category struct {
	g.Model `xorm:"extends"`

	Name   string `xorm:"name" json:"name"`
	Remark string `xorm:"remark" json:"remark"`
}

func (Category) TableName() string {
	return "category"
}

func (Category) New() g.ModelX {
	return &Category{}
}

func (c *Category) Save(sess *g.Sess) error {
	return c.SaveBean(sess, c)
}

func (c *Category) Delete(sess *g.Sess) error {
	return c.DeleteBean(sess, c)
}
