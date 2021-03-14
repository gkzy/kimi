package ws_conn

import (
	"encoding/json"
	"github.com/gkzy/gow/lib/logy"
	"github.com/gkzy/kimi/common/repository/pb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"strings"
	"time"
)

const (
	// PreConn 设备第二次连接时，标记设备的上一条连接
	PreConn = -1

	// deadLineTimeOut 超时min
	deadLineTimeOut = 10
)

// WsContext websocket context
type WsContext struct {
	DeviceId int64
	UserId   int64
	Conn     *websocket.Conn
}

// NewWsContext return a new WsContext
func NewWsContext(deviceId, userId int64, conn *websocket.Conn) *WsContext {
	return &WsContext{
		DeviceId: deviceId,
		UserId:   userId,
		Conn:     conn,
	}
}

// StartConn 开始处理连接
func (c *WsContext) StartConn() {
	defer func() {
		err := recover()
		panic(err)
		//TODO: get stack
	}()

	for {
		err := c.Conn.SetReadDeadline(time.Now().Add(deadLineTimeOut * time.Minute))
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			c.HandleReadError(err)
			return
		}
		c.HandlePackage(data)
	}
}

// HandlePackage 处理请求的数据包
func (c *WsContext) HandlePackage(data []byte) {
	var input pb.Input
	err := json.Unmarshal(data, &input)
	if err != nil {
		logy.Errorf("handler package data error:%v", err)
		c.Release()
	}
	switch input.Type {
	case pb.PackageType_PT_SYNC:
		c.Sync(input)
	case pb.PackageType_PT_HEARTBEAT:
		c.Heartbeat(input)
	case pb.PackageType_PT_MESSAGE:
		c.MessageACK(input)
	default:
		logy.Errorf("unknown input type: %v ", input.Type)
	}
}

// Heartbeat 心跳
func (c *WsContext) Heartbeat(input pb.Input) {
	log.Printf("ws 心跳 \n")
	c.Output(pb.PackageType_PT_HEARTBEAT, input.RequestId, nil)
}

// Sync 同步离线消息
func (c *WsContext) Sync(input pb.Input) {
	var sync pb.SyncInput
	err := json.Unmarshal(input.Data, &sync)
	if err != nil {
		logy.Errorf("sync input error:%v", err)
		c.Release()
		return
	}

	//TODO: return true SyncResp
	resp := &pb.SyncResp{
		Messages: nil,
		HasMore:  false,
	}

	message := &pb.SyncOutput{
		Messages: resp.Messages,
		HasMore:  resp.HasMore,
	}
	c.Output(pb.PackageType_PT_SYNC, input.RequestId, message)
}

// MessageACK 消息回执
func (c *WsContext) MessageACK(input pb.Input) {
	var messageACK pb.MessageACK
	err := proto.Unmarshal(input.Data, &messageACK)
	if err != nil {
		logy.Errorf("message ack error:%v", err)
		c.Release()
		return
	}

	//TODO:具体的消息回执业务实现
	log.Printf("ws 消息回执 \n")
}

// Output 处理下行数据
func (c *WsContext) Output(pt pb.PackageType, requestId int64, message proto.Message) {
	var output = pb.Output{
		Type:      pt,
		RequestId: requestId,
	}

	if message != nil {
		msgBytes, err := json.Marshal(message)
		if err != nil {
			return
		}
		output.Data = msgBytes
	}

	outputByte, err := json.Marshal(&output)
	if err != nil {
		logy.Errorf("output message error:%v", err)
	}

	logy.Debugf("outputByte:%v", string(outputByte))

	err = c.Conn.WriteMessage(websocket.BinaryMessage, outputByte)
	if err != nil {
		logy.Errorf("write message error:%v", err)
		return
	}

}

// HandleReadError 读错误处理
func (c *WsContext) HandleReadError(err error) {

	errStr := err.Error()

	// 服务器主动关闭链接
	if strings.Contains(errStr, "use of closed network connection") {
		return
	}

	c.Release()

	// client主动关闭连接或者异常程序退出
	if err == io.EOF {
		return
	}

	if strings.HasSuffix(errStr, "i/o timeout") {
		return
	}

}

// Release release connection
func (c *WsContext) Release() {
	if c.DeviceId != PreConn {
		Delete(c.DeviceId)
	}

	err := c.Conn.Close()
	if err != nil {
		logy.Errorf("conn close error:%v", err)
	}

	//TODO:通知业务服务器，设备已经下线
	if c.DeviceId != PreConn {

	}
}
