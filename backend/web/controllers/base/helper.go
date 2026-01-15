package base

import (
	"backend/g"

	"github.com/gin-gonic/gin"
)

type Helper struct {
	*g.Web
}

func NewHelper() *Helper {
	return &Helper{Web: g.NewWeb()}
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
