package prod

import (
	"backend/g"
	"fmt"
)

type Store struct {
	g.Model `xorm:"extends"`

	Category string `xorm:"category" json:"category"` // 类型
	Material string `xorm:"material" json:"material"` // 材质
	Color    string `xorm:"color" json:"color"`       // 颜色
	Count    int    `xorm:"count" json:"count"`       // 数量
	Remark   string `xorm:"remark" json:"remark"`     // 备注
}

func (Store) TableName() string {
	return "store"
}

func (Store) New() g.ModelX {
	return &Store{}
}

func (s *Store) Save(sess *g.Sess) error {
	return s.SaveBean(sess, s)
}

func (s *Store) Delete(sess *g.Sess) error {
	return s.DeleteBean(sess, s)
}

func (s *Store) Plus(sess *g.Sess, ticket *Ticket, count int, remark string) error {
	s.Count += count
	if err := s.Save(sess); err != nil {
		return err
	}
	op := &StoreOp{
		StoreId:  s.Id,
		TicketId: ticket.Id,
		Op:       OpPlus,
		Count:    count,
		Remark:   remark,
		User:     sess.Ctx.GetString("username"),
	}
	return op.Save(sess)
}

func (s *Store) Minus(sess *g.Sess, ticket *Ticket, count int, remark string) error {
	if s.Count < count {
		return fmt.Errorf("库存[%d]不足 %d，无法出库", s.Count, count)
	}
	s.Count -= count
	if err := s.Save(sess); err != nil {
		return err
	}
	op := &StoreOp{
		StoreId:  s.Id,
		TicketId: ticket.Id,
		Op:       OpMinus,
		Count:    count,
		Remark:   remark,
		User:     sess.Ctx.GetString("username"),
	}
	return op.Save(sess)
}

func (s *Store) Edit(sess *g.Sess, ticket *Ticket, count int, remark string) error {
	if count < 0 {
		return fmt.Errorf("库存数不能小于0")
	}
	s.Count = count
	s.Remark = remark
	if err := s.Save(sess); err != nil {
		return err
	}
	op := &StoreOp{
		StoreId:  s.Id,
		TicketId: ticket.Id,
		Op:       OpEdit,
		Count:    count,
		Remark:   remark,
		User:     sess.Ctx.GetString("username"),
	}
	return op.Save(sess)
}
