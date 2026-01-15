package base

import (
	"backend/web/frm"

	"github.com/gin-gonic/gin"
)

type Helper struct {
	*frm.Web
}

func NewHelper() *Helper {
	return &Helper{Web: frm.NewWeb()}
}

func (h *Helper) ActionFaIcon(c *gin.Context) {
	h.Render(c)
}

func (h *Helper) ActionWidget(c *gin.Context) {
	h.Render(c)
}

func (h *Helper) ActionPlayground(c *gin.Context) {
	h.Render(c)
}
