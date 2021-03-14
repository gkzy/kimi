/*
register deviceId
*/
package ws_conn

import "sync"

var manager sync.Map

// Store save deviceId and conn
func Store(deviceId int64, ctx *WsContext) {
	manager.Store(deviceId, ctx)
}

// Load get ws context
func Load(deviceId int64) *WsContext {
	val, ok := manager.Load(deviceId)
	if ok {
		return val.(*WsContext)
	}
	return nil
}

// Delete delete conn
func Delete(deviceId int64) {
	manager.Delete(deviceId)
}
