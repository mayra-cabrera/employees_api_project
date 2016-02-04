package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "employees_api_project/routers"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:secret_yeih@tcp(127.0.0.1:3306)/employees")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
