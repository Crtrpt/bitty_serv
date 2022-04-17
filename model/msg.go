package model

import "time"

type Msg struct {
	Id int64
	//发送者
	createUserId int64
	//发送给谁
	EndpointId int64
	//消息的内容 文字形式
	Content string `xorm:"varchar(100) not null  'content' comment('content')"`
	//消息类型
	Type int `xorm:"tinyint  'type' comment('type ')"`
	//消息负载
	Payload string `xorm:"json not null  'payload' comment('payload')"`
	//创建时间
	CreatedAt time.Time `xorm:"timestamp  created comment('created_at')"`
}
