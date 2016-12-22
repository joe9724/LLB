package main

import (
	_ "LLB/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	logs.SetLogger(logs.AdapterFile,`{"filename":"opration.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":100}`)
	orm.RegisterDataBase("default", "mysql", "pf:123456@tcp(192.168.30.103:63306)/ibus?charset=utf8", 30)



	beego.Run()
}

