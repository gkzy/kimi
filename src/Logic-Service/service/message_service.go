package service

import (
	"context"
	"fmt"
	"github.com/gkzy/gow/lib/logy"
	"github.com/gkzy/gow/lib/rpc"
	"github.com/gkzy/kimi/common/constname"
	"github.com/gkzy/kimi/common/repository/entity"
	"github.com/gkzy/kimi/common/repository/pb"
	"github.com/gkzy/kimi/src/Logic-Service/models"
)

const (
	MessageLimit     = 50
	MaxSyncBufLength = 65535
)

// MessageService 消息的业务 实现
type MessageService struct{}

func (m *MessageService) Send(ctx context.Context, sender *models.Sender, req *pb.SendMessageReq) (int64, error) {

	switch req.ReceiverType {
	case pb.ReceiverType_RT_USER:
		//发送者是用户
		if sender.SenderType == pb.SenderType_ST_USER {
			return m.SendToFriend(ctx, sender, req)
		} else {
			return m.SendToUser(ctx, sender, req.ReceiverId, req)
		}
	case pb.ReceiverType_RT_GROUP:
		logy.Info("send to group")
	}
	return 0, nil
}

// SendToUser 将消息发送给用户
func (m *MessageService) SendToUser(ctx context.Context, sender *models.Sender, toUserId int64, req *pb.SendMessageReq) (seq int64, err error) {
	//持久化消息到存储
	if req.IsPersist {

	}
	message := &pb.Message{
		Sender: &pb.Sender{
			SenderType: sender.SenderType,
			SenderId:   sender.SenderId,
			Avatar:     sender.Avatar,
			Nickname:   sender.Nickname,
		},
		ReceiverType:   req.ReceiverType,
		ReceiverId:     req.ReceiverId,
		ToUserIds:      req.ToUserIds,
		MessageType:    req.MessageType,
		MessageContent: req.MessageContent,
		Seq:            seq,
		SendTime:       req.SendTime,
		Status:         pb.MessageStatus_MS_NORMAL,
	}

	devices, err := new(models.Device).GetOnlineDeviceByUserId(toUserId)
	if err != nil {
		return
	}

	for _, item := range devices {
		if sender.DeviceId == item.Id {
			continue
		}
		err = m.SendToDevice(item, message)
		if err != nil {
			logy.Errorf("send to user error :%v", err)
			return
		}

	}
	return

}

// SendToFriend  消息发送到用户
func (m *MessageService) SendToFriend(ctx context.Context, sender *models.Sender, req *pb.SendMessageReq) (seq int64, err error) {
	seq, err = m.SendToUser(ctx, sender, sender.SenderId, req)
	if err != nil {
		return
	}

	//TODO: 用户需要增加自己的已经同步的序列号

	_, err = m.SendToUser(ctx, sender, req.ReceiverId, req)
	if err != nil {
		return
	}

	return
}

// SendToDevice 调用connect-service grpc ，发送消息
func (m *MessageService) SendToDevice(device *models.Device, message *pb.Message) error {
	if device.Status == entity.DeviceOnLine {
		messageSend := &pb.MessageSend{
			Message: message,
		}

		// 连接 connect-service 的grpc发送消息
		conn, err := rpc.NewClient(constname.ConnectServerRPCAddr, constname.ConnectServerRPCPort)
		if err != nil {
			err = fmt.Errorf("连接GRPC失败:%v", err)
			return err
		}

		defer conn.Close()
		client := pb.NewConnMessageClient(conn)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		req := &pb.DeliverMessageReq{
			DeviceId:    device.Id,
			Fd:          device.ConnFd,
			MessageSend: messageSend,
		}
		_, err = client.DeliverMessage(ctx, req)

		if err != nil {
			return err
		}

	}

	return nil

}
