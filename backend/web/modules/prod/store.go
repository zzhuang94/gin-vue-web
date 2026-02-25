package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"

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

	r.Tool = []*frm.Tool{}
	r.BuildOption = func(c *gin.Context) []*frm.Option {
		return r.buildOption(c)
	}
	return r
}

func (s *Store) buildOption(c *gin.Context) []*frm.Option {
	user := s.GetUser(c)
	ans := []*frm.Option{}
	if user.IsStorekeeper || user.IsManager {
		ans = append(ans, s.WrapOption([]any{
			"出库", "upload", "out", "modal", []string{"id"},
		}))
		ans = append(ans, s.WrapOption([]any{
			"编辑", "edit", "edit", "modal", []string{"id"},
		}))
	}

	ans = append(ans, s.WrapOption([]any{
		"变更记录", "history", "history", "modal", []string{"id"},
	}))
	return ans
}

func (s *Store) ActionHistory(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	ops := []*prod.StoreOp{}
	g.CoreDB.Where("store_id = ?", id).OrderBy("id DESC").Limit(100).Find(&ops)
	s.Modal(c, map[string]any{"ops": ops})
}

func (s *Store) ActionOut(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	store := &prod.Store{}
	has, err := g.CoreDB.ID(id).Get(store)
	if err != nil || !has {
		s.JsonFail(c, err)
		return
	}
	s.Modal(c, map[string]any{"store": store})
}
