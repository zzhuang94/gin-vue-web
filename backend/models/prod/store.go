package prod

import (
	"backend/g"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Store struct {
	g.Model `xorm:"extends"`

	Category string `xorm:"category" json:"category"`  // 类型
	Material string `xorm:"material" json:"material"`  // 材质
	Color    string `xorm:"color" json:"color"`        // 颜色
	Goods    int    `xorm:"goods" json:"goods,string"` // 良品
	Bads     int    `xorm:"bads" json:"bads,string"`   // 劣品
	Remark   string `xorm:"remark" json:"remark"`      // 备注
}

func (Store) TableName() string {
	return "store"
}

func (Store) New() g.ModelX {
	return &Store{}
}

func (s *Store) Save(sess *g.Sess) error {
	if s.Id == 0 {
		exists, err := sess.Where(
			"category = ? AND material = ? AND color = ?",
			s.Category, s.Material, s.Color,
		).Exist(&Store{})
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("该类型、材质、颜色已存在")
		}
		return s.SaveBean(sess, s)
	}
	if s.Goods < 0 || s.Bads < 0 {
		return fmt.Errorf("良品库存[%d]和劣品库存[%d]不能小于0", s.Goods, s.Bads)
	}
	op := &StoreOp{
		StoreId: s.Id,
		Op:      OpEdit,
		Count:   s.Goods,
		Remark:  fmt.Sprintf("良品:%d,劣品:%d, 备注:%s", s.Goods, s.Bads, s.Remark),
		User:    sess.Ctx.GetString("username"),
	}
	if err := op.Save(sess); err != nil {
		return err
	}
	return s.SaveBean(sess, s)
}

func (s *Store) Delete(sess *g.Sess) error {
	return s.DeleteBean(sess, s)
}

func (s *Store) Plus(sess *g.Sess, count int, remark, src string) error {
	s.Goods += count
	if err := s.SaveBean(sess, s); err != nil {
		return err
	}
	op := &StoreOp{
		StoreId: s.Id,
		Op:      OpPlus,
		Count:   count,
		Remark:  remark,
		User:    sess.Ctx.GetString("username"),
	}
	if err := op.Save(sess); err != nil {
		return err
	}
	if src == "gen" {
		s.updateTickets(sess, count, 0)
	}
	return nil
}

func (s *Store) Minus(sess *g.Sess, count int, remark string) error {
	if s.Goods < count {
		return fmt.Errorf("良品库存[%d]不足 %d，无法出库", s.Goods, count)
	}
	s.Goods -= count
	if err := s.SaveBean(sess, s); err != nil {
		return err
	}
	op := &StoreOp{
		StoreId: s.Id,
		Op:      OpMinus,
		Count:   count,
		Remark:  remark,
		User:    sess.Ctx.GetString("username"),
	}
	return op.Save(sess)
}

func (s *Store) Reject(sess *g.Sess, count int, remark, src string) error {
	s.Bads += count
	if err := s.SaveBean(sess, s); err != nil {
		return err
	}
	op := &StoreOp{
		StoreId: s.Id,
		Op:      OpReject,
		Count:   count,
		Remark:  remark,
		User:    sess.Ctx.GetString("username"),
	}
	if err := op.Save(sess); err != nil {
		return err
	}
	if src == "gen" {
		s.updateTickets(sess, 0, count)
	}
	return nil
}

func (s *Store) updateTickets(sess *g.Sess, goods, bads int) (int, int) {
	logrus.Infof("updateTickets: goods=%d, bads=%d", goods, bads)

	ts := s.getRunningTickets(sess)
	if len(ts) == 0 {
		return goods, bads
	}
	for _, t := range ts {
		goods, bads = t.tryUpdateProgress(sess, goods, bads)
		if goods == 0 && bads == 0 {
			break
		}
	}
	return goods, bads
}

func (s *Store) getRunningTickets(sess *g.Sess) []*Ticket {
	ts := []*Ticket{}
	err := sess.Where(
		"store_id = ? AND status = ?",
		s.Id, StatusRunning,
	).OrderBy("lead_time ASC").Find(&ts)
	if err != nil {
		logrus.Errorf("getRunningTickets: %v", err)
		return nil
	}
	return ts
}
