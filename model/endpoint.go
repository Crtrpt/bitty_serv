package model

import "time"

//可以理解为某个用户的通讯录列表
type Endpoint struct {
	Id int64
	//用户的id
	UserId int64 `xorm:"varchar(25)  comment('name')"`
	//显示的用户名
	Name string `xorm:"varchar(25)  comment('name')"`
	//用户名头像
	Avatar string `xorm:"varchar(100)  comment('avatar')"`
	//创建时间
	CreatedAt time.Time `xorm:"timestamp  created comment('created_at')"`
	//成员数量
	MemberAmount int `xorm:"tinyint  comment('endpoint 0 1to1')"`
	//聊条场景
	Scenes int `xorm:"tinyint  comment('endpoint 0 1to1')"`
}
