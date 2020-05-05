package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/qianbaidu/chrome-server/util"
	_ "github.com/qianbaidu/chrome-server/routers"
	"github.com/qianbaidu/chrome-server/models"
	_ "github.com/astaxie/beego/session/mysql"
)

func init() {
	models.RegisterDb()
	orm.RunSyncdb("default", false, false)
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = util.GetMysqlDns()
}

func main() {
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}


	beego.Run()
}
