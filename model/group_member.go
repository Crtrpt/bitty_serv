package model

import "time"

//群组
type GroupMember struct {
	Id int64 `json:"-"`
	//群组id
	GroupId string `json:"group_id" xorm:"varchar(32) group_id comment('group_id')"`
	//sessionId
	UserId string `json:"user_id" xorm:"varchar(32) user_id comment('user_id')"`
	//加入群组时间
	CreatedAt time.Time `json:"created_at" xorm:"timestamp  created comment('created_at')"`
	//消息类型 0  管理员
	Type int `json:"type" xorm:"tinyint  'type' comment('type ')"`
}
