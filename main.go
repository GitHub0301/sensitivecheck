package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"sensitivecheck/dao"
	_ "sensitivecheck/router"
)

func init() {

	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		//return err
	}

	addr := beego.AppConfig.String("mysql::addr")
	userName := beego.AppConfig.String("mysql::user")
	password := beego.AppConfig.String("mysql::password")
	dbName := beego.AppConfig.String("mysql::database")

	orm.RegisterDataBase("default", "mysql", userName+":"+password+"@tcp("+addr+")/"+dbName+"?charset=utf8&parseTime=true&loc=Local")

	// 打开调试模式，开发的时候方便查看orm生成什么样子的sql语句
	dao.InitModel()
	orm.Debug = true

}

func main() {

	// flag redefined: graceful 可能是引用到其他文件的路径
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 1800

	beego.Run()

}
