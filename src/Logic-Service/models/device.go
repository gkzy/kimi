package models

import (
	"github.com/gkzy/gow/lib/mysql"
	"github.com/gkzy/kimi/common/repository/entity"
	"time"
)

type Device struct {
	*entity.Device
}

func (m *Device) TableName() string {
	return "tbl_device"
}

// Register 注册设备
func (m *Device) Register() error {
	db := mysql.GetORM()
	m.Created = time.Now().Unix()
	m.Updated = time.Now().Unix()
	return db.Model(m).Create(&m).Error
}

// Get 查询数据表的device
func (m *Device) Get(id int64) error {
	db := mysql.GetORM()
	return db.Model(m).Where("device_id=?", id).First(&m).Error
}

// GetOnlineDeviceByUserId 查询用户的所有在线设备
func (m *Device) GetOnlineDeviceByUserId(userId int64) (data []*Device, err error) {
	db:=mysql.GetORM()
	err = db.Model(m).Where("user_id=?",userId).Where("status=?",entity.DeviceOnLine).Find(&data).Error
	return
}
