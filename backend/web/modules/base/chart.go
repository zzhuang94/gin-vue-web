package base

import (
	"backend/web/frm"

	"github.com/gin-gonic/gin"
)

type Chart struct {
	*frm.Web
}

func NewChart() *Chart {
	return &Chart{Web: frm.NewWeb()}
}

func (r *Chart) ActionIndex(c *gin.Context) {
	r.Render(c)
}
