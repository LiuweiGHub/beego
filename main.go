package main

import (
	_ "bee/routers"
	_ "github.com/prometheus/client_golang/prometheus"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

