package prod

import "backend/g"

type Color struct {
	g.Model `xorm:"extends"`

	Name   string `xorm:"name" json:"name"`
	Remark string `xorm:"remark" json:"remark"`
}

func (Color) TableName() string {
	return "color"
}

func (Color) New() g.ModelX {
	return &Color{}
}

func (c *Color) Save(sess *g.Sess) error {
	return c.SaveBean(sess, c)
}

func (c *Color) Delete(sess *g.Sess) error {
	return c.DeleteBean(sess, c)
}
