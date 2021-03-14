package models



// User 用户
type User struct {
	Userid   int64  //  userid
	NickName string // 昵称
	Sex      int    // 性别
	Avatar   string // 头像
	Updated  int64  // 更新
	Created  int64  // 入库时间
}
