package modules

import (
	"backend/g"
	base "backend/modules/base/controllers"
	res "backend/modules/res/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup) {
	initSession(rg)

	rg.Use(g.WebAuth)

	routeBase()
	routeRes()

	g.BindActions(rg)
}

func initSession(rg *gin.RouterGroup) {
	store, err := redis.NewStore(10, "tcp", "gin-vue.web.domain:6379", "", "hzz123", []byte("hzz_session_secret"))
	if err != nil {
		panic(err)
	}

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   24 * 60 * 60,
		HttpOnly: true,
		Secure:   false,
	})
	rg.Use(sessions.Sessions("mysession", store))
}

func routeBase() {
	g.RegController("base", "helper", base.NewHelper())
	g.RegController("base", "user", base.NewUser())
	g.RegController("base", "user-log", base.NewUserLog())
	g.RegController("base", "action", base.NewAction())
	g.RegController("base", "navtree", base.NewNavtree())
	g.RegController("base", "role", base.NewRole())
	g.RegController("base", "role-user", base.NewRoleUser())
	g.RegController("base", "trans", base.NewTrans())
	g.RegController("base", "chart", base.NewChart())
}

func routeRes() {
	g.RegController("res", "vidc", res.NewVidc())
	g.RegController("res", "ip", res.NewIp())
	g.RegController("res", "vidc-ip", res.NewVidcIp())
	g.RegController("res", "service", res.NewService())
	g.RegController("res", "policy", res.NewPolicy())
}
