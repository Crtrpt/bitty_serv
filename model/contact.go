package model

import "time"

//可以理解为某个用户的通讯录列表
type Contact struct {
	Id int64 `json:"id"`
	//用户的id
	UserId string `json:"user_id" xorm:"varchar(32) user_id comment('name')"`
	//目标用户的id
	TargetId string `json:"target_id" xorm:"varchar(32) target_id comment('TargetId')"`
	//sessionId
	SessionId string `json:"session_id" xorm:"varchar(32) session_id comment('session_id')"`
	//用户的备注信息
	Name string `json:"name" xorm:"varchar(32)  name comment('name')"`
	//创建时间
	CreatedAt time.Time `xorm:"timestamp  created comment('created_at')"`
}
