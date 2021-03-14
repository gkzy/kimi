package main

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/kimi/src/Connect-Service/conn"
	"github.com/gkzy/kimi/src/Connect-Service/handler"
	"github.com/gkzy/kimi/src/Connect-Service/router"
	"github.com/gkzy/kimi/src/Connect-Service/rpc"
)

func init() {
	conn.InitLog()
	rpc.InitGRPCServer()
}

func main() {
	r := gow.Default()
	r.SetAppConfig(gow.GetAppConfig())
	r.Static("/static","static")
	r.GET("/", handler.IndexHandler)
	router.APIRouter(r)
	r.Run()
}
