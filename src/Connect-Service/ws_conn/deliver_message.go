package ws_conn

import (
	"context"
	"github.com/gkzy/gow/lib/logy"
	"github.com/gkzy/kimi/common/repository/pb"
)

// DeliverMessage 投递消息
func DeliverMessage(ctx context.Context, requestId int64, req *pb.DeliverMessageReq) error {
	conn := Load(req.DeviceId)
	if conn == nil {
		logy.Errorf("conn is nil")
		return nil
	}

	conn.Output(pb.PackageType_PT_MESSAGE, requestId, req.MessageSend)

	return nil
}
