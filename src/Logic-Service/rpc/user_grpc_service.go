package rpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/gkzy/kimi/common/repository/pb"
	"github.com/gkzy/kimi/src/Logic-Service/models"
)

// UserGRPCService 用户相关 GRPC
type UserGRPCService struct{}

// RegisterDevice 注册设备
func (*UserGRPCService) RegisterDevice(ctx context.Context, req *pb.RegisterDeviceReq) (resp *pb.RegisterDeviceResp, err error) {
	device := &models.Device{}
	device.Type = req.Type
	device.Brand = req.Brand
	device.Model = req.Brand
	device.SystemVersion = req.SystemVersion
	device.SDKVersion = req.SystemVersion

	if device.Type == 0 || device.Brand == "" || device.Model == "" || device.SystemVersion == "" || device.SDKVersion == "" {
		err = errors.New("请求参数错误")
		return
	}

	err = device.Register()
	if err != nil {
		err = fmt.Errorf("写入DB失败:%v", err)
		return
	}

	resp = &pb.RegisterDeviceResp{
		DeviceId: device.Id,
	}

	return
}

// GetDevice 获取设备信息
func (*UserGRPCService) GetDevice(ctx context.Context, req *pb.GetDeviceReq) (resp *pb.GetDeviceResp, err error) {
	device := new(models.Device)
	err = device.Get(req.DeviceId)
	if err != nil {
		err = fmt.Errorf("设备不存在:%v", err)
		return
	}

	resp = &pb.GetDeviceResp{
		Device: &pb.Device{
			DeviceId:      device.Id,
			UserId:        device.UserId,
			Type:          device.Type,
			Brand:         device.Brand,
			Model:         device.Model,
			SystemVersion: device.SystemVersion,
			SDKVersion:    device.SDKVersion,
			Status:        device.Status,
			Created:       device.Created,
			Updated:       device.Updated,
		},
	}
	return
}

// SendMessage 发送消息
func (*UserGRPCService) SendMessage(ctx context.Context, req *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	return nil, nil

}

func (*UserGRPCService) ConnUserLogin(ctx context.Context, req *pb.ConnUserLoginReq) (resp *pb.ConnUserLoginResp, err error) {
	return nil, nil
}

func (*UserGRPCService) Offline(ctx context.Context, req *pb.OfflineReq) (resp *pb.OfflineResp, err error) {
	return nil, nil
}
