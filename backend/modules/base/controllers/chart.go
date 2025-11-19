package controllers

import (
	"backend/g"

	"github.com/gin-gonic/gin"
)

type Chart struct {
	*g.Web
}

func NewChart() *Chart {
	return &Chart{Web: g.NewWeb()}
}

func (r *Chart) ActionIndex(c *gin.Context) {
	r.Render(c)
}
