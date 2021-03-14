package handler

import (
	"fmt"
	"github.com/gkzy/gow"
	"github.com/gkzy/kimi/common/repository/pb"
	"github.com/gkzy/kimi/src/Logic-Service/models"
	"github.com/gkzy/kimi/src/Logic-Service/service"
	"time"
)

func Send(c *gow.Context) {
	sender := &models.Sender{
		Nickname:   "新月却泽滨",
		SenderType: 1,
		SenderId:   1,
		DeviceId:   1234567890,
		Avatar:     "https://p1.ymzy.cn/ave/20210311/d58400fb555a34aa.jpg",
	}

	req := &pb.SendMessageReq{
		ReceiverType:   1,
		ReceiverId:     2,
		MessageType:    1,
		MessageContent: []byte("这是我的第一条消息"),
		SendTime:       time.Now().Unix(),
		IsPersist:      false,
	}

	seq, err := new(service.MessageService).Send(nil, sender, req)
	if err != nil {
		c.DataJSON(1, fmt.Sprintf("Send ERROR:%v", err))
		return
	}

	c.DataJSON(0, seq)
}
