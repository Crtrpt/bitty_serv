package model

import "time"

//可以理解为某个用户的通讯录列表
type Endpoint struct {
	Id int64	`json:"-"`
	//用户的id
	UserId string `json:"userid" xorm:"varchar(32) user_id comment('name')"`
	//显示的用户名
	Name string `json:"name" xorm:"varchar(32)  name comment('name')"`
	//用户名头像
	Avatar string `json:"avatar" xorm:"varchar(100)  avatar comment('avatar')"`
	//创建时间
	CreatedAt time.Time `xorm:"timestamp  created comment('created_at')"`
	//成员数量
	MemberAmount int `json:"member_amount" xorm:"tinyint  member_amount comment('endpoint 0 1to1')"`
	//聊条场景
	Scenes int `json:"scenes"  xorm:"tinyint  scenes comment('endpoint 0 1to1')"`
	//场景id
	EndpointId string `json:"endpoint_id" xorm:"varchar(32)  not null endpoint_id comment('endpoint_id')"`
	//对端id
}
