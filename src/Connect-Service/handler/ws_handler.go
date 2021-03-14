package handler

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/gow/lib/logy"
	"github.com/gkzy/kimi/src/Connect-Service/ws_conn"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 65535,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// WebSocketHandler
func WebSocketHandler(c *gow.Context) {
	userId, _ := c.GetInt64("user_id", 0)
	deviceId, _ := c.GetInt64("device_id", 0)
	token := c.GetString("token")
	// requestId, _ := c.GetInt64("request_id", 0)

	if userId == 0 || deviceId == 0 || token == "" {
		c.DataJSON(100, "请重新登录")
		return
	}

	//TODO:去登录和验证

	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logy.Errorf("websocket upgrade error:%v", err)
		return
	}

	preCtx := ws_conn.Load(deviceId)
	if preCtx != nil {
		preCtx.DeviceId = ws_conn.PreConn
	}

	ctx := ws_conn.NewWsContext(deviceId, userId, conn)
	ws_conn.Store(deviceId, ctx)
	ctx.StartConn()

}
