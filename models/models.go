package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/common/log"
	"github.com/qianbaidu/chrome-server/util"
)

func RegisterDb() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		log.Error("orm.RegisterDriver error ", err)
	}
	err = orm.RegisterDataBase("default", "mysql", util.GetMysqlDns())
	if err != nil {
		log.Error("orm.RegisterDataBase error ", err)
	}

	orm.RegisterModel(new(User), new(Collection))
}
