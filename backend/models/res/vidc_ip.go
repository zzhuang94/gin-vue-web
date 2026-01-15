package res

import (
	"backend/g"
)

type VidcIp struct {
	g.Model `xorm:"extends"`

	VidcId int    `xorm:"vidc_id" json:"vidc_id,string"`
	IpId   int    `xorm:"ip_id" json:"ip_id,string"`
	Weight int    `xorm:"weight" json:"weight,string"`
	Remark string `xorm:"remark" json:"remark"`
	Status int    `xorm:"status" json:"status,string"`
}

func (VidcIp) TableName() string {
	return "vidc_ip"
}

func (VidcIp) New() g.ModelX {
	return &VidcIp{}
}

func (v *VidcIp) Save(sess *g.Sess) error {
	return v.SaveBean(sess, v)
}

func (v *VidcIp) Delete(sess *g.Sess) error {
	return v.DeleteBean(sess, v)
}
