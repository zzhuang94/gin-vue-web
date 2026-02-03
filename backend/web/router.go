package web

import (
	"backend/g"
	"backend/web/frm"
	"backend/web/modules/base"
	"backend/web/modules/prod"
	"backend/web/modules/res"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup) {
	initSession(rg)

	rg.Use(frm.Middleware)

	routeBase()
	routeRes()
	routeProd()

	frm.BindActions(rg)
}

func initSession(rg *gin.RouterGroup) {
	store, err := redis.NewStore(10, "tcp",
		g.C.Redis.Addr, "", g.C.Redis.Passwd,
		[]byte("hzz_session_secret"),
	)
	if err != nil {
		panic(err)
	}

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60,
		HttpOnly: true,
		Secure:   false,
	})
	rg.Use(sessions.Sessions("mysession", store))
}

func routeBase() {
	frm.RegController("base", "helper", base.NewHelper())
	frm.RegController("base", "user", base.NewUser())
	frm.RegController("base", "user-log", base.NewUserLog())
	frm.RegController("base", "action", base.NewAction())
	frm.RegController("base", "navtree", base.NewNavtree())
	frm.RegController("base", "role", base.NewRole())
	frm.RegController("base", "role-user", base.NewRoleUser())
	frm.RegController("base", "trans", base.NewTrans())
	frm.RegController("base", "chart", base.NewChart())
	frm.RegController("base", "op", base.NewOp())
}

func routeRes() {
	frm.RegController("res", "vidc", res.NewVidc())
	frm.RegController("res", "ip", res.NewIp())
	frm.RegController("res", "vidc-ip", res.NewVidcIp())
	frm.RegController("res", "service", res.NewService())
	frm.RegController("res", "policy", res.NewPolicy())
}

func routeProd() {
	frm.RegController("prod", "ticket", prod.NewTicket())
	frm.RegController("prod", "material", prod.NewMaterial())
	frm.RegController("prod", "color", prod.NewColor())
	frm.RegController("prod", "machine", prod.NewMachine())
}
