package main

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/kimi/src/Logic-Service/conn"
	"github.com/gkzy/kimi/src/Logic-Service/handler"
	"github.com/gkzy/kimi/src/Logic-Service/middleware"
	"github.com/gkzy/kimi/src/Logic-Service/rpc"
)

func init() {
	conn.InitLog()
	conn.InitMysql()
	rpc.InitGRPCServer()
}

func main() {
	r := gow.Default()
	r.SetAppConfig(gow.GetAppConfig())
	r.Use(middleware.Auth())
	r.GET("/send", handler.Send)
	r.Run()
}
