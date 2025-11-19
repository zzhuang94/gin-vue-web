package models

import (
	"backend/g"
	"fmt"

	"github.com/asaskevich/govalidator"
	"xorm.io/xorm"
)

type Ip struct {
	g.Model `xorm:"extends"`

	Ip        string `xorm:"ip" json:"ip"`
	Type      string `xorm:"type" json:"type"`
	Sn        string `xorm:"sn" json:"sn"`
	Isp       string `xorm:"isp" json:"isp"`
	Mac       string `xorm:"mac" json:"mac"`
	Bandwidth int    `xorm:"bandwidth" json:"bandwidth,string"`
	Status    int    `xorm:"status" json:"status,string"`
}

func (Ip) TableName() string {
	return "ip"
}

func (Ip) New() g.ModelX {
	return &Ip{}
}

func (i *Ip) Save(sess *xorm.Session) error {
	if !govalidator.IsIP(i.Ip) {
		return fmt.Errorf("IP地址格式错误")
	}
	if i.Type == "IPv4" && !govalidator.IsIPv4(i.Ip) {
		return fmt.Errorf("IPv4地址格式错误")
	}
	if i.Type == "v6" && !govalidator.IsIPv6(i.Ip) {
		return fmt.Errorf("IPv6地址格式错误")
	}
	return i.SaveBean(sess, i)
}

func (i *Ip) Delete(sess *xorm.Session) error {
	vis := []*VidcIp{}
	sess.Where("ip_id = ?", i.ID).Find(&vis)
	for _, vi := range vis {
		vi.Delete(sess)
	}
	return i.DeleteBean(sess, i)
}
