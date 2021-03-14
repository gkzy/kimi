package router

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/kimi/src/Connect-Service/handler"
)

// APIRouter
func APIRouter(r *gow.Engine) {
	r.GET("/v1/ws", handler.WebSocketHandler)
}
