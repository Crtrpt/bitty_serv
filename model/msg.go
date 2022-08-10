package model

import "time"

type Msg struct {
	Id int64 `json:"id"`

	SourceId string `json:"source_id" xorm:"varchar(32)  not null source_id comment('source_id')"`
	//接受者
	TargetId string `json:"target_id" xorm:"varchar(32)  not null target_id comment('target_id')"`
	//消息的内容 文字形式
	Content string `json:"content" xorm:"varchar(100) not null  'content' comment('content')"`
	//消息类型
	Type int `json:"type" xorm:"tinyint  'type' comment('type ')"`
	//消息负载
	Payload string `json:"payload" xorm:"json not null  'payload' comment('payload')"`
	//创建时间
	CreatedAt time.Time `json:"created_at" xorm:"timestamp  created comment('created_at')"`
	//处理结果
	Result int `json:"result" xorm:"tinyint  'result' comment('result')"`
	//处理结果附带直
	ResultPayload string `json:"result_payload" xorm:"varchar(32)  'result_payload' comment('result_payload')"`
	//消息级别
	Level int `json:"level" xorm:"tinyint  'level' comment('level')"`
	//是否已读
	IsRead int ` json:"isread" xorm:"tinyint  'isread' comment('isread')"`
	//处理状态
	Status int `json:"status" xorm:"tinyint  'status' comment('status')"`
}
