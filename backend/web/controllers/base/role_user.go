package base

import (
	"backend/g"
	"backend/models/base"
)

type RoleUser struct {
	*g.XB[*base.RoleUser]
}

func NewRoleUser() *RoleUser {
	return &RoleUser{XB: g.NewXB(&base.RoleUser{})}
}
