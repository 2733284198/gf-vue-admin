package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type Base struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewBaseGroup(router *ghttp.RouterGroup) interfaces.Router {
	return &Base{router: router, response: &response.Handler{}}
}

func (b *Base) Init() {
	var base =  b.router.Group("/base")
	{
		base.POST("captcha", b.response.Handler()(api.Base.Captcha))
	}
}