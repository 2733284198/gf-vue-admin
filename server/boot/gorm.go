package boot

import (
	"database/sql"
	system "gf-vue-admin/app/model/system"
	"gf-vue-admin/boot/internal"
	"gf-vue-admin/library/config"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) Initialize() {
	switch global.Config.System.DbType {
	case "mysql":
		Mysql.Initialize()
	}
}

func init() {
	global.Config.Mysql = config.Mysql{
		Path:          g.Cfg().GetString("database.default.Host") + ":" + g.Cfg().GetString("database.default.Port"),
		Config:        "charset=utf8mb4&parseTime=True&loc=Local",
		Dbname:        g.Cfg().GetString("database.default.Name"),
		Username:      g.Cfg().GetString("database.default.User"),
		Password:      g.Cfg().GetString("database.default.Pass"),
		MaxIdleConnes: 10,
		MaxOpenConnes: 10,
		LogMode:       true,
		LogZap:        "",
	}
}

var Mysql = new(_mysql)

type _mysql struct {
	db  *gorm.DB
	err error
	sql *sql.DB
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: gorm连接mysql数据库
func (m *_mysql) Initialize() {
	if m.db, m.err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       global.Config.Mysql.Dsn(), // DSN data source name
		DefaultStringSize:         191,                       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                     // 根据版本自动配置
	}), internal.GenerateConfig()); m.err != nil {
		g.Log().Error(`Gorm连接MySQL异常!`, g.Map{"err": m.err})
		os.Exit(0)
	} else {
		m.AutoMigrateTables()
		if m.sql, m.err = m.db.DB(); m.err != nil {
			g.Log().Error(`DatabaseSql对象获取异常!`, g.Map{"err": m.err})
		}
		m.sql.SetMaxIdleConns(global.Config.Mysql.GetMaxIdleConnes())
		m.sql.SetMaxOpenConns(global.Config.Mysql.GetMaxOpenConnes())
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: gorm 同步模型 生成mysql表
func (m *_mysql) AutoMigrateTables() {
	m.err = m.db.AutoMigrate(
		new(system.Api),
		new(system.Admin),
		new(system.BaseMenu),
		new(system.Authority),
		new(system.Dictionary),
		new(system.JwtBlacklist),
		new(system.DictionaryDetail),
		new(system.BaseMenuParameter),
	)
	if m.err != nil {
		g.Log().Error(`注册表失败!`, g.Map{"err": m.err})
		os.Exit(0)
	}
	g.Log().Info(`注册表成功!`)
}
