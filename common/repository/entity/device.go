package entity

const (
	DeviceOnLine  = 1 // 设备在线
	DeviceOffLine = 0 // 设备离线
)

// Device 设备
type Device struct {
	Id            int64     // 设备id
	UserId        int64     // 用户id
	Type          int32    // 设备类型,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web
	Brand         string    // 手机厂商
	Model         string    // 机型
	SystemVersion string    // 系统版本
	SDKVersion    string    // SDK版本
	Status        int32     // 在线状态，0：不在线；1：在线
	ConnAddr      string    // 连接层服务层地址
	ConnFd        int64     // TCP连接对应的文件描述符
	ClientAddr    string    // 客户端地址
	Created   int64 // 创建时间
	Updated    int64 // 更新时间
}
