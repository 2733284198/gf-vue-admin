package boot

import (
	"fmt"
	"gf-vue-admin/library/global"
	"gf-vue-admin/library/utils"
	"gf-vue-admin/router"
	"time"

	"github.com/gogf/gf/frame/g"
)

var Server = new(_server)

type _server struct{}

func (s *_server) Initialize() {
	var server = g.Server()
	server.SetReadTimeout(10 * time.Second)
	server.SetWriteTimeout(10 * time.Second)
	server.SetMaxHeaderBytes(1 << 20)
	server.SetIndexFolder(true)
	if global.Config.System.OssType == "local" {
		_ = utils.CreateDir(g.Cfg("oss").GetString("local.LocalPath"))
		server.AddStaticPath("/"+g.Cfg("oss").GetString("local.LocalPath"), g.Cfg("oss").GetString("local.LocalPath"))
	}
	server.AddStaticPath("/form-generator", "public/page")
	router.Routers.Init()
	fmt.Printf(`
	欢迎使用 Gf-Vue-Admin
	当前版本:V1.1.2
	默认前端文件运行地址:http://127.0.0.1:8080
	`)
	server.SetPort(8888)
	server.Run()
}
