package res

import (
	"backend/g"
	"backend/models/res"
	"backend/web/frm"
)

type Vidc struct {
	*frm.XB[*res.Vidc]
}

func NewVidc() *Vidc {
	r := &Vidc{XB: frm.NewXB(&res.Vidc{})}
	r.TableMenu = append([][]string{{"IP管理", "list", "/res/vidc-ip/list-ip"}}, r.TableMenu...)
	r.DB = g.CoreDB
	return r
}
