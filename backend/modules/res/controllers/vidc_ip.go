package controllers

import (
	"backend/g"
	"backend/modules/res/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type VidcIp struct {
	*g.XB[*models.VidcIp]
}

func NewVidcIp() *VidcIp {
	r := &VidcIp{XB: g.NewXB(&models.VidcIp{})}
	r.DB = g.CoreDB
	return r
}

func (r *VidcIp) ActionListIp(c *gin.Context) {
	id := c.Query("id")
	vidc := new(models.Vidc)
	r.DB.ID(id).Get(vidc)
	args := map[string]string{"vidc_id": id}
	title := fmt.Sprintf("<code>%s</code> IP列表", vidc.Name)
	r.List(c, args, title, "80%", "", true)
}
