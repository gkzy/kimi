package rpc

import (
	"context"
	"github.com/gkzy/kimi/common/repository/pb"
)

type MessageGRPCService struct {
}

func (*MessageGRPCService) SendMessage(ctx context.Context, req *pb.SendMessageReq) (resp *pb.SendMessageResp, err error) {
	return nil, nil
}

func (*MessageGRPCService) Sync(ctx context.Context, req *pb.SyncReq) (resp *pb.SyncResp, err error) {
	return nil, nil
}

func (*MessageGRPCService) MessageACK(ctx context.Context, req *pb.MessageACKReq) (resp *pb.MessageACKResp, err error) {
	return nil, nil
}
