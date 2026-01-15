package base

import (
	"backend/g"
	"fmt"
	"strconv"
	"strings"
)

type Navtree struct {
	g.Model  `xorm:"extends"`
	ActionId int    `xorm:"action_id" json:"action_id,string"`
	ParentId int    `xorm:"parent_id" json:"parent_id,string"`
	KeyPath  string `xorm:"key_path" json:"key_path"`
	Level    int    `xorm:"level" json:"level,string"`
	Rank     int    `xorm:"rank" json:"rank,string"`
	Name     string `xorm:"name" json:"name"`
	Icon     string `xorm:"icon" json:"icon"`
	Status   int    `xorm:"status" json:"status,string"`
}

func (Navtree) TableName() string {
	return "navtree"
}

func (Navtree) New() g.ModelX {
	return &Navtree{}
}

func (n *Navtree) Save(sess *g.Sess) error {
	parent := &Navtree{}
	sess.Where("id = ?", n.ParentId).Get(parent)
	if parent.Id == 0 {
		return fmt.Errorf("父节点不存在")
	}
	if n.Id > 0 {
		if n.Id == n.ParentId {
			return fmt.Errorf("不能指定自己为父节点")
		}
		pids := strings.SplitSeq(n.KeyPath, "_")
		for pid := range pids {
			if pid == strconv.Itoa(n.Id) {
				return fmt.Errorf("不能循环指定父节点")
			}
		}
	}
	n.Level = parent.Level + 1
	if parent.KeyPath == "" {
		n.KeyPath = strconv.Itoa(parent.Id)
	} else {
		n.KeyPath = parent.KeyPath + "_" + strconv.Itoa(parent.Id)
	}
	if err := n.SaveBean(sess, n); err != nil {
		return err
	}
	subs := []*Navtree{}
	sess.Where("parent_id = ?", n.Id).Find(&subs)
	for _, s := range subs {
		if err := s.Save(sess); err != nil {
			return err
		}
	}
	return nil
}

func (n *Navtree) Delete(sess *g.Sess) error {
	subs := []*Navtree{}
	sess.Where("parent_id = ?", n.Id).Find(&subs)
	for _, s := range subs {
		s.Delete(sess)
	}
	return n.DeleteBean(sess, n)
}
