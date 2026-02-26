package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Store struct {
	*frm.X
}

func NewStore() *Store {
	r := &Store{X: frm.NewX(&prod.Store{})}
	r.NoID = true
	r.DB = g.CoreDB
	r.Dump = true
	r.BuildOption = func(c *gin.Context) []*frm.Option {
		return r.buildOption(c)
	}
	return r
}

func (s *Store) buildOption(c *gin.Context) []*frm.Option {
	user := s.GetUser(c)
	ans := []*frm.Option{}
	ans = append(ans, s.WrapOption([]any{
		"查看二维码", "qrcode", "qrcode", "modal", []string{"id"},
	}))
	if user.IsStorekeeper || user.IsManager {
		ans = append(ans, s.WrapOption([]any{
			"良品入库", "download", "plus", "modal", []string{"id"},
		}))
		ans = append(ans, s.WrapOption([]any{
			"良品出库", "upload", "minus", "modal", []string{"id"},
		}))
		ans = append(ans, s.WrapOption([]any{
			"劣品上报", "warning", "reject", "modal", []string{"id"},
		}))
		ans = append(ans, s.WrapOption([]any{
			"库存编辑", "edit", "edit", "modal", []string{"id"},
		}))
	}
	ans = append(ans, s.WrapOption([]any{
		"变更历史", "history", "history", "modal", []string{"id"},
	}))
	return ans
}

func (s *Store) ActionDetail(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	store := &prod.Store{}
	has, err := g.CoreDB.ID(id).Get(store)
	if err != nil || !has {
		s.JsonFail(c, err)
		return
	}
	s.RenderData(c, map[string]any{"store": store})
}

func (s *Store) ActionQrcode(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	store := &prod.Store{}
	has, err := g.CoreDB.ID(id).Get(store)
	if err != nil || !has {
		s.JsonFail(c, err)
		return
	}
	url := fmt.Sprintf("http://%s/prod/store/detail?id=%s", g.C.Host, id)
	s.Modal(c, map[string]any{"store": store, "url": url})
}

func (s *Store) ActionHistory(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	ops := []*prod.StoreOp{}
	g.CoreDB.Where("store_id = ?", id).OrderBy("id DESC").Limit(100).Find(&ops)
	s.Modal(c, map[string]any{"ops": ops})
}

func (s *Store) ActionPlus(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	store := &prod.Store{}
	has, err := g.CoreDB.ID(id).Get(store)
	if err != nil || !has {
		s.JsonFail(c, err)
		return
	}
	s.Modal(c, map[string]any{"store": store})
}

func (s *Store) ActionMinus(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	store := &prod.Store{}
	has, err := g.CoreDB.ID(id).Get(store)
	if err != nil || !has {
		s.JsonFail(c, err)
		return
	}
	s.Modal(c, map[string]any{"store": store})
}

func (s *Store) ActionReject(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	store := &prod.Store{}
	has, err := g.CoreDB.ID(id).Get(store)
	if err != nil || !has {
		s.JsonFail(c, err)
		return
	}
	s.Modal(c, map[string]any{"store": store})
}

func (s *Store) ActionOp(c *gin.Context) {
	arg := &struct {
		Op     string `json:"op"`
		Id     string `json:"id"`
		Count  int    `json:"count"`
		Remark string `json:"remark"`
		Src    string `json:"src"`
	}{}
	if err := c.ShouldBindJSON(arg); err != nil {
		s.JsonFail(c, err)
		return
	}
	if arg.Count <= 0 {
		s.JsonFail(c, fmt.Errorf("数量必须为正整数"))
		return
	}

	store := &prod.Store{}
	has, err := g.CoreDB.ID(arg.Id).Get(store)
	if err != nil || !has {
		s.JsonFail(c, err)
		return
	}

	sess := s.BeginSess(g.CoreDB, c)
	switch arg.Op {
	case "PLUS":
		err = store.Plus(sess, arg.Count, arg.Remark, arg.Src)
	case "MINUS":
		err = store.Minus(sess, arg.Count, arg.Remark)
	case "REJECT":
		err = store.Reject(sess, arg.Count, arg.Remark, arg.Src)
	default:
		err = fmt.Errorf("操作类型错误")
	}
	if err != nil {
		sess.Rollback()
		s.JsonFail(c, err)
		return
	}
	sess.Commit()
	s.JsonSucc(c, "操作成功")
}
