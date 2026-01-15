package main

import (
	"backend/api"
	"backend/g"
	"backend/web"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := g.Init(); err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	r := gin.Default()
	wg := r.Group("/web")
	web.Route(wg)

	ag := r.Group("/api")
	api.Route(ag)

	r.Run(fmt.Sprintf(":%d", g.C.Port))
}
