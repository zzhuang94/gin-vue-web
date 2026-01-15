package base

import (
	"backend/models/base"
	"backend/web/frm"
)

type RoleUser struct {
	*frm.XB[*base.RoleUser]
}

func NewRoleUser() *RoleUser {
	return &RoleUser{XB: frm.NewXB(&base.RoleUser{})}
}
