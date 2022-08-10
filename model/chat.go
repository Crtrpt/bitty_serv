package model

import "time"

//可以理解为某个用户的通讯录列表
type Chat struct {
	Id int64 `json:"id"`

	SessionId string `json:"session_id" xorm:"varchar(32) session_id comment('session_id')"`
	//发送者用户id
	SenderId string `json:"sender_id" xorm:"varchar(32) sender_id comment('sender_id')"`
	//发送者消息的序号
	Sn int `json:"sn" xorm:"int sn comment('sn')"`
	//发送消息的类型
	Type int `json:"type" xorm:"tinyint  'type' comment('type')"`
	//发送消息的内容
	Content string `json:"content" xorm:"varchar(256) content comment('content')"`
	//发送消息的负载
	Payload string `json:"payload" xorm:"varchar(256) payload comment('payload')"`
	//
	CreatedAt time.Time `json:"created_at" xorm:"timestamp  created comment('created_at')"`
}
