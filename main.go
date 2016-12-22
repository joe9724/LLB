package main

import (
	_ "LLB/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.SetLogger(logs.AdapterFile,`{"filename":"opration.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":100}`)
	beego.Run()
}

