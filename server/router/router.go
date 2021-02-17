package router

import (
	"github.com/gogf/gf/frame/g"
)

var Routers = new(routers)

type routers struct{}

func (r *routers) Init() {
	var public = g.Server().Group("")
	{
		NewBaseGroup(public).Init()
	}
	//var private = g.Server().Group("").Middleware(middleware.JwtAuth, middleware.CasbinMiddleware)
	var private = g.Server().Group("").Middleware(JwtAuth)
	{ // 需要Jwt鉴权, casbin鉴权
		NewApiRouter(private).Init()
		NewAdminGroup(private).Init()
		NewMenuRouter(private).Init()
		NewCasbinRouter(private).Init()
		NewAuthorityRouter(private).Init()
		NewDictionaryRouter(private).Init()
		NewJwtBlacklistGroup(private).Init()
		NewDictionaryDetailRouter(private).Init()
	}
}
