package rpc

import (
	"context"
	"github.com/gkzy/kimi/common/repository/pb"
	"github.com/gkzy/kimi/src/Connect-Service/ws_conn"
)

// ConnMessageGRPCService connect-service 提供的消息服务相关grpc
type ConnMessageGRPCService struct{}

// DeliverMessage 消息投递
func (*ConnMessageGRPCService) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (resp *pb.DeliverMessageResp, err error) {
	return &pb.DeliverMessageResp{}, ws_conn.DeliverMessage(ctx, req.DeviceId, req)
}
