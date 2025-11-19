package controllers

import (
	"backend/g"
	"backend/modules/base/models"
)

type RoleUser struct {
	*g.XB[*models.RoleUser]
}

func NewRoleUser() *RoleUser {
	return &RoleUser{XB: g.NewXB(&models.RoleUser{})}
}
