package models

import (
	"backend/g"

	"xorm.io/xorm"
)

type Vidc struct {
	g.Model `xorm:"extends"`

	Name     string `xorm:"name" json:"name"`
	Type     string `xorm:"type" json:"type"`
	Location string `xorm:"location" json:"location"`
	Status   int    `xorm:"status" json:"status,string"`
}

func (Vidc) TableName() string {
	return "vidc"
}

func (Vidc) New() g.ModelX {
	return &Vidc{}
}

func (v *Vidc) Save(sess *xorm.Session) error {
	return v.SaveBean(sess, v)
}

func (v *Vidc) Delete(sess *xorm.Session) error {
	vis := []*VidcIp{}
	sess.Where("vidc_id = ?", v.ID).Find(&vis)
	for _, vi := range vis {
		vi.Delete(sess)
	}
	return v.DeleteBean(sess, v)
}
